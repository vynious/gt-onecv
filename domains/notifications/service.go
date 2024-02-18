package notifications

import (
	"context"
	"github.com/vynious/gt-onecv/db"
	"github.com/vynious/gt-onecv/utils"
	"regexp"
)

type NotificationService struct {
	db.IRepository
}

func SpawnNotificationService(db db.IRepository) *NotificationService {
	return &NotificationService{
		db,
	}
}

func (ns *NotificationService) GetNotifiableStudents(ctx context.Context, teacherEmail, content string) ([]string, error) {

	existingTeacher, err := ns.GetTeacherByEmail(ctx, teacherEmail)
	if err != nil {
		return nil, utils.ErrTeacherNotFound
	}

	studentEmails, err := ns.GetNotSuspendedStudentEmailsUnderTeacherId(ctx, existingTeacher.ID)
	if err != nil {
		return nil, utils.ErrStudentNotFound
	}

	extractedEmails := ExtractEmails(content)

	emailSet := make(map[string]struct{})
	var uniqueEmails []string

	for _, email := range studentEmails {
		if _, exists := emailSet[email]; !exists {
			emailSet[email] = struct{}{}
			uniqueEmails = append(uniqueEmails, email)
		}
	}

	for _, email := range extractedEmails {
		if _, exists := emailSet[email]; !exists {
			uniqueEmails = append(uniqueEmails, email)
		}
	}
	return uniqueEmails, nil
}

func ExtractEmails(input string) []string {
	emailRegex := regexp.MustCompile(`[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}`)
	return emailRegex.FindAllString(input, -1)
}
