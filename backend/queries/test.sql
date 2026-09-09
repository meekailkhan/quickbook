-- name: TestDB :one
SELECT
    id,
    test_name
FROM test
WHERE id = $1
LIMIT 1;
