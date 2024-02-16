-- name: CreateStudent :one
insert into students
(name, email)
values
    ($1 , $2)
    returning *;

-- name: GetStudentByEmail :one
select *
from students
where email = $1;


-- name: GetStudentById :one
select *
from students
where id = $1;

-- name: UpdateStudentSuspensionByEmail :one
update students
set is_suspended = $2
where email = $1
returning *;


-- name: GetStudentEmailsByIds :many
SELECT email FROM students WHERE id = ANY($1);



