CREATE TABLE IF NOT EXISTS "comments" (
  id                BIGSERIAL                 PRIMARY KEY,
  message           text                      NOT NULL,
  document_id       BIGINT                    NOT NULL,
  group_id          integer                   NOT NULL,
  created_at        timestamp with time zone  NOT NULL    DEFAULT NOW(),
  updated_at        timestamp with time zone  NOT NULL    DEFAULT NOW()
);