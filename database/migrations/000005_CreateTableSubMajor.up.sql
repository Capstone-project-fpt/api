CREATE TABLE IF NOT EXISTS "sub_majors" (
  id          BIGSERIAL                 PRIMARY KEY,
  name        text                      NOT NULL,
  major_id    BIGINT                    NOT NULL,
  created_at  timestamp with time zone  NOT NULL    DEFAULT NOW(),
  updated_at  timestamp with time zone  NOT NULL    DEFAULT NOW()
);

ALTER TABLE "sub_majors"
ADD CONSTRAINT fk_sub_majors_majors
FOREIGN KEY (major_id) REFERENCES majors(id)
ON DELETE CASCADE;

CREATE INDEX "idx_major_id" ON "sub_majors" (major_id);