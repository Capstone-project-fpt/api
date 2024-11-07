CREATE TABLE IF NOT EXISTS "evaluation_committees" (
  id                      BIGSERIAL                 PRIMARY KEY,
  name                    VARCHAR(255)              NOT NULL,
  semester_id             BIGINT                    NOT NULL,
  teacher_ids             JSONB                     NOT NULL,
  created_at              timestamp with time zone  NOT NULL    DEFAULT NOW(),
  updated_at              timestamp with time zone  NOT NULL    DEFAULT NOW()
);

ALTER TABLE "evaluation_committees"
ADD CONSTRAINT fk_evaluation_committees_semesters
FOREIGN KEY (semester_id) REFERENCES semesters(id);

CREATE INDEX "idx_evaluation_committees_semester_id" ON "evaluation_committees" (semester_id);