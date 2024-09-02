CREATE TABLE IF NOT EXISTS "users" (
  id                BIGSERIAL                 PRIMARY KEY,
  name              text                      NOT NULL,
  user_type         text                      NOT NULL,
  password          text                      NULL,
  email             text                      NOT NULL,
  phone_number      text                      NOT NULL,
  created_at        timestamp with time zone  NOT NULL    DEFAULT NOW(),
  updated_at        timestamp with time zone  NOT NULL    DEFAULT NOW()
);