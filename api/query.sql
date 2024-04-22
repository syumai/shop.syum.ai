-- name: GetProduct :one
SELECT * FROM products
WHERE id = ? LIMIT 1;

-- name: ListProducts :many
SELECT * FROM products
ORDER BY id;

-- name: CreateProduct :one
INSERT INTO products (
  name, description, product_status, created_at, updated_at, reserved_at
) VALUES (
  ?, ?, ?, ?, ?, ?
)
RETURNING *;

-- name: UpdateProduct :exec
UPDATE products
SET
  name = ?,
  description = ?,
  product_status = ?,
  created_at = ?,
  updated_at = ?,
  reserved_at = ?
WHERE id = ?;

-- name: DeleteProduct :exec
DELETE FROM products
WHERE id = ?;
