-- name: InsertBlock :exec
INSERT INTO blocks (id, name, created_at, updated_at, deleted_at) VALUES (?, ?, ?, ?, ?);
