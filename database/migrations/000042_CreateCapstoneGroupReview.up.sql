CREATE TABLE IF NOT EXISTS "capstone_group_reviews" (
  id                      BIGSERIAL                 PRIMARY KEY,
  capstone_group_id       BIGINT                    NOT NULL,
  schedule_review_id      BIGINT                    NOT NULL,
  report_files            TEXT[]                    NULL,
  feedback                TEXT                      NULL,
  created_at              timestamp with time zone  NOT NULL    DEFAULT NOW(),
  updated_at              timestamp with time zone  NOT NULL    DEFAULT NOW()
);

ALTER TABLE "capstone_group_reviews"
ADD CONSTRAINT fk_capstone_group_reviews_capstone_groups
FOREIGN KEY (capstone_group_id) REFERENCES capstone_groups(id);

ALTER TABLE "capstone_group_reviews"
ADD CONSTRAINT fk_capstone_group_reviews_schedule_reviews
FOREIGN KEY (schedule_review_id) REFERENCES schedule_reviews(id);

CREATE INDEX "idx_capstone_group_reviews_capstone_group_id" ON "capstone_group_reviews" (capstone_group_id);
CREATE INDEX "idx_capstone_group_reviews_schedule_review_id" ON "capstone_group_reviews" (schedule_review_id);