CREATE TABLE IF NOT EXISTS "comments" (
  id                BIGSERIAL                 PRIMARY KEY,
  message           text                      NOT NULL,
  document_id       BIGINT                    NOT NULL,
  group_id          integer                   NOT NULL,
  created_at        timestamp with time zone  NOT NULL    DEFAULT NOW(),
  updated_at        timestamp with time zone  NOT NULL    DEFAULT NOW()
);

ALTER TABLE "comments"
ADD CONSTRAINT fk_comments_documents
FOREIGN KEY (document_id)
REFERENCES "documents"(id)
ON DELETE CASCADE;

CREATE INDEX "idx_document_id" ON "comments" (document_id);