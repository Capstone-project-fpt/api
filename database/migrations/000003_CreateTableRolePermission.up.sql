CREATE TABLE IF NOT EXISTS "roles_permissions" (
  id              BIGSERIAL PRIMARY KEY,
  role_id         BIGINT    NOT NULL,
  permission_id   BIGINT    NOT NULL,
  UNIQUE (role_id, permission_id)
);

ALTER TABLE "roles_permissions"
ADD CONSTRAINT fk_roles_permissions_roles
FOREIGN KEY (role_id)
REFERENCES "roles"(id)
ON DELETE CASCADE;

ALTER TABLE "roles_permissions"
ADD CONSTRAINT fk_roles_permissions_permissions
FOREIGN KEY (permission_id)
REFERENCES "permissions"(id)
ON DELETE CASCADE;

CREATE INDEX "idx_role_id_permission_id"" ON "roles_permissions" (role_id, permission_id)