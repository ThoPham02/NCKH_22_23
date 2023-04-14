CREATE TABLE "user" (
  "id" INT PRIMARY KEY,
  "name" VARCHAR(255) NOT NULL,
  "hash_password" VARCHAR(255) NOT NULL,
  "email" VARCHAR(255) NOT NULL,
  "type_account" INT NOT NULL
);

CREATE TABLE "notification" (
  "id" INT PRIMARY KEY,
  "name" VARCHAR(255) NOT NULL,
  "url" VARCHAR(255)
);

CREATE TABLE "library" (
  "id" INT PRIMARY KEY,
  "name" VARCHAR(255) NOT NULL,
  "url" VARCHAR(255),
  "owner_id" INT NOT NULL
);

CREATE TABLE "user_info" (
  "id" INT PRIMARY KEY,
  "user_id" INT NOT NULL,
  "name" VARCHAR(255) NOT NULL,
  "email" VARCHAR(255) NOT NULL,
  "phone" VARCHAR(20) NOT NULL,
  "faculty_id" INT NOT NULL,
  "degree" INT NOT NULL,
  "year_start" INT NOT NULL,
  "avata_url" VARCHAR(255),
  "birthday" VARCHAR(255),
  "bank_account" VARCHAR(20)
);

CREATE TABLE "department" (
  "id" INT PRIMARY KEY,
  "name" VARCHAR(255) NOT NULL,
  "faculty_id" INT NOT NULL
);

CREATE TABLE "faculty" (
  "id" INT PRIMARY KEY,
  "name" VARCHAR(255) NOT NULL
);

CREATE TABLE "topic_registration" (
  "id" INT PRIMARY KEY,
  "name" VARCHAR(255) NOT NULL,
  "lecture_id" INT NOT NULL,
  "faculty_id" INT NOT NULL,
  "status" INT,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "topic" (
  "id" INT PRIMARY KEY,
  "name" VARCHAR(255) NOT NULL,
  "lecture_id" INT NOT NULL,
  "faculty_id" INT NOT NULL,
  "status" INT NOT NULL DEFAULT 1,
  "result_url" VARCHAR(255),
  "conference_id" INT NOT NULL,
  "group_id" INT,
  "time_start" timestamptz NOT NULL DEFAULT (now()),
  "time_end" timestamptz NOT NULL
);

CREATE TABLE "student_topic" (
  "id" INT PRIMARY KEY,
  "student_id" INT NOT NULL,
  "topic_id" INT NOT NULL
);

CREATE TABLE "conference" (
  "id" INT PRIMARY KEY,
  "name" VARCHAR(255) NOT NULL,
  "cash_support" INT,
  "school_year" VARCHAR(20)
);

CREATE TABLE "group" (
  "id" INT PRIMARY KEY,
  "name" VARCHAR(255) NOT NULL,
  "conference_id" INT NOT NULL,
  "faculty_id" INT NOT NULL
);

CREATE TABLE "group_supervisor" (
  "id" INT PRIMARY KEY,
  "supervisor_id" INT NOT NULL,
  "role" INT NOT NULL,
  "group_report_id" INT NOT NULL
);

CREATE TABLE "topic_result" (
  "id" INT PRIMARY KEY,
  "score" float NOT NULL DEFAULT 0,
  "comment" VARCHAR(255) NOT NULL,
  "topic_id" INT NOT NULL,
  "supervisor_id" INT NOT NULL
);
