CREATE TABLE IF NOT EXISTS "files" (
  id                BIGSERIAL                 PRIMARY KEY,
  path              TEXT                      NOT NULL,
  created_at        timestamp with time zone  NOT NULL    DEFAULT NOW(),
  updated_at        timestamp with time zone  NOT NULL    DEFAULT NOW()
);