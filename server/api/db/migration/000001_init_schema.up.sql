CREATE TABLE "users" (
  "id" BIGSERIAL PRIMARY KEY,
  "username" VARCHAR(255) NOT NULL,
  "password" VARCHAR(255) NOT NULL,
  "email" VARCHAR(255) NOT NULL,
  "permission" BIGINT NOT NULL,
  "created_at" timestamp DEFAULT (now()),
  "updated_at" timestamp,
  "deleted_at" timestamp
);

CREATE TABLE "references" (
  "id" BIGSERIAL PRIMARY KEY,
  "reference_url" VARCHAR(255) NOT NULL,
  "name" VARCHAR(255) NOT NULL,
  "created_at" timestamp DEFAULT (now()),
  "update_at" timestamp,
  "deleted_at" timestamp
);

CREATE TABLE "faculties" (
  "id" BIGSERIAL PRIMARY KEY,
  "name" VARCHAR(255) NOT NULL,
  "user_id" BIGINT UNIQUE NOT NULL,
  "created_at" timestamp DEFAULT (now()),
  "updated_at" timestamp,
  "deleted_at" timestamp
);

CREATE TABLE "departments" (
  "id" BIGSERIAL PRIMARY KEY,
  "name" VARCHAR(255) NOT NULL,
  "faculty_id" BIGINT NOT NULL,
  "user_id" BIGINT UNIQUE NOT NULL,
  "created_at" timestamp DEFAULT (now()),
  "updated_at" timestamp,
  "deleted_at" timestamp
);

CREATE TABLE "lecturers" (
  "id" BIGSERIAL PRIMARY KEY,
  "name" VARCHAR(255) UNIQUE NOT NULL,
  "user_id" BIGINT NOT NULL,
  "created_at" timestamp DEFAULT (now()),
  "updated_at" timestamp,
  "deleted_at" timestamp
);

CREATE TABLE "students" (
  "id" BIGSERIAL PRIMARY KEY,
  "name" VARCHAR(255) NOT NULL,
  "user_id" BIGINT UNIQUE NOT NULL,
  "created_at" timestamp DEFAULT (now()),
  "updated_at" timestamp,
  "deleted_at" timestamp
);

CREATE TABLE "status" (
  "id" BIGSERIAL PRIMARY KEY,
  "name" VARCHAR(255) NOT NULL
);

CREATE TABLE "student_groups" (
  "group_id" BIGint NOT NULL,
  "student_id" BIGint NOT NULL,
  PRIMARY KEY ("group_id", "student_id")
);

CREATE TABLE "groups" (
  "id" BIGSERIAL PRIMARY KEY,
  "topic_id" BIGint NOT NULL,
  "lecturer_id" BIGint NOT NULL,
  "term_id" BIGint NOT NULL,
  "created_at" timestamp DEFAULT (now()),
  "updated_at" timestamp,
  "deleted_at" timestamp
);

CREATE TABLE "terms" (
  "id" BIGSERIAL PRIMARY KEY,
  "name" VARCHAR(255) NOT NULL,
  "created_at" timestamp DEFAULT (now()),
  "updated_at" timestamp,
  "deleted_at" timestamp
);

CREATE TABLE "topics" (
  "id" BIGSERIAL PRIMARY KEY,
  "name" VARCHAR(255) NOT NULL,
  "description" VARCHAR(255),
  "department_id" BIGINT NOT NULL,
  "time_start" date,
  "time_end" date,
  "deadline" date,
  "status_id" BIGint NOT NULL,
  "created_at" timestamp DEFAULT (now()),
  "updated_at" timestamp,
  "deleted_at" timestamp
);

ALTER TABLE "student_groups" ADD FOREIGN KEY ("group_id") REFERENCES "groups" ("id");

ALTER TABLE "student_groups" ADD FOREIGN KEY ("student_id") REFERENCES "students" ("id");

ALTER TABLE "groups" ADD FOREIGN KEY ("topic_id") REFERENCES "topics" ("id");

ALTER TABLE "groups" ADD FOREIGN KEY ("lecturer_id") REFERENCES "lecturers" ("id");

ALTER TABLE "groups" ADD FOREIGN KEY ("term_id") REFERENCES "terms" ("id");

ALTER TABLE "topics" ADD FOREIGN KEY ("status_id") REFERENCES "status" ("id");

ALTER TABLE "faculties" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "departments" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "departments" ADD FOREIGN KEY ("faculty_id") REFERENCES "faculties" ("id");

ALTER TABLE "lecturers" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "students" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "topics" ADD FOREIGN KEY ("department_id") REFERENCES "departments" ("id");


INSERT INTO
  "terms" ("name")
VALUES
  (N'2020-2021'),
  (N'2021-2022'),
  (N'2022-2023'),
  (N'2023-2024');

INSERT INTO
  "references"(reference_url, "name")
VALUES
  ('https://www.youtube.com/', N'YouTube'),
  ('https://www.w3schools.com/', N'W3School'),
  ('https://courses.dev-academy.com/', N'Academy');

INSERT INTO
  "status" ("name")
VALUES
  (N'Đề xuất'),
  (N'Đang đăng ký'),
  (N'Đang thực hiện'),
  (N'Sắp hết hạn'),
  (N'Quá hạn'),
  (N'Đã hoàn thành'),
  (N'Hoàn thành quá hạn');

INSERT INTO
  "users" (username, "password",email,  permission)
VALUES
  ('humgsv01', '12345678', 'sv01@humg.edu.vn', 1),
  ('humgsv02', '12345678', 'sv02@humg.edu.vn', 1),
  ('humgsv03', '12345678', 'sv03@humg.edu.vn', 1),
  ('humgsv04', '12345678', 'sv04@humg.edu.vn', 1),
  ('humgsv05', '12345678', 'sv05@humg.edu.vn', 1),
  ('humgsv06', '12345678', 'sv06@humg.edu.vn', 1),
  ('humgsv07', '12345678', 'sv07@humg.edu.vn', 1),
  ('humgsv08', '12345678', 'sv08@humg.edu.vn', 1),
  ('humgsv09', '12345678', 'sv09@humg.edu.vn', 1),
  ('humgsv10', '12345678', 'sv10@humg.edu.vn', 1),
  ('humgsv11', '12345678', 'sv11@humg.edu.vn', 1),
  ('humgsv12', '12345678', 'sv12@humg.edu.vn', 1),
  ('humgsv13', '12345678', 'sv13@humg.edu.vn', 1),
  ('humgsv14', '12345678', 'sv14@humg.edu.vn', 1),
  ('humgsv15', '12345678', 'sv15@humg.edu.vn', 1),
  ('humgsv16', '12345678', 'sv16@humg.edu.vn', 1),
  ('humgsv17', '12345678', 'sv17@humg.edu.vn', 1),
  ('humgsv18', '12345678', 'sv18@humg.edu.vn', 1),
  ('humgsv19', '12345678', 'sv19@humg.edu.vn', 1),
  ('humgsv20', '12345678', 'sv20@humg.edu.vn', 1),
  ('humggv01', '12345678', 'gv01@humg.edu.vn', 2),
  ('humggv02', '12345678', 'gv02@humg.edu.vn', 2),
  ('humggv03', '12345678', 'gv03@humg.edu.vn', 2),
  ('humggv04', '12345678', 'gv04@humg.edu.vn', 2),
  ('humggv05', '12345678', 'gv05@humg.edu.vn', 2),
  ('humggv06', '12345678', 'gv06@humg.edu.vn', 2),
  ('humggv07', '12345678', 'gv07@humg.edu.vn', 2),
  ('humggv08', '12345678', 'gv08@humg.edu.vn', 2),
  ('humggv09', '12345678', 'gv09@humg.edu.vn', 2),
  ('humggv10', '12345678', 'gv10@humg.edu.vn', 2),
  ('humggv11', '12345678', 'gv11@humg.edu.vn', 2),
  ('humggv12', '12345678', 'gv12@humg.edu.vn', 2),
  ('humggv13', '12345678', 'gv13@humg.edu.vn', 2),
  ('humggv14', '12345678', 'gv14@humg.edu.vn', 2),
  ('humggv15', '12345678', 'gv15@humg.edu.vn', 2),
  ('humggv16', '12345678', 'gv16@humg.edu.vn', 2),
  ('humggv17', '12345678', 'gv17@humg.edu.vn', 2),
  ('humggv18', '12345678', 'gv18@humg.edu.vn', 2),
  ('humggv19', '12345678', 'gv19@humg.edu.vn', 2),
  ('humggv20', '12345678', 'gv20@humg.edu.vn', 2),
  ('humgbm01', '12345678', 'bm01@humg.edu.vn', 3),
  ('humgbm02', '12345678', 'bm02@humg.edu.vn', 3),
  ('humgbm03', '12345678', 'bm03@humg.edu.vn', 3),
  ('humgbm04', '12345678', 'bm04@humg.edu.vn', 3),
  ('humgbm05', '12345678', 'bm05@humg.edu.vn', 3),
  ('humgbm06', '12345678', 'bm06@humg.edu.vn', 3),
  ('humgbm07', '12345678', 'bm07@humg.edu.vn', 3),
  ('humgbm08', '12345678', 'bm08@humg.edu.vn', 3),
  ('humgbm09', '12345678', 'bm09@humg.edu.vn', 3),
  ('humgbm10', '12345678', 'bm10@humg.edu.vn', 3),
  ('humgbm11', '12345678', 'bm11@humg.edu.vn', 3),
  ('humgbm12', '12345678', 'bm12@humg.edu.vn', 3),
  ('humgbm13', '12345678', 'bm13@humg.edu.vn', 3),
  ('humgbm14', '12345678', 'bm14@humg.edu.vn', 3),
  ('humgbm15', '12345678', 'bm15@humg.edu.vn', 3),
  ('humgbm16', '12345678', 'bm16@humg.edu.vn', 3),
  ('humgbm17', '12345678', 'bm17@humg.edu.vn', 3),
  ('humgbm18', '12345678', 'bm18@humg.edu.vn', 3),
  ('humgbm19', '12345678', 'bm19@humg.edu.vn', 3),
  ('humgbm20', '12345678', 'bm20@humg.edu.vn', 3),
  ('humgbm21', '12345678', 'bm21@humg.edu.vn', 3),
  ('humgbm22', '12345678', 'bm22@humg.edu.vn', 3),
  ('humgbm23', '12345678', 'bm23@humg.edu.vn', 3),
  ('humgbm24', '12345678', 'bm24@humg.edu.vn', 3),
  ('humgbm25', '12345678', 'bm25@humg.edu.vn', 3),
  ('humgbm26', '12345678', 'bm26@humg.edu.vn', 3),
  ('humgbm27', '12345678', 'bm27@humg.edu.vn', 3),
  ('humgbm28', '12345678', 'bm28@humg.edu.vn', 3),
  ('humgbm29', '12345678', 'bm29@humg.edu.vn', 3),
  ('humgbm30', '12345678', 'bm30@humg.edu.vn', 3),
  ('humgbm31', '12345678', 'bm31@humg.edu.vn', 3),
  ('humgbm32', '12345678', 'bm32@humg.edu.vn', 3),
  ('humgbm33', '12345678', 'bm33@humg.edu.vn', 3),
  ('humgbm34', '12345678', 'bm34@humg.edu.vn', 3),
  ('humgbm35', '12345678', 'bm35@humg.edu.vn', 3),
  ('humgbm36', '12345678', 'bm36@humg.edu.vn', 3),
  ('humgbm37', '12345678', 'bm37@humg.edu.vn', 3),
  ('humgbm38', '12345678', 'bm38@humg.edu.vn', 3),
  ('humgbm39', '12345678', 'bm39@humg.edu.vn', 3),
  ('humgbm40', '12345678', 'bm40@humg.edu.vn', 3),
  ('humgbm41', '12345678', 'bm41@humg.edu.vn', 3),
  ('humgbm42', '12345678', 'bm42@humg.edu.vn', 3),
  ('humgbm43', '12345678', 'bm43@humg.edu.vn', 3),
  ('humgbm44', '12345678', 'bm44@humg.edu.vn', 3),
  ('humgbm45', '12345678', 'bm45@humg.edu.vn', 3),
  ('humgbm46', '12345678', 'bm46@humg.edu.vn', 3),
  ('humgbm47', '12345678', 'bm47@humg.edu.vn', 3),
  ('humgbm48', '12345678', 'bm48@humg.edu.vn', 3),
  ('humgbm49', '12345678', 'bm49@humg.edu.vn', 3),
  ('humgbm50', '12345678', 'bm50@humg.edu.vn', 3),
  ('humgbm51', '12345678', 'bm51@humg.edu.vn', 3),
  ('humgbm52', '12345678', 'bm52@humg.edu.vn', 3),
  ('humgbm53', '12345678', 'bm53@humg.edu.vn', 3),
  ('humgbm54', '12345678', 'bm54@humg.edu.vn', 3),
  ('humgbm55', '12345678', 'bm55@humg.edu.vn', 3),
  ('humgbm56', '12345678', 'bm56@humg.edu.vn', 3),
  ('humgbm57', '12345678', 'bm57@humg.edu.vn', 3),
  ('humgbm58', '12345678', 'bm58@humg.edu.vn', 3),
  ('humgbm59', '12345678', 'bm59@humg.edu.vn', 3),
  ('humgbm60', '12345678', 'bm60@humg.edu.vn', 3),
  ('humgkh61', '12345678', 'kh61@humg.edu.vn', 4),
  ('humgkh62', '12345678', 'kh62@humg.edu.vn', 4),
  ('humgkh63', '12345678', 'kh63@humg.edu.vn', 4),
  ('humgkh64', '12345678', 'kh64@humg.edu.vn', 4),
  ('humgkh65', '12345678', 'kh65@humg.edu.vn', 4),
  ('humgkh66', '12345678', 'kh66@humg.edu.vn', 4),
  ('humgkh67', '12345678', 'kh67@humg.edu.vn', 4),
  ('humgkh68', '12345678', 'kh68@humg.edu.vn', 4),
  ('humgkh69', '12345678', 'kh69@humg.edu.vn', 4),
  ('humgkh70', '12345678', 'kh70@humg.edu.vn', 4),
  ('humgkh71', '12345678', 'kh71@humg.edu.vn', 4),
  ('humgkh72', '12345678', 'kh72@humg.edu.vn', 4),
  ('humgkhcn', '12345678', 'khcn@humg.edu.vn', 5);

INSERT INTO
  "students" ("name", user_id)
VALUES
  (N'Nguyễn Văn A01', 1),
  (N'Nguyễn Văn A02', 2),
  (N'Nguyễn Văn A03', 3),
  (N'Nguyễn Văn A04', 4),
  (N'Nguyễn Văn A05', 5),
  (N'Nguyễn Văn A06', 6),
  (N'Nguyễn Văn A07', 7),
  (N'Nguyễn Văn A08', 8),
  (N'Nguyễn Văn A09', 9),
  (N'Nguyễn Văn A10', 10),
  (N'Nguyễn Văn A11', 11),
  (N'Nguyễn Văn A12', 12),
  (N'Nguyễn Văn A13', 13),
  (N'Nguyễn Văn A14', 14),
  (N'Nguyễn Văn A15', 15),
  (N'Nguyễn Văn A16', 16),
  (N'Nguyễn Văn A17', 17),
  (N'Nguyễn Văn A18', 18),
  (N'Nguyễn Văn A19', 19),
  (N'Nguyễn Văn A20', 20);

INSERT INTO
  "lecturers" ("name", user_id)
VALUES
  (N'Nguyễn Văn B01', 21),
  (N'Nguyễn Văn B02', 22),
  (N'Nguyễn Văn B03', 23),
  (N'Nguyễn Văn B04', 24),
  (N'Nguyễn Văn B05', 25),
  (N'Nguyễn Văn B06', 26),
  (N'Nguyễn Văn B07', 27),
  (N'Nguyễn Văn B08', 28),
  (N'Nguyễn Văn B09', 29),
  (N'Nguyễn Văn B10', 30),
  (N'Nguyễn Văn B11', 31),
  (N'Nguyễn Văn B12', 32),
  (N'Nguyễn Văn B13', 33),
  (N'Nguyễn Văn B14', 34),
  (N'Nguyễn Văn B15', 35),
  (N'Nguyễn Văn B16', 36),
  (N'Nguyễn Văn B17', 37),
  (N'Nguyễn Văn B18', 38),
  (N'Nguyễn Văn B19', 39),
  (N'Nguyễn Văn B20', 40);

INSERT INTO
  "faculties" ("name", user_id)
VALUES
  (N'Cơ Điện', 41),
  (N'Công Nghệ Thông Tin', 42),
  (N'Dầu Khí', 43),
  (N'Giáo dục quốc phòng', 44),
  (N'Khoa học cơ bản', 45),
  (N'Khoa học và kĩ thuật Địa chất', 46),
  (N'Kinh tế và Quản trị kinh doanh', 47),
  (N'Lý luận chính trị', 48),
  (N'Mỏ', 49),
  (N'Môi trường', 50),
  (N'Trắc địa bản đồ và QLĐĐ', 51),
  (N'Xây dựng', 52);

INSERT INTO
  "departments"("name", faculty_id, user_id)
VALUES
  (N'Máy và Thiết bị', 1, 53),
  (N'Máy và Thiết bị', 1, 54),
  (N'Máy và Thiết bị', 1, 55),
  (N'Máy và Thiết bị', 1, 56),
  (N'Máy và Thiết bị', 1, 57),
  (N'Máy và Thiết bị', 1, 58),
  (N'Công nghệ phần mềm', 2, 59),
  (N'Công nghệ phần mềm', 2, 60),
  (N'Công nghệ phần mềm', 2, 61),
  (N'Công nghệ phần mềm', 2, 62),
  (N'Công nghệ phần mềm', 2, 63),
  (N'Công nghệ phần mềm', 2, 64),
  (N'Công nghệ phần mềm', 2, 65),
  (N'Công nghệ phần mềm', 2, 66),
  (N'Thiết bị dầu khí', 3, 67),
  (N'Thiết bị dầu khí', 3, 68),
  (N'Thiết bị dầu khí', 3, 69),
  (N'Thiết bị dầu khí', 3, 70),
  (N'Thiết bị dầu khí', 3, 71),
  (N'Kĩ thuật Quân sự', 4, 72),
  (N'Kĩ thuật Quân sự', 4, 73),
  (N'Toán', 5, 74),
  (N'Toán', 5, 75),
  (N'Toán', 5, 76),
  (N'Toán', 5, 77),
  (N'Toán', 5, 78),
  (N'Toán', 5, 79),
  (N'Toán', 5, 80),
  (N'Địa chất công trình', 6, 81),
  (N'Địa chất công trình', 6, 82),
  (N'Địa chất công trình', 6, 83),
  (N'Địa chất công trình', 6, 84),
  (N'Địa chất công trình', 6, 85),
  (N'Địa chất công trình', 6, 86),
  (N'Quản trị doanh nghiệp Mỏ', 7, 87),
  (N'Quản trị doanh nghiệp Mỏ', 7, 88),
  (N'Quản trị doanh nghiệp Mỏ', 7, 89),
  (N'Quản trị doanh nghiệp Mỏ', 7, 90),
  (N'Tư tưởng Hồ Chí Minh', 8, 91),
  (N'Tư tưởng Hồ Chí Minh', 8, 92),
  (N'Tư tưởng Hồ Chí Minh', 8, 93),
  (N'Khai thác lộ thiên', 9, 94),
  (N'Khai thác lộ thiên', 9, 95),
  (N'Khai thác lộ thiên', 9, 96),
  (N'Khai thác lộ thiên', 9, 97),
  (N'Địa sinh thái và công nghệ môi trường', 10, 98),
  (N'Địa sinh thái và công nghệ môi trường', 10, 99),
  (N'Địa sinh thái và công nghệ môi trường', 10, 100),
  (N'Trắc địa cao cấp', 11, 101),
  (N'Trắc địa cao cấp', 11, 102),
  (N'Trắc địa cao cấp', 11, 103),
  (N'Trắc địa cao cấp', 11, 104),
  (N'Trắc địa cao cấp', 11, 105),
  (N'Trắc địa cao cấp', 11, 106),
  (N'Trắc địa cao cấp', 11, 107),
  (N'Trắc địa cao cấp', 11, 108),
  (N'Xây dựng hạ tầng cơ sở', 12, 109),
  (N'Xây dựng hạ tầng cơ sở', 12, 110),
  (N'Xây dựng hạ tầng cơ sở', 12, 111);