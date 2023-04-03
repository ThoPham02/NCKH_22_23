CREATE TABLE "user" (
  "id" BIGSERIAL PRIMARY KEY,
  "name" VARCHAR(255) NOT NULL,
  "hash_password" VARCHAR(255) NOT NULL,
  "type_account_id" BIGSERIAL NOT NULL,
  "email" VARCHAR(255) NOT NULL
);

CREATE TABLE "type_account" (
  "id" BIGSERIAL PRIMARY KEY,
  "name" VARCHAR(255)
);

CREATE TABLE "user_info" (
  "user_id" BIGSERIAL PRIMARY KEY,
  "name" VARCHAR(255) NOT NULL,
  "description" VARCHAR(255),
  "avata_url" VARCHAR(255),
  "birthday" date,
  "faculity_id" BIGSERIAL NOT NULL,
  "year_start" BIGSERIAL NOT NULL,
  "bank_account" VARCHAR(20),
  "phone" VARCHAR(20),
  "sex" BIGSERIAL
);

CREATE TABLE "notification" (
  "id" BIGSERIAL PRIMARY KEY,
  "name" VARCHAR(255) NOT NULL,
  "content" VARCHAR(255) NOT NULL,
  "from_user_id" BIGSERIAL NOT NULL,
  "to_user_id" BIGSERIAL,
  "created_at" timestamptz NOT NULL
);

CREATE TABLE "department" (
  "id" BIGSERIAL PRIMARY KEY,
  "name" VARCHAR(255) NOT NULL,
  "faculity_id" BIGSERIAL NOT NULL
);

CREATE TABLE "faculity" (
  "id" BIGSERIAL PRIMARY KEY,
  "name" VARCHAR(255) NOT NULL
);

CREATE TABLE "topic_registration" (
  "id" BIGSERIAL PRIMARY KEY,
  "name" VARCHAR(255) NOT NULL,
  "description" VARCHAR(255) NOT NULL,
  "description_url" VARCHAR(255),
  "lecture_id" BIGSERIAL NOT NULL,
  "faculity_id" BIGSERIAL,
  "created_at" timestamptz
);

CREATE TABLE "topic" (
  "id" BIGSERIAL PRIMARY KEY,
  "name" VARCHAR(255) NOT NULL,
  "description" VARCHAR(255) NOT NULL,
  "description_url" VARCHAR(255),
  "lecture_id" BIGSERIAL NOT NULL,
  "group_id" BIGSERIAL NOT NULL,
  "department_id" BIGSERIAL NOT NULL,
  "status_id" BIGSERIAL NOT NULL,
  "exam_id" BIGSERIAL NOT NULL,
  "topic_report_id" BIGSERIAL,
  "acceptance_board_id" BIGSERIAL,
  "time_start" timestamptz NOT NULL DEFAULT (now()),
  "time_end" timestamptz
);

CREATE TABLE "status" (
  "id" BIGSERIAL PRIMARY KEY,
  "name" VARCHAR(255) NOT NULL
);

CREATE TABLE "exam" (
  "id" BIGSERIAL PRIMARY KEY,
  "name" VARCHAR(20) NOT NULL,
  "year" INT NOT NULL,
  "estimated_cost" money
);

CREATE TABLE "group_student" (
  "id" BIGSERIAL PRIMARY KEY,
  "group_id" BIGSERIAL NOT NULL,
  "student_id" BIGSERIAL NOT NULL
);

CREATE TABLE "group_detail" (
  "id" BIGSERIAL PRIMARY KEY,
  "num_student" INT NOT NULL DEFAULT 1
);

CREATE TABLE "acceptance_board" (
  "id" BIGSERIAL PRIMARY KEY,
  "name" VARCHAR(255),
  "chairman" BIGSERIAL NOT NULL,
  "reviewer1" BIGSERIAL NOT NULL,
  "reviewer2" BIGSERIAL,
  "secretary" BIGSERIAL
);

CREATE TABLE "acceptance_report" (
  "id" BIGSERIAL PRIMARY KEY,
  "acceptance_board_id" BIGSERIAL NOT NULL,
  "topic_id" BIGSERIAL NOT NULL,
  "total_score" float NOT NULL DEFAULT 0,
  "comment" VARCHAR(255),
  "report_url" VARCHAR(255)
);

CREATE TABLE "topic_report" (
  "id" BIGSERIAL PRIMARY KEY,
  "summary_content" VARCHAR(255) NOT NULL,
  "content_url" VARCHAR(255)
);

ALTER TABLE "user" ADD FOREIGN KEY ("type_account_id") REFERENCES "type_account" ("id");

ALTER TABLE "user_info" ADD FOREIGN KEY ("user_id") REFERENCES "user" ("id");

ALTER TABLE "user_info" ADD FOREIGN KEY ("faculity_id") REFERENCES "faculity" ("id");

ALTER TABLE "department" ADD FOREIGN KEY ("faculity_id") REFERENCES "faculity" ("id");

ALTER TABLE "topic_registration" ADD FOREIGN KEY ("lecture_id") REFERENCES "user" ("id");

ALTER TABLE "topic_registration" ADD FOREIGN KEY ("faculity_id") REFERENCES "user" ("id");

ALTER TABLE "topic" ADD FOREIGN KEY ("lecture_id") REFERENCES "user" ("id");

ALTER TABLE "topic" ADD FOREIGN KEY ("group_id") REFERENCES "group_detail" ("id");

ALTER TABLE "topic" ADD FOREIGN KEY ("department_id") REFERENCES "department" ("id");

ALTER TABLE "topic" ADD FOREIGN KEY ("status_id") REFERENCES "status" ("id");

ALTER TABLE "topic" ADD FOREIGN KEY ("exam_id") REFERENCES "exam" ("id");

ALTER TABLE "topic" ADD FOREIGN KEY ("topic_report_id") REFERENCES "topic_report" ("id");

ALTER TABLE "topic" ADD FOREIGN KEY ("acceptance_board_id") REFERENCES "acceptance_board" ("id");

ALTER TABLE "group_student" ADD FOREIGN KEY ("group_id") REFERENCES "group_detail" ("id");

ALTER TABLE "group_student" ADD FOREIGN KEY ("student_id") REFERENCES "user" ("id");

ALTER TABLE "acceptance_board" ADD FOREIGN KEY ("chairman") REFERENCES "user" ("id");

ALTER TABLE "acceptance_board" ADD FOREIGN KEY ("reviewer1") REFERENCES "user" ("id");

ALTER TABLE "acceptance_board" ADD FOREIGN KEY ("reviewer2") REFERENCES "user" ("id");

ALTER TABLE "acceptance_board" ADD FOREIGN KEY ("secretary") REFERENCES "user" ("id");

ALTER TABLE "acceptance_report" ADD FOREIGN KEY ("acceptance_board_id") REFERENCES "acceptance_board" ("id");

ALTER TABLE "acceptance_report" ADD FOREIGN KEY ("topic_id") REFERENCES "topic" ("id");
