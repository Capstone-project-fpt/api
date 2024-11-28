CREATE TABLE IF NOT EXISTS "report_comments" (
  id                      BIGSERIAL                 PRIMARY KEY,
  user_id                 BIGINT                    NOT NULL,
  report_document_id      BIGINT                    NOT NULL,
  message                 TEXT                      NOT NULL,
  group_comment           INTEGER                   NOT NULL,
  created_at              timestamp with time zone  NOT NULL    DEFAULT NOW(),
  updated_at              timestamp with time zone  NOT NULL    DEFAULT NOW()
);

ALTER TABLE "report_comments"
ADD CONSTRAINT fk_report_comments_users
FOREIGN KEY (user_id) REFERENCES users(id);

ALTER TABLE "report_comments"
ADD CONSTRAINT fk_report_comments_report_documents
FOREIGN KEY (report_document_id) REFERENCES report_documents(id);

CREATE INDEX "idx_report_comments_user_id" ON "report_comments" (user_id);
CREATE INDEX "idx_report_comments_report_document_id" ON "report_comments" (report_document_id);