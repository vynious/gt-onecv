package registrations

import (
	"context"
	"github.com/vynious/gt-onecv/db"
	"github.com/vynious/gt-onecv/db/sqlc"
	"github.com/vynious/gt-onecv/utils"
)

type RegistrationService struct {
	*db.Repository
}

func SpawnRegistrationService(db *db.Repository) *RegistrationService {
	return &RegistrationService{
		db,
	}
}

func (rs *RegistrationService) CreateRegistration(ctx context.Context, tEmail, sEmail string) error {

	existingStudent, err := rs.Queries.GetStudentByEmail(ctx, sEmail)
	if err != nil {
		return utils.ErrStudentNotFound
	}
	existingTeacher, err := rs.Queries.GetTeacherByEmail(ctx, tEmail)
	if err != nil {
		return utils.ErrTeacherNotFound
	}

	arg := sqlc.RegisterStudentUnderTeacherParams{
		TeacherID: existingTeacher.ID,
		StudentID: existingStudent.ID,
	}

	_, err = rs.Queries.RegisterStudentUnderTeacher(ctx, arg)
	if err != nil {
		return utils.ConvertToAPIError(err)
	}

	return nil
}

func (rs *RegistrationService) GetRegistrationsByTEmail(ctx context.Context, teacherEmail string) ([]string, error) {
	existingTeacher, err := rs.Queries.GetTeacherByEmail(ctx, teacherEmail)
	if err != nil {
		return nil, utils.ErrTeacherNotFound
	}
	students, err := rs.Queries.GetStudentsUnderTeacher(ctx, existingTeacher.ID)
	if err != nil {
		return nil, utils.ErrStudentNotFound
	}

	var studentEmails []string

	for _, studentId := range students {
		student, err := rs.Queries.GetStudentById(ctx, studentId)
		if err != nil {
			return nil, utils.ConvertToAPIError(err)
		}
		studentEmails = append(studentEmails, student.Email)
	}
	return studentEmails, nil
}

func (rs *RegistrationService) GetCommonRegistrationsByTEmails(ctx context.Context, teacherEmails []string) ([]string, error) {
	emailCounts := make(map[string]int)
	var totalEmails []string // Collect all emails across teachers

	for _, teacherEmail := range teacherEmails {
		teacher, err := rs.Queries.GetTeacherByEmail(ctx, teacherEmail)
		if err != nil {
			return nil, utils.ConvertToAPIError(err)
		}

		// directly fetch student emails for the teacher
		studentEmails, err := rs.Queries.GetStudentEmailsByTeacherId(ctx, teacher.ID)
		if err != nil {
			return nil, utils.ConvertToAPIError(err)
		}

		for _, email := range studentEmails {
			emailCounts[email]++
			totalEmails = append(totalEmails, email)
		}
	}

	// filter to find common emails
	var commonEmails []string
	for _, email := range totalEmails {
		if emailCounts[email] == len(teacherEmails) {
			commonEmails = append(commonEmails, email)
			emailCounts[email] = -1
		}
	}

	return commonEmails, nil
}
