CREATE TABLE animeSearchCache (
  id   BIGSERIAL PRIMARY KEY,
  searchTerm text NOT NULL,
  response jsonb NOT NULL
);

CREATE TABLE animeResultCache (
  id integer PRIMARY KEY,
  response jsonb NOT NULL
);
