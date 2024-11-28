CREATE TABLE IF NOT EXISTS "verifier_topics" (
  id                      BIGSERIAL                 PRIMARY KEY,
  teacher_id              BIGINT                    NOT NULL,
  semester_id             BIGINT                    NOT NULL,
  created_at              timestamp with time zone  NOT NULL    DEFAULT NOW(),
  updated_at              timestamp with time zone  NOT NULL    DEFAULT NOW()
);

ALTER TABLE "verifier_topics"
ADD CONSTRAINT fk_verifier_topics_teachers
FOREIGN KEY (teacher_id) REFERENCES teachers(id);

ALTER TABLE "verifier_topics"
ADD CONSTRAINT fk_verifier_topics_semesters
FOREIGN KEY (semester_id) REFERENCES semesters(id);

CREATE INDEX "idx_verifier_topics_teacher_id" ON "verifier_topics" (teacher_id);
CREATE INDEX "idx_verifier_topics_semester_id" ON "verifier_topics" (semester_id);
