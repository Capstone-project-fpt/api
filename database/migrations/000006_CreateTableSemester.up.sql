CREATE TABLE IF NOT EXISTS "semesters" (
  id          BIGSERIAL                 PRIMARY KEY,
  name        text                      NOT NULL,
  start_time  timestamp with time zone  NOT NULL,
  end_time    timestamp with time zone  NOT NULL,
  created_at  timestamp with time zone  NOT NULL    DEFAULT NOW(),
  updated_at  timestamp with time zone  NOT NULL    DEFAULT NOW()
);
