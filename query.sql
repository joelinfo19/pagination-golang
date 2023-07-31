-- name: GetPeople :many
SELECT id, first_name, last_name FROM people WHERE id > ? ORDER BY id LIMIT ?