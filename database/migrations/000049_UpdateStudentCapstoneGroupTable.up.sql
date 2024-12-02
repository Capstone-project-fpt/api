ALTER TABLE "student_capstone_groups"
ADD COLUMN IF NOT EXISTS "score" NUMERIC NULL;

ALTER TABLE "student_capstone_groups"
ADD COLUMN IF NOT EXISTS "status" VARCHAR(50) NULL;