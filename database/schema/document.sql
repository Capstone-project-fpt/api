CREATE TABLE IF NOT EXISTS "documents" (
  id                BIGSERIAL                 PRIMARY KEY,
  name              text                      NOT NULL,
  file_ids          text[]                     NULL,
  capstone_group_id BIGINT                    NOT NULL,
  score             SMALLINT                  NULL,
  created_at        timestamp with time zone  NOT NULL    DEFAULT NOW(),
  updated_at        timestamp with time zone  NOT NULL    DEFAULT NOW()
);
