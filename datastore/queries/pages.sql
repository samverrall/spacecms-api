-- name: InsertPage :exec
INSERT INTO pages
(id, template_id, url, canonical_url, title, description, thumbnail, is_active,
is_in_sitemap, no_follow, no_archive, created_at, updated_at, deleted_at)
VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?);

-- name: LookupPageURL :one
SELECT id, url FROM pages WHERE url = ?;
