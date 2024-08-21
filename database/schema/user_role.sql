CREATE TABLE IF NOT EXISTS "users_roles" (
  id              BIGSERIAL PRIMARY KEY,
  role_id         BIGINT    NOT NULL,
  user_id         BIGINT    NOT NULL,
  created_at      timestamp with time zone  NOT NULL    DEFAULT NOW(),
  updated_at      timestamp with time zone  NOT NULL    DEFAULT NOW(),
  UNIQUE (role_id, user_id)
);
