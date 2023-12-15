-- migrate:up
CREATE TABLE article_hash (
  article_hash TEXT PRIMARY KEY NOT NULL,
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- migrate:down
DROP TABLE article_hash;