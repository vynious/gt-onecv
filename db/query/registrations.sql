-- name: RegisterStudentUnderTeacher :one
insert into registrations
(teacher_id, student_id)
values ($1, $2)
returning *;


-- name: GetStudentsUnderTeacher :many
select student_id
from registrations
where teacher_id = $1;



-- name: GetStudentEmailsByTeacherId :many
SELECT s.email
FROM students s
    JOIN registrations r ON s.id = r.student_id
WHERE r.teacher_id = $1;

-- name: GetNotSuspendedStudentEmailsUnderTeacherId :many
SELECT s.email
FROM students s
         JOIN registrations r ON s.id = r.student_id
WHERE r.teacher_id = $1 AND s.is_suspended = FALSE;
