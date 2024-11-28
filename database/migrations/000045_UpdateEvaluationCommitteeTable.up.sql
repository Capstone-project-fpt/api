ALTER TABLE "evaluation_committees" ADD COLUMN IF NOT EXISTS "assign_group_ids" JSONB;

UPDATE "evaluation_committees" SET "assign_group_ids" = '[]' WHERE "assign_group_ids" IS NULL;

ALTER TABLE "evaluation_committees" ALTER COLUMN "assign_group_ids" SET NOT NULL;
