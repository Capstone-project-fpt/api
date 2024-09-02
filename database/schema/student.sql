CREATE TABLE IF NOT EXISTS "students" (
  id                BIGSERIAL                 PRIMARY KEY,
  code              text                      NOT NULL,
  sub_major_id      BIGINT                    NOT NULL,
  user_id           BIGINT                    NOT NULL,
  capstone_group_id BIGINT                    NULL,
  created_at        timestamp with time zone  NOT NULL    DEFAULT NOW(),
  updated_at        timestamp with time zone  NOT NULL    DEFAULT NOW()
);