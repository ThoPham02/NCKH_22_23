CREATE TABLE `user_type` (
  `id` INT PRIMARY KEY AUTO_INCREMENT,
  `name` VARCHAR(255) NOT NULL,
  `created_at` datetime DEFAULT (now()),
  `update_at` datetime,
  `deleted_at` datetime
);

CREATE TABLE `users` (
  `id` INT PRIMARY KEY AUTO_INCREMENT,
  `username` VARCHAR(255) NOT NULL,
  `password` VARCHAR(255) NOT NULL,
  `user_type_id` INT NOT NULL,
  `created_at` datetime DEFAULT (now()),
  `update_at` datetime,
  `deleted_at` datetime
);

CREATE TABLE `reference` (
  `id` INT PRIMARY KEY AUTO_INCREMENT,
  `reference_url` VARCHAR(255) NOT NULL,
  `name` VARCHAR(255) NOT NULL,
  `created_at` datetime DEFAULT (now()),
  `update_at` datetime,
  `deleted_at` datetime
);

CREATE TABLE `faculties` (
  `id` INT PRIMARY KEY AUTO_INCREMENT,
  `name` VARCHAR(255) NOT NULL,
  `user_id` INT NOT NULL,
  `created_at` datetime DEFAULT (now()),
  `update_at` datetime,
  `deleted_at` datetime
);

CREATE TABLE `departments` (
  `id` INT PRIMARY KEY AUTO_INCREMENT,
  `name` VARCHAR(255) NOT NULL,
  `faculty_id` INT NOT NULL,
  `user_id` INT NOT NULL,
  `created_at` datetime DEFAULT (now()),
  `update_at` datetime,
  `deleted_at` datetime
);

CREATE TABLE `lectureres` (
  `id` INT PRIMARY KEY AUTO_INCREMENT,
  `name` VARCHAR(255) NOT NULL,
  `user_id` INT NOT NULL,
  `created_at` datetime DEFAULT (now()),
  `update_at` datetime,
  `deleted_at` datetime
);

CREATE TABLE `students` (
  `id` INT PRIMARY KEY AUTO_INCREMENT,
  `name` VARCHAR(255) NOT NULL,
  `user_id` INT NOT NULL,
  `created_at` datetime DEFAULT (now()),
  `update_at` datetime,
  `deleted_at` datetime
);

CREATE TABLE `events` (
  `id` INT PRIMARY KEY AUTO_INCREMENT,
  `name` VARCHAR(255) NOT NULL,
  `url` VARCHAR(255),
  `time_start` datetime DEFAULT (now()),
  `time_end` datetime,
  `created_at` datetime DEFAULT (now()),
  `update_at` datetime,
  `deleted_at` datetime
);

CREATE TABLE `groups` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `topic_id` int NOT NULL,
  `created_at` datetime DEFAULT (now()),
  `update_at` datetime,
  `deleted_at` datetime
);

CREATE TABLE `student_groups` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `student_id` int NOT NULL,
  `group_id` int NOT NULL,
  `created_at` datetime DEFAULT (now()),
  `update_at` datetime,
  `deleted_at` datetime
);

CREATE TABLE `status` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `name` VARCHAR(255) NOT NULL,
  `created_at` datetime DEFAULT (now()),
  `update_at` datetime,
  `deleted_at` datetime
);

CREATE TABLE `topics` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `name` VARCHAR(255) NOT NULL,
  `description` VARCHAR(255),
  `lecturer_id` INT NOT NULL,
  `department_id` INT NOT NULL,
  `event_id` INT NOT NULL,
  `status_id` INT DEFAULT 1,
  `created_at` datetime DEFAULT (now()),
  `update_at` datetime,
  `deleted_at` datetime
);

ALTER TABLE `faculties` ADD FOREIGN KEY (`user_id`) REFERENCES `users` (`id`);

ALTER TABLE `departments` ADD FOREIGN KEY (`user_id`) REFERENCES `users` (`id`);

ALTER TABLE `departments` ADD FOREIGN KEY (`faculty_id`) REFERENCES `faculties` (`id`);

ALTER TABLE `lectureres` ADD FOREIGN KEY (`user_id`) REFERENCES `users` (`id`);

ALTER TABLE `user_type` ADD FOREIGN KEY (`id`) REFERENCES `users` (`user_type_id`);

ALTER TABLE `students` ADD FOREIGN KEY (`user_id`) REFERENCES `users` (`id`);

ALTER TABLE `groups` ADD FOREIGN KEY (`topic_id`) REFERENCES `topics` (`id`);

ALTER TABLE `student_groups` ADD FOREIGN KEY (`student_id`) REFERENCES `students` (`id`);

ALTER TABLE `topics` ADD FOREIGN KEY (`event_id`) REFERENCES `events` (`id`);

ALTER TABLE `topics` ADD FOREIGN KEY (`status_id`) REFERENCES `status` (`id`);

ALTER TABLE `student_groups` ADD FOREIGN KEY (`group_id`) REFERENCES `groups` (`id`);

ALTER TABLE `topics` ADD FOREIGN KEY (`lecturer_id`) REFERENCES `lectureres` (`id`);

ALTER TABLE `topics` ADD FOREIGN KEY (`department_id`) REFERENCES `departments` (`id`);
