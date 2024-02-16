
-- name: CreateNotification :one
insert into notifications
(teacher_id, content)
values
    ($1, $2)
returning *;


-- name: GetEligibleRecipients :many
