CREATE TABLE IF NOT EXISTS "capstone_group_topics" (
  id                BIGSERIAL                 PRIMARY KEY,
  topic             VARCHAR(255),
  document_path     VARCHAR(255),
  status_review     VARCHAR(50),
  approved_at       timestamp with time zone  NULL,
  approved_by_id    BIGINT                    NULL,
  rejected_at       timestamp with time zone  NULL,
  rejected_by_id    BIGINT                    NULL,
  created_at        timestamp with time zone  NOT NULL    DEFAULT NOW(),
  updated_at        timestamp with time zone  NOT NULL    DEFAULT NOW(),
  capstone_group_id BIGINT                    NOT NULL
);

ALTER TABLE "capstone_group_topics"
ADD CONSTRAINT fk_capstone_group_topics_capstone_groups
FOREIGN KEY (capstone_group_id) REFERENCES capstone_groups(id);

ALTER TABLE "capstone_group_topics"
ADD CONSTRAINT fk_capstone_group_topics_approved_by
FOREIGN KEY (approved_by_id) REFERENCES teachers(id);

ALTER TABLE "capstone_group_topics"
ADD CONSTRAINT fk_capstone_group_topics_rejected_by
FOREIGN KEY (rejected_by_id) REFERENCES teachers(id);

CREATE INDEX "idx_capstone_group_topics_capstone_group_id" ON "capstone_group_topics" (capstone_group_id);
CREATE INDEX "idx_capstone_group_topics_approved_by_id" ON "capstone_group_topics" (approved_by_id);
CREATE INDEX "idx_capstone_group_topics_rejected_by_id" ON "capstone_group_topics" (rejected_by_id);