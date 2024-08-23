CREATE TABLE IF NOT EXISTS "users_roles" (
  id              BIGSERIAL PRIMARY KEY,
  role_id         BIGINT    NOT NULL,
  user_id         BIGINT    NOT NULL,
  created_at      timestamp with time zone  NOT NULL    DEFAULT NOW(),
  updated_at      timestamp with time zone  NOT NULL    DEFAULT NOW(),
  UNIQUE (role_id, user_id)
);

ALTER TABLE "users_roles"
ADD CONSTRAINT fk_users_roles_roles
FOREIGN KEY (role_id)
REFERENCES "roles"(id)
ON DELETE CASCADE;

ALTER TABLE "users_roles"
ADD CONSTRAINT fk_users_roles_users
FOREIGN KEY (user_id)
REFERENCES "users"(id)
ON DELETE CASCADE;

CREATE INDEX "idx_role_id_user_id" ON "users_roles" (role_id, user_id);