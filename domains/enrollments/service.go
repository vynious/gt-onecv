package enrollments

import (
	"context"
	"fmt"
	"github.com/vynious/gt-onecv/db"
	"github.com/vynious/gt-onecv/db/sqlc"
)

type EnrollmentService struct {
	*db.Repository
}

func SpawnEnrollmentService(db *db.Repository) *EnrollmentService {
	return &EnrollmentService{
		db,
	}
}

func (es *EnrollmentService) CreateEnrollment(ctx context.Context, tEmail, sEmail string) (sqlc.Enrollment, error) {
	enrollment, err := es.Queries.CreateEnrollment(ctx, sqlc.CreateEnrollmentParams{
		Email:   tEmail,
		Email_2: sEmail,
	})
	if err != nil {
		return sqlc.Enrollment{}, fmt.Errorf("failed to create enrollment %w", err)
	}
	return enrollment, nil
}

func (es *EnrollmentService) ReadEnrollmentsByTEmail(ctx context.Context, email string) ([]sqlc.GetEnrollmentsByTeacherEmailRow, error) {
	enrollments, err := es.Queries.GetEnrollmentsByTeacherEmail(ctx, email)
	if err != nil {
		return nil, fmt.Errorf("failed to get enrollments %w", err)
	}
	return enrollments, nil
}
