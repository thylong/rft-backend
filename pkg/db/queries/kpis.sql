-- kpis.sql

-- Get kpis
-- name: GetKpis :many
SELECT * FROM kpis
WHERE (LOWER(name) LIKE LOWER('%' || $1 || '%') OR $1 = '')
LIMIT $2 OFFSET $3;

-- Get total count for pagination
-- name: GetKpisCount :one
SELECT COUNT(*) FROM kpis
WHERE (LOWER(name) LIKE LOWER('%' || $1 || '%') OR $1 = '');

-- Get a single kpi by ID
-- name: GetKpiByID :one
SELECT * FROM kpis WHERE id = $1;

-- Insert a new kpi
-- name: InsertKpi :one
INSERT INTO kpis (
    name, value, target, day
) VALUES ($1, $2, $3, $4)
RETURNING *;

-- Delete a kpi by ID
-- name: DeleteKpi :exec
DELETE FROM kpis WHERE id = $1;
