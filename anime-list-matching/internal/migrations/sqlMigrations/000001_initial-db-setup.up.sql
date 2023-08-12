CREATE TABLE animeSearchCache (
  id   BIGSERIAL PRIMARY KEY,
  searchTerm text NOT NULL,
  response json NOT NULL
);

CREATE TABLE animeResultCache (
  id integer PRIMARY KEY,
  response json NOT NULL
);
