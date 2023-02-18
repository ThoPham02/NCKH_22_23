CREATE TABLE IF NOT EXISTS `users` (
    id INT AUTO_INCREMENT PRIMARY KEY,
    username VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL,
    user_type ENUM('students', 'lecturers', 'faculties', 'departments', 'admin') NOT NULL
);

INSERT INTO `users` (username, password, user_type) VALUES
('fos', 'password789', 'faculties'),
('foa', 'password321', 'faculties'),
('foe', 'password321', 'faculties'),
('fob', 'password321', 'faculties'),
('fom', 'password321', 'faculties'),
('dop', 'password789', 'departments'),
('doe', 'password789', 'departments'),
('dom', 'password789', 'departments'),
('dof', 'password789', 'departments'),
('dos', 'password789', 'departments'),
('nva', 'password789', 'lecturers'),
('nvb', 'password321', 'lecturers'),
('nvc', 'password321', 'lecturers'),
('nvd', 'password321', 'lecturers'),
('nve', 'password321', 'lecturers'),
('2021050001', 'password123', 'students'),
('2021050002', 'password456', 'students'),
('2021050003', 'password456', 'students'),
('2021050004', 'password456', 'students'),
('2021050005', 'password456', 'students'),
('admin', 'admin123', 'admin');

CREATE TABLE IF NOT EXISTS `reference` (
    id INT AUTO_INCREMENT PRIMARY KEY,
    reference_url VARCHAR(255) NOT NULL
);
INSERT INTO `reference` (reference_url) VALUES
('https://www.example.com/'),
('https://www.google.com/'),
('https://www.github.com/'),
('https://www.youtube.com/'),
('https://www.stackoverflow.com/');

CREATE TABLE IF NOT EXISTS `faculties` (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    user_id INT NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users(id)
);

INSERT INTO `faculties` (name, user_id) VALUES
('Faculty of Science', 1),
('Faculty of Arts', 2),
('Faculty of Engineering', 3),
('Faculty of Business', 4),
('Faculty of Medicine', 5);

CREATE TABLE IF NOT EXISTS `departments` (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    faculty_id INT NOT NULL,
    user_id INT NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users(id),
    FOREIGN KEY (faculty_id) REFERENCES faculties(id)
);

INSERT INTO `departments` (name, faculty_id, user_id) VALUES
('Department of Physics', 1, 6),
('Department of English', 2, 7),
('Department of Mechanical Engineering', 3, 8),
('Department of Finance', 4, 9),
('Department of Surgery', 5, 10);

CREATE TABLE IF NOT EXISTS `lectureres` (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    user_id INT NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users(id)
);
INSERT INTO `lectureres` (name, user_id) VALUES
('Professor Johnson', 11),
('Dr. Smith', 12),
('Professor Doe', 13),
('Dr. Adams', 14),
('Dr. Brown', 15);

CREATE TABLE IF NOT EXISTS `students` (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    user_id INT NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users(id)
);
INSERT INTO `students` (name, user_id) VALUES
('John Doe', 16),
('Jane Doe', 17),
('Alex Smith', 18),
('Chris Johnson', 19),
('Lisa Adams', 20);

CREATE TABLE IF NOT EXISTS `registration_topics` (
    id int AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    description VARCHAR(255),
    lecturer_id INT NOT NULL,
    department_id INT NOT NULL,
    FOREIGN KEY (lecturer_id) REFERENCES lectureres(id),
    FOREIGN KEY (department_id) REFERENCES departments(id)
);
INSERT INTO `registration_topics` (name, description, lecturer_id, department_id) VALUES
('Introduction to Physics', 'Introduction to Physics', 1, 1),
('English Literature', 'Study of English Literature', 2, 2),
('Mechanical Engineering Fundamentals', 'Fundamentals of Mechanical Engineering', 3, 3),
('Finance Principles', 'Study of Finance Principles', 4, 4),
('Introduction to Surgery', 'Introduction to Surgery', 5, 5);

CREATE TABLE IF NOT EXISTS `groups` (
    id int AUTO_INCREMENT PRIMARY KEY,
    registration_topic_id int not null,
    FOREIGN KEY (registration_topic_id) REFERENCES registration_topics(id)
);
INSERT INTO `groups` (registration_topic_id) VALUES (1);

CREATE TABLE IF NOT EXISTS `student_groups`(
    id int AUTO_INCREMENT PRIMARY KEY,
    student_id int NOT NULL,
    group_id int NOT NULL,
    FOREIGN KEY (student_id) REFERENCES students(id),
    FOREIGN KEY (group_id) REFERENCES `groups`(id)
);

INSERT INTO `student_groups` (student_id, group_id) VALUES
(1, 1),
(2, 1),
(3, 1),
(4, 1),
(5, 1);

CREATE TABLE IF NOT EXISTS `topics` (
    id int AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    description VARCHAR(255),
    lecturer_id INT NOT NULL,
    department_id INT NOT NULL,
    group_id INT NOT NULL,
    FOREIGN KEY (lecturer_id) REFERENCES lectureres(id),
    FOREIGN KEY (department_id) REFERENCES departments(id),
    FOREIGN KEY (group_id) REFERENCES `groups`(id)
);

INSERT INTO `topics` (name, description, lecturer_id, department_id, group_id) VALUES
('Electricity and Magnetism', 'Study of Electricity and Magnetism', 1, 1, 1);