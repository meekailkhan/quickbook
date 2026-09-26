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
FROM inventories AS i
LEFT JOIN lessors AS l
    ON l.id = i.lessor_id
WHERE i.available_units > 0
  AND (
      sqlc.arg(category)::text = ''
      OR i.category = sqlc.arg(category)::text
  )
ORDER BY i.created_at DESC
LIMIT sqlc.arg('limit')
OFFSET sqlc.arg('offset');


-- name: CountInventories :one
SELECT COUNT(*)
FROM inventories AS i
WHERE i.available_units > 0
  AND (
      sqlc.arg(category)::text = ''
      OR i.category = sqlc.arg(category)::text
);