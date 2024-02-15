-- name: CreateStudent :one
insert into students
(name, email)
values
    ($1 , $2)
    returning *;


-- name: GetAllStudents :many

-- name: GetStudent :one
select *
from students
where id = $1;

-- name: GetStudentEmail :one
select email
from students
where id = $1;


-- name: UpdateStudentSuspensionByEmail :one
update students
set is_suspended = $2
where email = $1
returning *;

-- name: UpdateStudentEmail :one
update students
set email = $2
where id = $1
returning *;

-- name: DeleteStudent :one
delete from students
where id = $1
returning *;

