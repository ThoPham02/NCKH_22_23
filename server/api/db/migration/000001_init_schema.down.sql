ALTER TABLE "acceptance_report" DROP CONSTRAINT "acceptance_report_topic_id_fkey";
ALTER TABLE "acceptance_report" DROP CONSTRAINT "acceptance_report_acceptance_board_id_fkey";

ALTER TABLE "acceptance_board" DROP CONSTRAINT "acceptance_board_secretary_fkey";
ALTER TABLE "acceptance_board" DROP CONSTRAINT "acceptance_board_reviewer2_fkey";
ALTER TABLE "acceptance_board" DROP CONSTRAINT "acceptance_board_reviewer1_fkey";
ALTER TABLE "acceptance_board" DROP CONSTRAINT "acceptance_board_chairman_fkey";

ALTER TABLE "group_student" DROP CONSTRAINT "group_student_student_id_fkey";
ALTER TABLE "group_student" DROP CONSTRAINT "group_student_group_id_fkey";

ALTER TABLE "topic" DROP CONSTRAINT "topic_acceptance_board_id_fkey";
ALTER TABLE "topic" DROP CONSTRAINT "topic_topic_report_id_fkey";
ALTER TABLE "topic" DROP CONSTRAINT "topic_exam_id_fkey";
ALTER TABLE "topic" DROP CONSTRAINT "topic_status_id_fkey";
ALTER TABLE "topic" DROP CONSTRAINT "topic_department_id_fkey";
ALTER TABLE "topic" DROP CONSTRAINT "topic_group_id_fkey";
ALTER TABLE "topic" DROP CONSTRAINT "topic_lecture_id_fkey";

ALTER TABLE "topic_registration" DROP CONSTRAINT "topic_registration_lecture_id_fkey";
ALTER TABLE "topic_registration" DROP CONSTRAINT "topic_registration_faculity_id_fkey";

ALTER TABLE "department" DROP CONSTRAINT "department_faculity_id_fkey";

ALTER TABLE "user_info" DROP CONSTRAINT "user_info_faculity_id_fkey";

ALTER TABLE "user_info" DROP CONSTRAINT "user_info_user_id_fkey";

ALTER TABLE "user" DROP CONSTRAINT "user_type_account_id_fkey";

DROP TABLE IF EXISTS "type_account";
DROP TABLE IF EXISTS "user";
DROP TABLE IF EXISTS "user_info";
DROP TABLE IF EXISTS "notification";
DROP TABLE IF EXISTS "department";
DROP TABLE IF EXISTS "faculity";
DROP TABLE IF EXISTS "topic_registration";
DROP TABLE IF EXISTS "topic";
DROP TABLE IF EXISTS "status";
DROP TABLE IF EXISTS "exam";
DROP TABLE IF EXISTS "group_student";
DROP TABLE IF EXISTS "group_detail";
DROP TABLE IF EXISTS "acceptance_board";
DROP TABLE IF EXISTS "acceptance_report";
DROP TABLE IF EXISTS "topic_report";