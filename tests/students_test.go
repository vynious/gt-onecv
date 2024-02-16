package tests

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/stretchr/testify/suite"
	"github.com/vynious/gt-onecv/db"
	"github.com/vynious/gt-onecv/domains/students"
	"log"
	"os"
	"testing"
)

type StudentServiceSuite struct {
	suite.Suite
	DB             *pgxpool.Pool
	StudentService *students.StudentService
}

func TestStudentServiceSuite(t *testing.T) {
	suite.Run(t, new(StudentServiceSuite))
}

func (s *StudentServiceSuite) SetupSuite() {
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
	s.StudentService = students.SpawnStudentService(database)
}

func (s *StudentServiceSuite) TearDownSuite() {
	s.DB.Close()
}

func (s *StudentServiceSuite) SetupTest() {
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

func (s *StudentServiceSuite) TearDownTest() {
	ctx := context.Background()
	_, err := s.DB.Exec(ctx, "DELETE FROM registrations; DELETE FROM students; DELETE FROM teachers;")
	if err != nil {
		s.T().Fatal("Failed to clean up test data:", err)
	}
}

func (s *StudentServiceSuite) TestChangeStudentSuspensionStatusSuspend() {
	ctx := context.Background()
	studentEmail := "student1@example.com"

	student, err := s.StudentService.ChangeStudentSuspensionStatus(ctx, studentEmail, true)
	s.NoError(err)
	s.True(student.IsSuspended, "The student should be suspended")
}

func (s *StudentServiceSuite) TestChangeStudentSuspensionStatusUnsuspend() {
	ctx := context.Background()
	studentEmail := "student2@example.com"

	student, err := s.StudentService.ChangeStudentSuspensionStatus(ctx, studentEmail, false)
	s.NoError(err)
	s.False(student.IsSuspended, "The student should not be suspended")
}

func (s *StudentServiceSuite) TestChangeStudentSuspensionStatusStudentNotFound() {
	ctx := context.Background()
	studentEmail := "nonexistentstudent@example.com"

	_, err := s.StudentService.ChangeStudentSuspensionStatus(ctx, studentEmail, true)
	s.Error(err, "Should return an error for a non-existing student")
}
