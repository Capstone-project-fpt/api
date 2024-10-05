CREATE TABLE IF NOT EXISTS "topic_references" (
  id                BIGSERIAL                 PRIMARY KEY,
  name              TEXT                      NOT NULL,
  path              TEXT                      NOT NULL,
  teacher_id        BIGINT                    NOT NULL,
  created_at        timestamp with time zone  NOT NULL    DEFAULT NOW(),
  updated_at        timestamp with time zone  NOT NULL    DEFAULT NOW()
);

ALTER TABLE "topic_references"
ADD CONSTRAINT fk_topic_references_teacher
FOREIGN KEY (teacher_id) REFERENCES teachers(id);

CREATE INDEX "idx_topic_references_teacher_id" ON "topic_references" (teacher_id);