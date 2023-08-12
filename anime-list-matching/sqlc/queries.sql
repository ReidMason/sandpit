-- name: CacheAnimeResult :one
INSERT INTO animeResultCache (
  id, response
) VALUES (
  $1, $2
)
ON CONFLICT(id) DO UPDATE SET response = EXCLUDED.response
RETURNING *;

-- name: GetCachedAnimeResult :one
SELECT response FROM animeResultCache WHERE id = $1;

