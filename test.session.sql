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

INSERT INTO `reference` (reference_url) VALUES
('https://www.example.com/'),
('https://www.google.com/'),
('https://www.github.com/'),
('https://www.youtube.com/'),
('https://www.stackoverflow.com/');

INSERT INTO `faculties` (name, user_id) VALUES
('Faculty of Science', 1),
('Faculty of Arts', 2),
('Faculty of Engineering', 3),
('Faculty of Business', 4),
('Faculty of Medicine', 5);

INSERT INTO `departments` (name, faculty_id, user_id) VALUES
('Department of Physics', 1, 6),
('Department of English', 2, 7),
('Department of Mechanical Engineering', 3, 8),
('Department of Finance', 4, 9),
('Department of Surgery', 5, 10);

INSERT INTO `lectureres` (name, user_id) VALUES
('Professor Johnson', 11),
('Dr. Smith', 12),
('Professor Doe', 13),
('Dr. Adams', 14),
('Dr. Brown', 15);

INSERT INTO `students` (name, user_id) VALUES
('John Doe', 16),
('Jane Doe', 17),
('Alex Smith', 18),
('Chris Johnson', 19),
('Lisa Adams', 20);

INSERT INTO `registration_topics` (name, description, lecturer_id, department_id) VALUES
('Introduction to Physics', 'Introduction to Physics', 1, 1),
('English Literature', 'Study of English Literature', 2, 2),
('Mechanical Engineering Fundamentals', 'Fundamentals of Mechanical Engineering', 3, 3),
('Finance Principles', 'Study of Finance Principles', 4, 4),
('Introduction to Surgery', 'Introduction to Surgery', 5, 5);

INSERT INTO `groups` (registration_topic_id) VALUES (1);

INSERT INTO `student_groups` (student_id, group_id) VALUES
(1, 1),
(2, 1),
(3, 1),
(4, 1),
(5, 1);

INSERT INTO `topics` (name, description, lecturer_id, department_id, group_id) VALUES
('Electricity and Magnetism', 'Study of Electricity and Magnetism', 1, 1, 1);