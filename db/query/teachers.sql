-- name: CreateTeacher :one
insert into teachers
    (name, email)
values
    ($1 , $2)
    returning *;


-- name: GetAllTeachers :many


-- name: GetTeacher :one
select *
from teachers
where id = $1;


-- name: UpdateTeacherEmail :one
update teachers
set email = $2
where id = $1
returning *;



-- name: DeleteTeacher :one
delete from teachers
where id = $1
returning *;