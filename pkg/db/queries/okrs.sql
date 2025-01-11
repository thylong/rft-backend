-- okrs.sql

-- Get okrs with pagination and optional search
-- name: GetOkrs :many
SELECT 
    o.id AS okr_id, 
    o.name AS okr_name, 
    o.number AS okr_number, 
    o.year AS okr_year, 
    o.description AS okr_description,
    COALESCE(
        ARRAY_AGG(
            JSON_BUILD_OBJECT(
                'id', k.id,
                'name', k.name,
                'number', k.number,
                'description', k.description,
                'sponsor', k.sponsor,
                'kpis', k.kpis,
                'scope', k.scope,
                'initiatives', k.initiatives
            )
        ) FILTER (WHERE k.id IS NOT NULL), 
        '{}'::json[]
    ) AS key_results
FROM okrs o
LEFT JOIN okr_krs k ON o.id = k.okr_id
GROUP BY o.id, o.name, o.number, o.year, o.description
LIMIT $1 OFFSET $2;

-- Get total count for pagination
-- name: GetOkrsCount :one
SELECT COUNT(*) FROM okrs
WHERE (LOWER(name) LIKE LOWER('%' || $1 || '%') OR $1 = '');

-- Get a single okr by ID
-- name: GetOkrByID :one
SELECT 
    o.id AS okr_id, 
    o.name AS okr_name, 
    o.number AS okr_number, 
    o.year AS okr_year, 
    o.description AS okr_description,
    ARRAY_AGG(
        JSON_BUILD_OBJECT(
            'id', k.id,
            'name', k.name,
            'number', k.number,
            'description', k.description,
            'sponsor', k.sponsor,
            'kpis', k.kpis,
            'scope', k.scope,
            'initiatives', k.initiatives
        )
    ) AS key_results
FROM okrs o
LEFT JOIN okr_krs k ON o.id = k.okr_id
WHERE o.id = $1
GROUP BY o.id;

-- Insert a new okr
-- name: InsertOkr :one
WITH new_okr AS (
    INSERT INTO okrs (name, number, year, description)
    VALUES ($1, $2, $3, $4)
    RETURNING id, name, number, year, description
)
SELECT 
    o.id AS okr_id,
    o.name AS okr_name,
    o.number AS okr_number,
    o.year AS okr_year,
    o.description AS okr_description
FROM 
    new_okr o;

-- Delete an okr by ID
-- name: DeleteOkr :exec
WITH deleted_krs AS (
    DELETE FROM okr_krs
    WHERE okr_krs.okr_id = $1
)
DELETE FROM okrs
WHERE okrs.id = $1;

-- Insert a new key result
-- name: InsertKeyResult :one
INSERT INTO okr_krs (
    okr_id, 
    name, 
    number, 
    description, 
    sponsor, 
    kpis, 
    scope, 
    initiatives
)
VALUES 
    ($1, $2, $3, $4, $5, $6, $7, $8)
RETURNING 
    id, 
    okr_id, 
    name, 
    number, 
    description, 
    sponsor, 
    kpis, 
    scope, 
    initiatives;
