-- name: ListInventories :many
SELECT
    id,
    lessor_id,
    name,
    category,
    rent_per_day,
    rent_per_month,
    available_units,
    health,
    created_at,
    updated_at
FROM inventories
WHERE available_units > 0
ORDER BY created_at DESC
LIMIT $1 OFFSET $2;

-- name: ListInventoriesByCategory :many
SELECT
    id,
    lessor_id,
    name,
    category,
    rent_per_day,
    rent_per_month,
    available_units,
    health,
    created_at,
    updated_at
FROM inventories
WHERE available_units > 0
  AND category = $1
ORDER BY created_at DESC
LIMIT $2 OFFSET $3;

-- name: CountInventories :one
SELECT COUNT(*) FROM inventories
WHERE available_units > 0;

-- name: CountInventoriesByCategory :one
SELECT COUNT(*) FROM inventories
WHERE available_units > 0
  AND category = $1;
