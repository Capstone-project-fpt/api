CREATE TABLE IF NOT EXISTS "roles_permissions" (
  id              BIGSERIAL PRIMARY KEY,
  role_id         BIGINT    NOT NULL,
  permission_id   BIGINT    NOT NULL,
  UNIQUE (role_id, permission_id)
);
