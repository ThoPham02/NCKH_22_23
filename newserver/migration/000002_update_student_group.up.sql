ALTER TABLE "topic_tbl" DROP COLUMN IF EXISTS "list_students";
ALTER TABLE "topic_tbl" ADD COLUMN IF NOT EXISTS "group_students_id" BIGINT;

CREATE TABLE IF NOT EXISTS "student_group_tbl" (
  "id" BIGINT PRIMARY KEY,
  "event_id" BIGINT NOT NULL,
  "student_id" BIGINT NOT NULL,
  "group_id" BIGINT NOT NULL
);