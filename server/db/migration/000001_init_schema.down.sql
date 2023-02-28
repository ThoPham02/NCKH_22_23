ALTER TABLE IF EXISTS "topics" DROP CONSTRAINT "topics_status_id_fkey";
ALTER TABLE IF EXISTS "topics" DROP CONSTRAINT "topics_department_id_fkey";

ALTER TABLE IF EXISTS "students" DROP CONSTRAINT "students_user_id_fkey";

ALTER TABLE IF EXISTS "lecturers" DROP CONSTRAINT "lecturers_user_id_fkey";

ALTER TABLE IF EXISTS "groups" DROP CONSTRAINT "groups_topic_id_fkey";
ALTER TABLE IF EXISTS "groups" DROP CONSTRAINT "groups_lecturer_id_fkey";
ALTER TABLE IF EXISTS "groups" DROP CONSTRAINT "groups_term_id_fkey";

ALTER TABLE IF EXISTS "departments" DROP CONSTRAINT "departments_user_id_fkey";
ALTER TABLE IF EXISTS "departments" DROP CONSTRAINT "departments_faculty_id_fkey";

ALTER TABLE IF EXISTS "faculties" DROP CONSTRAINT "faculties_user_id_fkey";

ALTER TABLE IF EXISTS "student_groups" DROP CONSTRAINT "student_groups_group_id_fkey";
ALTER TABLE IF EXISTS "student_groups" DROP CONSTRAINT "student_groups_student_id_fkey";

DROP TABLE IF EXISTS "departments";
DROP TABLE IF EXISTS "faculties";
DROP TABLE IF EXISTS "groups";
DROP TABLE IF EXISTS "lecturers";
DROP TABLE IF EXISTS "references";
DROP TABLE IF EXISTS "status";
DROP TABLE IF EXISTS "student_groups";
DROP TABLE IF EXISTS "students";
DROP TABLE IF EXISTS "terms";
DROP TABLE IF EXISTS "topics";
DROP TABLE IF EXISTS "users";