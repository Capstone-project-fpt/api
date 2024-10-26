CREATE TABLE IF NOT EXISTS "invitation_mentor_capstone_groups" (
  id                      BIGSERIAL                 PRIMARY KEY,
  status                  varchar(50)               NOT NULL,
  mentor_id               BIGINT                    NOT NULL,
  capstone_group_id       BIGINT                    NOT NULL,
  expired_at              timestamp with time zone  NOT NULL,
  created_at              timestamp with time zone  NOT NULL    DEFAULT NOW(),
  updated_at              timestamp with time zone  NOT NULL    DEFAULT NOW()
);

ALTER TABLE "invitation_mentor_capstone_groups"
ADD CONSTRAINT fk_invitation_mentor_capstone_groups_mentors
FOREIGN KEY (mentor_id) REFERENCES teachers(id);

ALTER TABLE "invitation_mentor_capstone_groups"
ADD CONSTRAINT fk_invitation_mentor_capstone_groups_capstone_groups
FOREIGN KEY (capstone_group_id) REFERENCES capstone_groups(id);

CREATE INDEX "idx_invitation_mentor_capstone_groups_mentor_id" ON "invitation_mentor_capstone_groups" (mentor_id);
CREATE INDEX "idx_invitation_mentor_capstone_groups_capstone_group_id" ON "invitation_mentor_capstone_groups" (capstone_group_id);