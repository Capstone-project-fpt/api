CREATE TABLE IF NOT EXISTS "documents" (
  id                BIGSERIAL                 PRIMARY KEY,
  name              text                      NOT NULL,
  file_ids          text[]                     NULL,
  capstone_group_id BIGINT                    NOT NULL,
  score             SMALLINT                  NULL,
  created_at        timestamp with time zone  NOT NULL    DEFAULT NOW(),
  updated_at        timestamp with time zone  NOT NULL    DEFAULT NOW()
);

ALTER TABLE "documents"
ADD CONSTRAINT fk_documents_capstone_groups
FOREIGN KEY (capstone_group_id) REFERENCES capstone_groups(id)
ON DELETE CASCADE;

CREATE INDEX "idx_capstone_id" ON "documents" (capstone_group_id);