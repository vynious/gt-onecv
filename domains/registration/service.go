package registration

import (
	"context"
	"fmt"
	"github.com/vynious/gt-onecv/db"
	"github.com/vynious/gt-onecv/db/sqlc"
	"strings"
)

type RegistrationService struct {
	*db.Repository
}

func SpawnRegistrationService(db *db.Repository) *RegistrationService {
	return &RegistrationService{
		db,
	}
}

func (rs *RegistrationService) CreateRegistration(ctx context.Context, tEmail, sEmail string) (sqlc.Registration, error) {
	registration, err := rs.Queries.CreateRegistration(ctx, sqlc.CreateRegistrationParams{
		Email:   tEmail,
		Email_2: sEmail,
	})
	if err != nil {
		return sqlc.Registration{}, fmt.Errorf("failed to create enrollment %w", err)
	}
	return registration, nil
}

func (rs *RegistrationService) GetRegistrationsByTEmail(ctx context.Context, email string) ([]string, error) {
	registrations, err := rs.Queries.GetRegistrationsByTeacherEmail(ctx, email)
	if err != nil {
		return nil, fmt.Errorf("failed to get registration %w", err)
	}
	return registrations, nil
}

func (rs *RegistrationService) GetCommonRegistrationsByTEmails(ctx context.Context, emails []string) ([]string, error) {
	strArray := "{" + strings.Join(emails, ",") + "}"
	common, err := rs.Queries.GetCommonRegistrationsByTeachersEmail(ctx, sqlc.GetCommonRegistrationsByTeachersEmailParams{
		Email: strArray,
		ID:    int32(len(emails)),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to get common registrations %w", err)
	}
	return common, nil
}
