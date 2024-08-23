CREATE TABLE IF NOT EXISTS "sub_majors" (
  id          BIGSERIAL                 PRIMARY KEY,
  name        text                      NOT NULL,
  major_id    BIGINT                    NOT NULL,
  created_at  timestamp with time zone  NOT NULL    DEFAULT NOW(),
  updated_at  timestamp with time zone  NOT NULL    DEFAULT NOW()
);