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

-- name: GetAuthorById :one
SELECT name FROM authors
WHERE id = $1 LIMIT 1;

-- name: GetAuthorsWithBooks :many
SELECT * FROM authors
INNER JOIN books on books.authorid = authors.id
-- For Databases other than Postgres
-- WHERE authors.id in (sqlc.slice('ids'));
WHERE authors.id = ANY($1::int[]);
