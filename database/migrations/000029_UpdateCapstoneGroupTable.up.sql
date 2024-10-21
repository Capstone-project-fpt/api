ALTER TABLE "capstone_groups" ALTER COLUMN "topic" DROP NOT NULL;

ALTER TABLE "capstone_groups" ADD COLUMN IF NOT EXISTS "leader_id" BIGINT;

ALTER TABLE "capstone_groups"
ADD CONSTRAINT fk_capstone_groups_leader
FOREIGN KEY (leader_id)
REFERENCES "students"(id);

CREATE INDEX "idx_capstone_groups_leader_id" ON "capstone_groups" (leader_id);