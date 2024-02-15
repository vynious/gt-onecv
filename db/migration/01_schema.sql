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


-- +goose Down

DROP TABLE IF EXISTS notifications;
DROP TABLE IF EXISTS registration;
DROP TABLE IF EXISTS students;
DROP TABLE IF EXISTS teachers;
