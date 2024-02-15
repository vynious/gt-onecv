package notifications

import (
	"context"
	"fmt"
	"github.com/vynious/gt-onecv/db"
	"regexp"
)

type NotificationService struct {
	*db.Repository
}

func SpawnNotificationService(db *db.Repository) *NotificationService {
	return &NotificationService{
		db,
	}
}

func (ns *NotificationService) GetNotifiableStudents(ctx context.Context, tEmail, content string) ([]string, error) {

	// Get distinct unsuspended students' email addresses registered under the teacher
	registeredStudents, err := ns.Queries.GetUnsuspendedRegistrationsByTeacherEmail(ctx, tEmail)
	if err != nil {
		return nil, fmt.Errorf("failed to get unsuspended students under teacher: %w", err)
	}

	// Initialize the recipients with the registered students
	recipients := make([]string, len(registeredStudents))
	copy(recipients, registeredStudents)

	// Extract mentioned student emails from the notification content
	mentionedEmails := extractEmails(content)

	// Create a map from the recipients slice for quick lookup
	uniqueRecipients := make(map[string]bool, len(recipients))
	for _, email := range recipients {
		uniqueRecipients[email] = true
	}

	// Check the extracted emails against the map and add them if they are not duplicates
	for _, email := range mentionedEmails {
		if !uniqueRecipients[email] {
			recipients = append(recipients, email)
			uniqueRecipients[email] = true // Mark this email as seen
		}
	}

	return recipients, nil
}

func extractEmails(input string) []string {
	emailRegex := regexp.MustCompile(`[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}`)
	return emailRegex.FindAllString(input, -1)
}
