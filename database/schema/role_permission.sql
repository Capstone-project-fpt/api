CREATE TABLE IF NOT EXISTS "roles_permissions" (
  id              BIGSERIAL PRIMARY KEY,
  role_id         BIGINT    NOT NULL,
  permission_id   BIGINT    NOT NULL,
  created_at      timestamp with time zone  NOT NULL    DEFAULT NOW(),
  updated_at      timestamp with time zone  NOT NULL    DEFAULT NOW(),
  UNIQUE (role_id, permission_id)
);