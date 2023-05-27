ALTER TABLE
    "user_tbl"
ADD
    COLUMN IF NOT EXISTS "department" BIGINT;

ALTER TABLE
    "stage_tbl" DROP COLUMN IF EXISTS "faculty_id";

CREATE TABLE IF NOT EXISTS "stage_detail_tbl" (
    "id" BIGINT PRIMARY KEY,
    "description" VARCHAR(255),
    "url" VARCHAR(355),
    "stage_id" BIGINT NOT NULL,
    "department_id" BIGINT,
    "faculty_id" BIGINT,
    "time" BIGINT NOT NULL
);

ALTER TABLE"stage_tbl" 
DROP COLUMN IF EXISTS "time_start",
DROP COLUMN IF EXISTS "time_end",
ADD COLUMN IF NOT EXISTS "time_start" BIGINT,
ADD COLUMN IF NOT EXISTS "time_end" BIGINT;