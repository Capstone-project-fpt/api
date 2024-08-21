CREATE TABLE IF NOT EXISTS "majors" (
  id          BIGSERIAL                 PRIMARY KEY,
  name        text                      NOT NULL,
  created_at  timestamp with time zone  NOT NULL    DEFAULT NOW(),
  updated_at  timestamp with time zone  NOT NULL    DEFAULT NOW()
)