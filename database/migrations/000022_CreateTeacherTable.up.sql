CREATE TABLE IF NOT EXISTS "teachers" (
  id                BIGSERIAL                 PRIMARY KEY,
  user_id           BIGINT                    NOT NULL,
  sub_major_id      BIGINT                    NOT NULL,
  created_at        timestamp with time zone  NOT NULL    DEFAULT NOW(),
  updated_at        timestamp with time zone  NOT NULL    DEFAULT NOW()
);

ALTER TABLE "teachers"
ADD CONSTRAINT fk_teachers_sub_majors
FOREIGN KEY (sub_major_id) REFERENCES sub_majors(id);

ALTER TABLE "teachers"
ADD CONSTRAINT fk_teachers_users
FOREIGN KEY (user_id) REFERENCES users(id);

CREATE INDEX "idx_teachers_sub_major_id" ON "teachers" (sub_major_id);
CREATE INDEX "idx_teachers_user_id" ON "teachers" (user_id);