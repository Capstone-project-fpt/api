CREATE TABLE IF NOT EXISTS "users" (
  id                BIGSERIAL                 PRIMARY KEY,
  name              text                      NOT NULL,
  user_type         text                      NOT NULL,
  password          text                      NULL,
  email             text                      NOT NULL,
  code              text                      NULL,
  sub_major_id      BIGINT                    NULL, 
  capstone_group_id BIGINT                    NULL,
  created_at        timestamp with time zone  NOT NULL    DEFAULT NOW(),
  updated_at        timestamp with time zone  NOT NULL    DEFAULT NOW()
);

ALTER TABLE "users"
ADD CONSTRAINT fk_users_sub_majors
FOREIGN KEY (sub_major_id) REFERENCES sub_majors(id)
ON DELETE CASCADE;

ALTER TABLE "users"
ADD CONSTRAINT fk_users_capstone_groups
FOREIGN KEY (capstone_group_id) REFERENCES capstone_groups(id)
ON DELETE CASCADE;

CREATE INDEX "idx_sub_major_id" ON "users" (sub_major_id);
CREATE INDEX "idx_capstone_group_id" ON "users" (capstone_group_id);