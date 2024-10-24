ALTER TABLE "capstone_groups" DROP COLUMN IF EXISTS "topic";

ALTER TABLE "capstone_groups" ADD COLUMN IF NOT EXISTS "topic_id" BIGINT;

ALTER TABLE "capstone_groups"
ADD CONSTRAINT fk_capstone_groups_topic
FOREIGN KEY (topic_id) REFERENCES capstone_group_topics(id);

ALTER TABLE "capstone_groups" ADD COLUMN IF NOT EXISTS "status" VARCHAR(50);

UPDATE "capstone_groups" SET "status" = 'reviewing_topic';