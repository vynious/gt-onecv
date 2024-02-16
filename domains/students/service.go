package students

import (
	"context"
	"github.com/vynious/gt-onecv/db"
	"github.com/vynious/gt-onecv/db/sqlc"
	"github.com/vynious/gt-onecv/utils"
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
		return sqlc.Student{}, utils.ConvertToAPIError(err)
	}
	return student, nil
}
