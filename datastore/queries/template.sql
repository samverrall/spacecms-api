-- name: InsertTemplate :exec
INSERT INTO templates (id, entry_block_id, name, created_at, updated_at, deleted_at) VALUES (?, ?, ?, ?, ?, ?);

-- name: GetTemplateByID :one
SELECT id, name, entry_block_id FROM templates WHERE id = ?
