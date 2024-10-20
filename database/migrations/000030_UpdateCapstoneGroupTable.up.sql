ALTER TABLE "capstone_groups" ADD COLUMN IF NOT EXISTS "mentor_id" BIGINT;

ALTER TABLE "capstone_groups"
ADD CONSTRAINT fk_capstone_groups_mentor
FOREIGN KEY (mentor_id)
REFERENCES "teachers"(id);

CREATE INDEX "idx_capstone_groups_mentor_id" ON "capstone_groups" (mentor_id);