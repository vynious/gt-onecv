-- name: CreateTeacher :one
insert into teachers
    (name, email)
values
    ($1 , $2)
    returning *;

-- name: GetTeacherByEmail :one
select *
from teachers
where email = $1;

-- name: GetTeacherById :one
select *
from teachers
where id = $1;

