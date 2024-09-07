CREATE TABLE IF NOT EXISTS "capstone_groups" (
  id          BIGSERIAL                 PRIMARY KEY,
  name_group  text                      NOT NULL,
  topic       text                      NOT NULL,
  major_id    BIGINT                    NOT NULL,
  semester_id BIGINT                    NOT NULL,
  created_at  timestamp with time zone  NOT NULL    DEFAULT NOW(),
  updated_at  timestamp with time zone  NOT NULL    DEFAULT NOW()
);

ALTER TABLE "capstone_groups"
ADD CONSTRAINT fk_capstone_group_majors
FOREIGN KEY (major_id) REFERENCES majors(id)
ON DELETE CASCADE;

ALTER TABLE "capstone_groups"
ADD CONSTRAINT fk_capstone_group_semesters
FOREIGN KEY (semester_id) REFERENCES semesters(id)
ON DELETE CASCADE;

CREATE INDEX "idx_capstone_groups_major_id" ON "capstone_groups" (major_id);
CREATE INDEX "idx_capstone_groups_semester_id" ON "capstone_groups" (semester_id);

