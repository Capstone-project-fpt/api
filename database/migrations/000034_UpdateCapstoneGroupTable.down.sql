ALTER TABLE "capstone_groups" ADD COLUMN IF NOT EXISTS "topic" TEXT;
ALTER TABLE "capstone_groups" DROP COLUMN IF EXISTS "topic_id";
ALTER TABLE "capstone_groups" DROP COLUMN IF EXISTS "status";
DROP INDEX IF EXISTS "fk_capstone_groups_topic";