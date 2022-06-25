-- name: InsertUser :exec
INSERT INTO users (id, email, name, pass_hash) VALUES (?, ?, ?, ?);

-- name: GetUserByEmail :one
SELECT id, name, email, pass_hash FROM users WHERE email = ?;
