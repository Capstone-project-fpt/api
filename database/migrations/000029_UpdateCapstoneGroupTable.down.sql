ALTER TABLE "capstone_groups"
ALTER COLUMN "topic"
SET
  NOT NULL;

ALTER TABLE "capstone_groups"
DROP COLUMN IF EXISTS "leader_id";