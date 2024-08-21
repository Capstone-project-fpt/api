CREATE TABLE IF NOT EXISTS "capstone_groups" (
  id          BIGSERIAL                 PRIMARY KEY,
  name_group  text                      NOT NULL,
  topic       text                      NOT NULL,
  major_id    BIGINT                    NOT NULL,
  semester_id BIGINT                    NOT NULL,
  created_at  timestamp with time zone  NOT NULL    DEFAULT NOW(),
  updated_at  timestamp with time zone  NOT NULL    DEFAULT NOW()
);