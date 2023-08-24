CREATE TABLE animeSearchCache (
  id   integer PRIMARY KEY,
  searchTerm text NOT NULL,
  response jsonb NOT NULL
);

CREATE TABLE animeResultCache (
  id integer PRIMARY KEY,
  response jsonb NOT NULL
);
