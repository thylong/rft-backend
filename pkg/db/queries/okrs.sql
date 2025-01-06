-- okrs.sql

-- Get okrs with pagination and optional search
-- name: GetOkrs :many
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
            'year', k.year,
            'description', k.description
        )
    ) AS key_results
FROM okrs o
LEFT JOIN okr_krs k ON o.id = k.okr_id
GROUP BY o.id
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
            'year', k.year,
            'description', k.description
        )
    ) AS key_results
FROM okrs o
LEFT JOIN okr_krs k ON o.id = k.okr_id
WHERE o.id = $1
GROUP BY o.id;

-- Insert a new okr
-- name: InsertOkr :exec
WITH new_okr AS (
    INSERT INTO okrs (id, name, number, year, description)
    VALUES ($1, $2, $3, $4, $5)
    RETURNING id
),
expanded_krs AS (
    SELECT
        UNNEST($6::uuid[]) AS kr_id,
        UNNEST($7::text[]) AS kr_name,
        UNNEST($8::int[]) AS kr_number,
        UNNEST($9::int[]) AS kr_year,
        UNNEST($10::text[]) AS kr_description,
        (SELECT id FROM new_okr) AS okr_id
)
INSERT INTO okr_krs (id, okr_id, name, number, year, description)
SELECT kr_id, okr_id, kr_name, kr_number, kr_year, kr_description
FROM expanded_krs;

-- Delete an okr by ID
-- name: DeleteOkr :exec
WITH deleted_krs AS (
    DELETE FROM okr_krs
    WHERE okr_krs.okr_id = $1
)
DELETE FROM okrs
WHERE okrs.id = $1;
