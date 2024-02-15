
-- name: CreateNotification :one
insert into notifications
(teacher_id, content)
values
    ($1, $2)
returning *;


-- name: GetEligibleRecipients :many
SELECT DISTINCT s.id, s.name, s.email
FROM students s
WHERE NOT s.is_suspended
  AND (
    s.id IN (
        SELECT e.student_id
        FROM enrollments e
                 JOIN teachers t ON e.teacher_id = t.id
        WHERE t.email = $1
    )
        OR s.email = ANY($2)
    );
