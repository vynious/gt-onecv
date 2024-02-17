package db

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/vynious/gt-onecv/db/sqlc"
)

type IRepository interface {
	GetTeacherByEmail(ctx context.Context, teacherEmail string) (sqlc.Teacher, error)
	GetStudentByEmail(ctx context.Context, studentEmail string) (sqlc.Student, error)
	GetStudentById(ctx context.Context, studentId int32) (sqlc.Student, error)
	GetStudentEmailsByTeacherId(ctx context.Context, teacherId int32) ([]string, error)
	RegisterStudentUnderTeacher(ctx context.Context, params sqlc.RegisterStudentUnderTeacherParams) (sqlc.Registration, error)
	GetNotSuspendedStudentEmailsUnderTeacherId(ctx context.Context, teacherId int32) ([]string, error)
	GetStudentsUnderTeacher(ctx context.Context, teacherId int32) ([]int32, error)
	UpdateStudentSuspensionByEmail(ctx context.Context, params sqlc.UpdateStudentSuspensionByEmailParams) (sqlc.Student, error)
}

type Repository struct {
	*sqlc.Queries
	DB *pgxpool.Pool
}

func SpawnRepository(db *pgxpool.Pool) *Repository {
	return &Repository{
		DB:      db,
		Queries: sqlc.New(db),
	}
}
