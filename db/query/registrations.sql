
-- name: CreateRegistration :one
INSERT INTO registrations (teacher_id, student_id)
SELECT t.id, s.id
FROM teachers t, students s
WHERE t.email = $1 AND s.email = $2
    RETURNING *;



-- name: GetRegistrationsByTeacherEmail :many
SELECT s.email
FROM registrations e
         JOIN students s ON e.student_id = s.id
         JOIN teachers t ON e.teacher_id = t.id
WHERE t.email = $1;


-- name: GetCommonRegistrationsByTeachersEmail :many
SELECT s.email
FROM students s
         JOIN registrations r ON s.id = r.student_id
         JOIN teachers t ON t.id = r.teacher_id
WHERE t.email = ANY($1)
GROUP BY s.id
HAVING COUNT(DISTINCT t.id) = $2;


-- name: GetUnsuspendedRegistrationsByTeacherEmail :many
SELECT s.email
FROM registrations e
         JOIN students s ON e.student_id = s.id
         JOIN teachers t ON e.teacher_id = t.id
WHERE t.email = $1 and s.is_suspended = false;