package students

import (
	"context"
	"fmt"
	"github.com/vynious/gt-onecv/db"
	"github.com/vynious/gt-onecv/db/sqlc"
)

type StudentService struct {
	*db.Repository
}

func SpawnStudentService(db *db.Repository) *StudentService {
	return &StudentService{
		db,
	}
}

func (ss *StudentService) ChangeStudentSuspensionStatus(ctx context.Context, sEmail string, suspend bool) (sqlc.Student, error) {
	student, err := ss.Queries.UpdateStudentSuspensionByEmail(ctx, sqlc.UpdateStudentSuspensionByEmailParams{
		Email:       sEmail,
		IsSuspended: suspend,
	})
	if err != nil {
		return sqlc.Student{}, fmt.Errorf("failed to update suspension status %w", err)
	}
	return student, nil
}
