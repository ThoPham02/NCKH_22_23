CREATE TABLE "user_tbl" (
  "id" BIGINT PRIMARY KEY,
  "username" VARCHAR(255) NOT NULL,
  "hash_password" VARCHAR(255) NOT NULL,
  "role" BIGINT NOT NULL,
  "name" VARCHAR(255) NOT NULL,
  "email" VARCHAR(255),
  "phone" VARCHAR(20),
  "faculty_id" BIGINT NOT NULL,
  "year_start" BIGINT NOT NULL,
  "degree" BIGINT NOT NULL,
  "avata_url" VARCHAR(255),
  "birthday" VARCHAR(255),
  "bank_account" VARCHAR(20)
);

CREATE TABLE "notification_tbl" (
  "id" BIGINT PRIMARY KEY,
  "name" VARCHAR(255) NOT NULL,
  "url" VARCHAR(255)
);

CREATE TABLE "library_tbl" (
  "id" BIGINT PRIMARY KEY,
  "name" VARCHAR(255) NOT NULL,
  "url" VARCHAR(255),
  "owner_id" BIGINT NOT NULL
);

CREATE TABLE "department_tbl" (
  "id" BIGINT PRIMARY KEY,
  "name" VARCHAR(255) NOT NULL,
  "faculty_id" BIGINT NOT NULL
);

CREATE TABLE "faculty_tbl" (
  "id" BIGINT PRIMARY KEY,
  "name" VARCHAR(255) NOT NULL
);

CREATE TABLE "topic_tbl" (
  "id" BIGINT PRIMARY KEY,
  "name" VARCHAR(255) NOT NULL,
  "lecture_id" BIGINT NOT NULL,
  "department_id" BIGINT NOT NULL,
  "status" BIGINT NOT NULL,
  "event_id" BIGINT NOT NULL,
  "subcommittee_id" BIGINT,
  "list_students" VARCHAR(255),
  "time_start" BIGINT,
  "time_end" BIGINT,
  "cash_support" BIGINT
);

CREATE TABLE "topic_result_tbl" (
  "id" BIGINT PRIMARY KEY,
  "topic_id" BIGINT NOT NULL,
  "description" VARCHAR(255),
  "result_url" VARCHAR(255) NOT NULL
);

CREATE TABLE "event_tbl" (
  "id" BIGINT PRIMARY KEY,
  "name" VARCHAR(255) NOT NULL,
  "school_year" VARCHAR(20),
  "is_current" BIGINT
);

CREATE TABLE "stage_tbl" (
  "id" BIGINT PRIMARY KEY,
  "name" VARCHAR(255) NOT NULL,
  "description" VARCHAR(1024),
  "url" VARCHAR(255),
  "event_id" BIGINT NOT NULL,
  "time_start" BIGINT NOT NULL,
  "time_end" BIGINT NOT NULL
);

CREATE TABLE "subcommittee_tbl" (
  "id" BIGINT PRIMARY KEY,
  "name" VARCHAR(255),
  "facult_id" BIGINT NOT NULL
);

CREATE TABLE "group_tbl" (
  "id" BIGINT PRIMARY KEY,
  "subcommittee_id" BIGINT NOT NULL,
  "lecture_id" BIGINT NOT NULL,
  "role" BIGINT NOT NULL
);

CREATE TABLE "topic_mark_tbl" (
  "id" BIGINT PRIMARY KEY,
  "topic_id" BIGINT NOT NULL,
  "lecture_id" BIGINT NOT NULL,
  "point" float NOT NULL,
  "comment" VARCHAR(255),
  "url" VARCHAR(255)
);
