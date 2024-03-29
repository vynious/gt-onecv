// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.24.0
// source: registrations.sql

package sqlc

import (
	"context"
)

const getNotSuspendedStudentEmailsUnderTeacherId = `-- name: GetNotSuspendedStudentEmailsUnderTeacherId :many
SELECT s.email
FROM students s
         JOIN registrations r ON s.id = r.student_id
WHERE r.teacher_id = $1 AND s.is_suspended = FALSE
`

func (q *Queries) GetNotSuspendedStudentEmailsUnderTeacherId(ctx context.Context, teacherID int32) ([]string, error) {
	rows, err := q.db.Query(ctx, getNotSuspendedStudentEmailsUnderTeacherId, teacherID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []string
	for rows.Next() {
		var email string
		if err := rows.Scan(&email); err != nil {
			return nil, err
		}
		items = append(items, email)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getStudentEmailsByTeacherId = `-- name: GetStudentEmailsByTeacherId :many
SELECT s.email
FROM students s
    JOIN registrations r ON s.id = r.student_id
WHERE r.teacher_id = $1
`

func (q *Queries) GetStudentEmailsByTeacherId(ctx context.Context, teacherID int32) ([]string, error) {
	rows, err := q.db.Query(ctx, getStudentEmailsByTeacherId, teacherID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []string
	for rows.Next() {
		var email string
		if err := rows.Scan(&email); err != nil {
			return nil, err
		}
		items = append(items, email)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getStudentsUnderTeacher = `-- name: GetStudentsUnderTeacher :many
select student_id
from registrations
where teacher_id = $1
`

func (q *Queries) GetStudentsUnderTeacher(ctx context.Context, teacherID int32) ([]int32, error) {
	rows, err := q.db.Query(ctx, getStudentsUnderTeacher, teacherID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []int32
	for rows.Next() {
		var student_id int32
		if err := rows.Scan(&student_id); err != nil {
			return nil, err
		}
		items = append(items, student_id)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const registerStudentUnderTeacher = `-- name: RegisterStudentUnderTeacher :one
insert into registrations
(teacher_id, student_id)
values ($1, $2)
returning id, student_id, teacher_id
`

type RegisterStudentUnderTeacherParams struct {
	TeacherID int32
	StudentID int32
}

func (q *Queries) RegisterStudentUnderTeacher(ctx context.Context, arg RegisterStudentUnderTeacherParams) (Registration, error) {
	row := q.db.QueryRow(ctx, registerStudentUnderTeacher, arg.TeacherID, arg.StudentID)
	var i Registration
	err := row.Scan(&i.ID, &i.StudentID, &i.TeacherID)
	return i, err
}
