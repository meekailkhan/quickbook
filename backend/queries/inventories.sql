-- name: ListInventories :many
SELECT
    i.id,
    i.lessor_id,
    i.name,
    i.category,
    i.rent_per_day,
    i.rent_per_month,
    i.available_units,
    i.health,
    l.username,
    l.email,
    l.phone,
    i.created_at,
    i.updated_at
FROM inventories as i
LEFT JOIN lessors as l
ON l.id = i.lessor_id
WHERE i.available_units > 0
ORDER BY i.created_at DESC
LIMIT $1 OFFSET $2;

-- name: ListInventoriesByCategory :many
SELECT
    i.id,
    i.lessor_id,
    i.name,
    i.category,
    i.rent_per_day,
    i.rent_per_month,
    i.available_units,
    i.health,
    l.username,
    l.email,
    l.phone,
    i.created_at,
    i.updated_at
FROM inventories AS i
LEFT JOIN lessors as l
ON l.id = i.lessor_id
WHERE available_units > 0
  AND i.category = $1
ORDER BY i.created_at DESC
LIMIT $2 OFFSET $3;

-- name: CountInventories :one
SELECT COUNT(*) FROM inventories
WHERE available_units > 0;

-- name: CountInventoriesByCategory :one
SELECT COUNT(*) FROM inventories
WHERE available_units > 0
  AND category = $1;
