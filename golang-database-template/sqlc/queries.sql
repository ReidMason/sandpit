 -- name: GetAuthor :one
SELECT * FROM authors
WHERE id = $1 LIMIT 1;

-- name: ListAuthors :many
SELECT * FROM authors
ORDER BY name;

-- name: CreateAuthor :one
INSERT INTO authors (
  id, name, bio
) VALUES (
  $1, $2, $3
)
RETURNING *;

-- name: DeleteAuthor :exec
DELETE FROM authors
WHERE id = $1;

-- name: GetAuthorsWithBooks :many
SELECT * FROM authors
LEFT JOIN books on books.authorid = authors.id
WHERE authors.id = ANY(sqlc.arg(ids)::int[]);

-- name: testing :many
SELECT * FROM authors WHERE id = $1;
