CREATE TABLE IF NOT EXISTS "report_documents" (
  id                      BIGSERIAL                 PRIMARY KEY,
  name                    TEXT                      NOT NULL,
  file_ids                TEXT[]                    NOT NULL,
  capstone_group_id       BIGINT                    NOT NULL,
  mentor_review_status    VARCHAR(50)               NOT NULL,
  type_report             VARCHAR(50)               NOT NULL,
  conclusion              TEXT                      NULL,
  created_at              timestamp with time zone  NOT NULL    DEFAULT NOW(),
  updated_at              timestamp with time zone  NOT NULL    DEFAULT NOW()
);

ALTER TABLE "report_documents"
ADD CONSTRAINT fk_report_documents_capstone_groups
FOREIGN KEY (capstone_group_id) REFERENCES capstone_groups(id);

CREATE INDEX "idx_report_documents_capstone_group_id" ON "report_documents" (capstone_group_id);