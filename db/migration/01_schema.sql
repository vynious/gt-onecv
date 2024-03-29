-- +goose Up

CREATE TABLE teachers (
                          id SERIAL PRIMARY KEY,
                          name VARCHAR(255) NOT NULL,
                          email VARCHAR(255) NOT NULL UNIQUE
);

CREATE INDEX idx_teachers_email ON teachers(email);

CREATE TABLE students (
                          id SERIAL PRIMARY KEY,
                          name VARCHAR(255) NOT NULL,
                          email VARCHAR(255) NOT NULL UNIQUE,
                          is_suspended BOOLEAN DEFAULT FALSE NOT NULL
);

CREATE INDEX idx_students_email ON students(email);

CREATE TABLE registrations (
                             id SERIAL PRIMARY KEY,
                             student_id INT NOT NULL,
                             teacher_id INT NOT NULL,
                             UNIQUE(student_id, teacher_id),
                             FOREIGN KEY (student_id) REFERENCES students(id) ON DELETE CASCADE,
                             FOREIGN KEY (teacher_id) REFERENCES teachers(id) ON DELETE CASCADE
);

CREATE TABLE notifications (
                               id SERIAL PRIMARY KEY,
                               teacher_id INT NOT NULL,
                               content TEXT NOT NULL,
                               FOREIGN KEY (teacher_id) REFERENCES teachers(id) ON DELETE CASCADE
);


INSERT INTO teachers (name, email) VALUES
                                       ('Teacher A', 'teachera@example.com'),
                                       ('Teacher B', 'teacherb@example.com'),
                                       ('Teacher C', 'teacherc@example.com'),
                                       ('Teacher D', 'teacherd@example.com'),
                                       ('Teacher E', 'teachere@example.com'),
                                       ('Teacher F', 'teacherf@example.com'),
                                       ('Teacher G', 'teacherg@example.com'),
                                       ('Teacher H', 'teacherh@example.com'),
                                       ('Teacher I', 'teacheri@example.com'),
                                       ('Teacher J', 'teacherj@example.com'),
                                       ('Teacher K', 'teacherk@example.com'),
                                       ('Teacher L', 'teacherl@example.com'),
                                       ('Teacher M', 'teacherm@example.com'),
                                       ('Teacher N', 'teachern@example.com'),
                                       ('Teacher O', 'teachero@example.com'),
                                       ('Teacher P', 'teacherp@example.com'),
                                       ('Teacher Q', 'teacherq@example.com'),
                                       ('Teacher R', 'teacherr@example.com'),
                                       ('Teacher S', 'teachers@example.com'),
                                       ('Teacher T', 'teachert@example.com');

INSERT INTO students (name, email, is_suspended) VALUES
                                                     ('Student 1', 'student1@example.com', FALSE),
                                                     ('Student 2', 'student2@example.com', FALSE),
                                                     ('Student 3', 'student3@example.com', FALSE),
                                                     ('Student 4', 'student4@example.com', FALSE),
                                                     ('Student 5', 'student5@example.com', FALSE),
                                                     ('Student 6', 'student6@example.com', FALSE),
                                                     ('Student 7', 'student7@example.com', FALSE),
                                                     ('Student 8', 'student8@example.com', FALSE),
                                                     ('Student 9', 'student9@example.com', FALSE),
                                                     ('Student 10', 'student10@example.com', FALSE),
                                                     ('Student 11', 'student11@example.com', FALSE),
                                                     ('Student 12', 'student12@example.com', FALSE),
                                                     ('Student 13', 'student13@example.com', FALSE),
                                                     ('Student 14', 'student14@example.com', FALSE),
                                                     ('Student 15', 'student15@example.com', FALSE),
                                                     ('Student 16', 'student16@example.com', FALSE),
                                                     ('Student 17', 'student17@example.com', FALSE),
                                                     ('Student 18', 'student18@example.com', FALSE),
                                                     ('Student 19', 'student19@example.com', FALSE),
                                                     ('Student 20', 'student20@example.com', FALSE);

-- +goose Down

DROP TABLE IF EXISTS notifications;
DROP TABLE IF EXISTS registrations;
DROP TABLE IF EXISTS students;
DROP TABLE IF EXISTS teachers;
