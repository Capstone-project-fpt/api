CREATE TABLE IF NOT EXISTS "student_capstone_groups" (
  id                      BIGSERIAL                 PRIMARY KEY,
  student_id              BIGINT                    NOT NULL,
  capstone_group_id       BIGINT                    NOT NULL,
  semester_id             BIGINT                    NOT NULL,
  created_at              timestamp with time zone  NOT NULL    DEFAULT NOW(),
  updated_at              timestamp with time zone  NOT NULL    DEFAULT NOW()
);

ALTER TABLE "student_capstone_groups"
ADD CONSTRAINT fk_student_capstone_groups_students
FOREIGN KEY (student_id) REFERENCES students(id);

ALTER TABLE "student_capstone_groups"
ADD CONSTRAINT fk_student_capstone_groups_capstone_groups
FOREIGN KEY (capstone_group_id) REFERENCES capstone_groups(id);

ALTER TABLE "student_capstone_groups"
ADD CONSTRAINT fk_student_capstone_groups_semesters
FOREIGN KEY (semester_id) REFERENCES semesters(id);

CREATE INDEX "idx_student_capstone_groups_student_id" ON "student_capstone_groups" (student_id);
CREATE INDEX "idx_student_capstone_groups_capstone_group_id" ON "student_capstone_groups" (capstone_group_id);
CREATE INDEX "idx_student_capstone_groups_semester_id" ON "student_capstone_groups" (semester_id);
