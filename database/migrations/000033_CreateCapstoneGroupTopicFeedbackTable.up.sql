CREATE TABLE IF NOT EXISTS "capstone_group_topic_feedbacks" (
  id                      BIGSERIAL                 PRIMARY KEY,
  feedback                VARCHAR(255)              NOT NULL,
  reviewer_id             BIGINT                    NOT NULL,
  created_at              timestamp with time zone  NOT NULL    DEFAULT NOW(),
  updated_at              timestamp with time zone  NOT NULL    DEFAULT NOW(),
  capstone_group_topic_id BIGINT                    NOT NULL
);

ALTER TABLE "capstone_group_topic_feedbacks"
ADD CONSTRAINT fk_capstone_group_topic_feedbacks_capstone_groups_topics
FOREIGN KEY (capstone_group_topic_id) REFERENCES capstone_group_topics(id);

ALTER TABLE "capstone_group_topic_feedbacks"
ADD CONSTRAINT fk_capstone_group_topic_feedbacks_reviewers
FOREIGN KEY (reviewer_id) REFERENCES teachers(id);

CREATE INDEX "idx_capstone_group_topic_feedbacks_capstone_group_topic_id" ON "capstone_group_topic_feedbacks" (capstone_group_topic_id);
CREATE INDEX "idx_capstone_group_topic_feedbacks_reviewer_id" ON "capstone_group_topic_feedbacks" (reviewer_id);