// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.24.0
// source: students.sql

package sqlc

import (
	"context"
)

const createStudent = `-- name: CreateStudent :one
insert into students
(name, email)
values
    ($1 , $2)
    returning id, name, email, is_suspended
`

type CreateStudentParams struct {
	Name  string
	Email string
}

func (q *Queries) CreateStudent(ctx context.Context, arg CreateStudentParams) (Student, error) {
	row := q.db.QueryRow(ctx, createStudent, arg.Name, arg.Email)
	var i Student
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Email,
		&i.IsSuspended,
	)
	return i, err
}

const deleteStudent = `-- name: DeleteStudent :one
delete from students
where id = $1
returning id, name, email, is_suspended
`

func (q *Queries) DeleteStudent(ctx context.Context, id int32) (Student, error) {
	row := q.db.QueryRow(ctx, deleteStudent, id)
	var i Student
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Email,
		&i.IsSuspended,
	)
	return i, err
}

const getAllStudents = `-- name: GetAllStudents :many

select id, name, email, is_suspended
from students
where id = $1
`

func (q *Queries) GetAllStudents(ctx context.Context, id int32) ([]Student, error) {
	rows, err := q.db.Query(ctx, getAllStudents, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Student
	for rows.Next() {
		var i Student
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Email,
			&i.IsSuspended,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getStudentEmail = `-- name: GetStudentEmail :one
select email
from students
where id = $1
`

func (q *Queries) GetStudentEmail(ctx context.Context, id int32) (string, error) {
	row := q.db.QueryRow(ctx, getStudentEmail, id)
	var email string
	err := row.Scan(&email)
	return email, err
}

const updateStudentEmail = `-- name: UpdateStudentEmail :one
update students
set email = $2
where id = $1
returning id, name, email, is_suspended
`

type UpdateStudentEmailParams struct {
	ID    int32
	Email string
}

func (q *Queries) UpdateStudentEmail(ctx context.Context, arg UpdateStudentEmailParams) (Student, error) {
	row := q.db.QueryRow(ctx, updateStudentEmail, arg.ID, arg.Email)
	var i Student
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Email,
		&i.IsSuspended,
	)
	return i, err
}

const updateStudentSuspension = `-- name: UpdateStudentSuspension :one
update students
set is_suspended = $2
where id = $1
returning id, name, email, is_suspended
`

type UpdateStudentSuspensionParams struct {
	ID          int32
	IsSuspended bool
}

func (q *Queries) UpdateStudentSuspension(ctx context.Context, arg UpdateStudentSuspensionParams) (Student, error) {
	row := q.db.QueryRow(ctx, updateStudentSuspension, arg.ID, arg.IsSuspended)
	var i Student
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Email,
		&i.IsSuspended,
	)
	return i, err
}
