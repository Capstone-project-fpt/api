ALTER TABLE "users"
ADD COLUMN IF NOT EXISTS "phone_number" text NULL;

ALTER TABLE "users"
DROP COLUMN IF EXISTS "code",
DROP COLUMN IF EXISTS "sub_major_id",
DROP COLUMN IF EXISTS "capstone_group_id";