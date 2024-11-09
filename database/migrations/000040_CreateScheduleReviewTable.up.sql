CREATE TABLE IF NOT EXISTS "schedule_reviews" (
  id                      BIGSERIAL                 PRIMARY KEY,
  title                   VARCHAR(100)              NOT NULL,
  description             TEXT                      NOT NULL,
  link_meeting            TEXT                      NOT NULL,
  start_time              timestamp with time zone  NOT NULL,
  end_time                timestamp with time zone  NOT NULL,
  evaluation_committee_id BIGINT                    NOT NULL,
  capstone_group_id       BIGINT                    NOT NULL,
  created_at              timestamp with time zone  NOT NULL    DEFAULT NOW(),
  updated_at              timestamp with time zone  NOT NULL    DEFAULT NOW()
);

ALTER TABLE "schedule_reviews"
ADD CONSTRAINT fk_schedule_reviews_evaluation_committees
FOREIGN KEY (evaluation_committee_id) REFERENCES evaluation_committees(id);

ALTER TABLE "schedule_reviews"
ADD CONSTRAINT fk_schedule_reviews_capstone_groups
FOREIGN KEY (capstone_group_id) REFERENCES capstone_groups(id);

CREATE INDEX "idx_schedule_reviews_evaluation_committee_id" ON "schedule_reviews" (evaluation_committee_id);
CREATE INDEX "idx_schedule_reviews_capstone_group_id" ON "schedule_reviews" (capstone_group_id);