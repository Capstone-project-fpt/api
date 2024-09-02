CREATE TABLE IF NOT EXISTS "students" (
  id                BIGSERIAL                 PRIMARY KEY,
  code              text                      NOT NULL,
  sub_major_id      BIGINT                    NOT NULL,
  user_id           BIGINT                    NOT NULL,
  capstone_group_id BIGINT                    NULL,
  created_at        timestamp with time zone  NOT NULL    DEFAULT NOW(),
  updated_at        timestamp with time zone  NOT NULL    DEFAULT NOW()
);

ALTER TABLE "students"
ADD CONSTRAINT fk_students_sub_majors
FOREIGN KEY (sub_major_id) REFERENCES sub_majors(id)
ON DELETE CASCADE;

ALTER TABLE "students"
ADD CONSTRAINT fk_students_users
FOREIGN KEY (user_id) REFERENCES users(id)
ON DELETE CASCADE;

ALTER TABLE "students"
ADD CONSTRAINT fk_students_capstone_groups
FOREIGN KEY (capstone_group_id) REFERENCES capstone_groups(id)
ON DELETE CASCADE;

CREATE INDEX "idx_sub_major_id" ON "students" (sub_major_id);
CREATE INDEX "idx_user_id" ON "students" (user_id);
CREATE INDEX "idx_capstone_group_id" ON "students" (capstone_group_id);