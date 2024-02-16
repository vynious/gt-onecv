package api

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/vynious/gt-onecv/db"
	"github.com/vynious/gt-onecv/domains/notifications"
	"github.com/vynious/gt-onecv/domains/registrations"
	"github.com/vynious/gt-onecv/domains/students"
	"log"
	"os"
)

type Server struct {
	Router *gin.Engine
	rh     *registrations.RegistrationHandler
	nh     *notifications.NotificationHandler
	sh     *students.StudentHandler
}

func SpawnServer() *Server {
	engine := gin.Default()
	conn, err := pgxpool.New(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalf("failed to connect to psql %v", err)
	}
	database := db.SpawnRepository(conn)
	notificationService := notifications.SpawnNotificationService(database)
	notificationHandler := notifications.SpawnNotificationHandler(notificationService)

	studentService := students.SpawnStudentService(database)
	studentHandler := students.SpawnStudentHandler(studentService)

	registrationService := registrations.SpawnRegistrationService(database)
	registrationHandler := registrations.SpawnRegistrationHandler(registrationService)

	return &Server{
		Router: engine,
		nh:     notificationHandler,
		sh:     studentHandler,
		rh:     registrationHandler,
	}
}

func (s *Server) MountHandlers() {
	api := s.Router.Group("/api")

	api.POST("/register", s.rh.RegisterStudents)
	api.GET("/commonstudents", s.rh.ViewCommonStudentsUnderTeachers)
	api.POST("/suspend", s.sh.SuspendStudent)
	api.POST("/retrievefornotifications", s.nh.RetrieveStudentsForNotification)

}

func (s *Server) Start(addr string) error {
	return s.Router.Run(addr)
}
