package tests

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/stretchr/testify/suite"
	"github.com/vynious/gt-onecv/db"
	"github.com/vynious/gt-onecv/domains/notifications"
	"log"
	"os"
	"testing"
)

type NotificationServiceSuite struct {
	suite.Suite
	DB                  *pgxpool.Pool
	NotificationService *notifications.NotificationService
}

func TestNotificationServiceSuite(t *testing.T) {
	suite.Run(t, new(NotificationServiceSuite))
}

func (s *NotificationServiceSuite) SetupSuite() {
	var err error
	s.DB, err = pgxpool.New(context.Background(), os.Getenv("DATABASE_TEST_URL"))
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	schemaStatements := []string{
		`CREATE TABLE IF NOT EXISTS teachers (
            id SERIAL PRIMARY KEY,
            name VARCHAR(255) NOT NULL,
            email VARCHAR(255) NOT NULL UNIQUE
        );`,
		`CREATE INDEX IF NOT EXISTS idx_teachers_email ON teachers(email);`,
		`CREATE TABLE IF NOT EXISTS students (
            id SERIAL PRIMARY KEY,
            name VARCHAR(255) NOT NULL,
            email VARCHAR(255) NOT NULL UNIQUE,
            is_suspended BOOLEAN DEFAULT FALSE NOT NULL
        );`,
		`CREATE INDEX IF NOT EXISTS idx_students_email ON students(email);`,
		`CREATE TABLE IF NOT EXISTS registrations (
            id SERIAL PRIMARY KEY,
            student_id INT NOT NULL,
            teacher_id INT NOT NULL,
            UNIQUE(student_id, teacher_id),
            FOREIGN KEY (student_id) REFERENCES students(id) ON DELETE CASCADE,
            FOREIGN KEY (teacher_id) REFERENCES teachers(id) ON DELETE CASCADE
        );`,
		`CREATE TABLE IF NOT EXISTS notifications (
            id SERIAL PRIMARY KEY,
            teacher_id INT NOT NULL,
            content TEXT NOT NULL,
            FOREIGN KEY (teacher_id) REFERENCES teachers(id) ON DELETE CASCADE
        );`,
	}

	for _, stmt := range schemaStatements {
		_, err := s.DB.Exec(context.Background(), stmt)
		if err != nil {
			log.Fatalf("Failed to create schema: %v", err)
		}
	}

	database := db.SpawnRepository(s.DB)
	s.NotificationService = notifications.SpawnNotificationService(database)
}

func (s *NotificationServiceSuite) TearDownSuite() {
	s.DB.Close()
}

func (s *NotificationServiceSuite) SetupTest() {
	ctx := context.Background()
	teachers := []string{
		"'Teacher A', 'teachera@example.com'",
		"'Teacher B', 'teacherb@example.com'",
		"'Teacher C', 'teacherc@example.com'",
		"'Teacher D', 'teacherd@example.com'",
		"'Teacher E', 'teachere@example.com'",
		"'Teacher F', 'teacherf@example.com'",
		"'Teacher G', 'teacherg@example.com'",
		"'Teacher H', 'teacherh@example.com'",
		"'Teacher I', 'teacheri@example.com'",
		"'Teacher J', 'teacherj@example.com'",
		"'Teacher K', 'teacherk@example.com'",
		"'Teacher L', 'teacherl@example.com'",
		"'Teacher M', 'teacherm@example.com'",
		"'Teacher N', 'teachern@example.com'",
		"'Teacher O', 'teachero@example.com'",
		"'Teacher P', 'teacherp@example.com'",
		"'Teacher Q', 'teacherq@example.com'",
		"'Teacher R', 'teacherr@example.com'",
		"'Teacher S', 'teachers@example.com'",
		"'Teacher T', 'teachert@example.com'",
	}

	students := []string{
		"'Student 1', 'student1@example.com', FALSE",
		"'Student 2', 'student2@example.com', FALSE",
		"'Student 3', 'student3@example.com', FALSE",
		"'Student 4', 'student4@example.com', FALSE",
		"'Student 5', 'student5@example.com', FALSE",
		"'Student 6', 'student6@example.com', FALSE",
		"'Student 7', 'student7@example.com', FALSE",
		"'Student 8', 'student8@example.com', FALSE",
		"'Student 9', 'student9@example.com', FALSE",
		"'Student 10', 'student10@example.com', FALSE",
		"'Student 11', 'student11@example.com', FALSE",
		"'Student 12', 'student12@example.com', FALSE",
		"'Student 13', 'student13@example.com', FALSE",
		"'Student 14', 'student14@example.com', FALSE",
		"'Student 15', 'student15@example.com', FALSE",
		"'Student 16', 'student16@example.com', FALSE",
		"'Student 17', 'student17@example.com', FALSE",
		"'Student 18', 'student18@example.com', FALSE",
		"'Student 19', 'student19@example.com', FALSE",
		"'Student 20', 'student20@example.com', FALSE",
	}

	for _, teacher := range teachers {
		_, err := s.DB.Exec(ctx, fmt.Sprintf("INSERT INTO teachers (name, email) VALUES (%s) ON CONFLICT (email) DO NOTHING;", teacher))
		if err != nil {
			s.T().Fatal("Failed to insert test teacher:", err)
		}
	}

	for _, student := range students {
		_, err := s.DB.Exec(ctx, fmt.Sprintf("INSERT INTO students (name, email, is_suspended) VALUES (%s) ON CONFLICT (email) DO NOTHING;", student))
		if err != nil {
			s.T().Fatal("Failed to insert test student:", err)
		}
	}
}

func (s *NotificationServiceSuite) TearDownTest() {
	ctx := context.Background()
	_, err := s.DB.Exec(ctx, "DELETE FROM registrations; DELETE FROM students; DELETE FROM teachers;")
	if err != nil {
		s.T().Fatal("Failed to clean up test data:", err)
	}
}

func (s *NotificationServiceSuite) TestGetNotifiableStudentsSuccess() {
	ctx := context.Background()
	teacherEmail := "teachera@example.com"
	content := "Hello students! [student1@example.com]"

	s.createRegistration(ctx, teacherEmail, "student1@example.com")

	notifiableEmails, err := s.NotificationService.GetNotifiableStudents(ctx, teacherEmail, content)
	s.NoError(err)
	s.Len(notifiableEmails, 1, "There should be one notifiable student")
	s.Contains(notifiableEmails, "student1@example.com", "The student email should be in the notifiable list")
}

func TestExtractEmails(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected []string
	}{
		{
			name:     "single email",
			input:    "test@example.com",
			expected: []string{"test@example.com"},
		},
		{
			name:     "multiple emails",
			input:    "first@example.com, second@test.com",
			expected: []string{"first@example.com", "second@test.com"},
		},
		{
			name:     "no emails",
			input:    "no emails here",
			expected: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := notifications.ExtractEmails(tt.input)
			if len(got) != len(tt.expected) {
				t.Errorf("Expected %d emails, got %d", len(tt.expected), len(got))
			}
			for i, email := range got {
				if email != tt.expected[i] {
					t.Errorf("Expected email %s, got %s", tt.expected[i], email)
				}
			}
		})
	}
}

func (s *NotificationServiceSuite) createRegistration(ctx context.Context, teacherEmail, studentEmail string) {
	_, err := s.DB.Exec(ctx, `INSERT INTO registrations (student_id, teacher_id) 
	SELECT s.id, t.id FROM students s, teachers t WHERE s.email = $1 AND t.email = $2;`, studentEmail, teacherEmail)
	if err != nil {
		s.T().Fatal("Failed to create registration:", err)
	}
}
