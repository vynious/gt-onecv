
-- name: CreateEnrollment :one
INSERT INTO enrollments (teacher_id, student_id)
SELECT t.id, s.id
FROM teachers t, students s
WHERE t.email = $1 AND s.email = $2
    RETURNING *;



-- name: GetEnrollmentsByTeacherEmail :many
SELECT e.id, e.student_id, s.name, s.email, s.is_suspended
FROM enrollments e
         JOIN students s ON e.student_id = s.id
         JOIN teachers t ON e.teacher_id = t.id
WHERE t.email = $1;


