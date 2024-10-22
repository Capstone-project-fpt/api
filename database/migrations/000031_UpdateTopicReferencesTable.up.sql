ALTER TABLE "topic_references" ADD COLUMN IF NOT EXISTS "status_review" Text;

UPDATE "topic_references" SET "status_review" = 'approved';