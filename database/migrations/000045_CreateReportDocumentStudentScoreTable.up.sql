DELETE FROM "report_documents";

CREATE TABLE IF NOT EXISTS "report_document_student_scores" (
  id                      BIGSERIAL                 PRIMARY KEY,
  report_document_id      BIGINT                    NOT NULL,
  student_id              BIGINT                    NOT NULL,
  score                   NUMERIC                   NULL,
  created_at              timestamp with time zone  NOT NULL    DEFAULT NOW(),
  updated_at              timestamp with time zone  NOT NULL    DEFAULT NOW()
);

ALTER TABLE "report_document_student_scores"
ADD CONSTRAINT fk_report_document_student_scores_report_documents
FOREIGN KEY (report_document_id) REFERENCES report_documents(id);

ALTER TABLE "report_document_student_scores"
ADD CONSTRAINT fk_report_document_student_scores_students
FOREIGN KEY (student_id) REFERENCES students(id);

CREATE INDEX "idx_report_document_student_scores_report_document_id" ON "report_document_student_scores" (report_document_id);
CREATE INDEX "idx_report_document_student_scores_student_id" ON "report_document_student_scores" (student_id);