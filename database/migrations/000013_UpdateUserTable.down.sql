ALTER TABLE "users"
DROP COLUMN IF EXISTS "phone_number";

ALTER TABLE "users"
ADD COLUMN IF NOT EXISTS "code" text null,
ADD COLUMN IF NOT EXISTS "sub_major_id" bigint null,
ADD COLUMN IF NOT EXISTS "capstone_group_id" bigint null;