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

-- name: CacheAnimeSearch :one
INSERT INTO animeSearchCache (
  searchTerm, response
) VALUES (
  $1, $2
)
ON CONFLICT(id) DO UPDATE SET response = EXCLUDED.response
RETURNING *;

-- name: GetCachedAnimeSearchResult :one
SELECT response FROM animeSearchCache WHERE searchTerm = $1;

-- name: SaveMapping :exec
INSERT INTO animeMapping (
  anilistId,
  plexSeriesId
) VALUES (
  $1, $2
);

-- name: GetMappings :many
SELECT * FROM animeMapping WHERE plexSeriesId = $1;
