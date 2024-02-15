package api

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"github.com/vynious/gt-onecv/db"
	"github.com/vynious/gt-onecv/domains/notifications"
	"github.com/vynious/gt-onecv/domains/registration"
	"github.com/vynious/gt-onecv/domains/students"
	"log"
	"net/http"
	"os"
)

type Server struct {
	Router *gin.Engine
	rh     *registration.RegistrationHandler
	nh     *notifications.NotificationHandler
	sh     *students.StudentHandler
}

func SpawnServer() *Server {
	engine := gin.Default() // Initialize the Gin router

	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalf("failed to connect to psql %v", err)
	}
	database := db.SpawnRepository(conn)
	notificationService := notifications.SpawnNotificationService(database)
	notificationHandler := notifications.SpawnNotificationHandler(notificationService)

	studentService := students.SpawnStudentService(database)
	studentHandler := students.SpawnStudentHandler(studentService)

	registrationService := registration.SpawnRegistrationService(database)
	registrationHandler := registration.SpawnRegistrationHandler(registrationService)

	return &Server{
		Router: engine,
		nh:     notificationHandler,
		sh:     studentHandler,
		rh:     registrationHandler,
	}
}

func (s *Server) MountHandlers() {
	api := s.router.Group("/api")

	api.POST("/register", s.rh.RegisterStudents)
	api.GET("/commonstudents", s.rh.ViewCommonStudentsUnderTeachers)
	api.POST("/suspend", s.sh.SuspendStudent)
	api.POST("/retrievefornotifications", s.nh.PublishNotification)

}

func (s *Server) Start(addr string) error {
	return s.router.Run(addr)
}

func (s *Server) Shutdown(ctx context.Context) error {
	if httpServer, ok := s.router.(*http.Server); ok {
		return httpServer.Shutdown(ctx)
	}
	return fmt.Errorf("router is not an *http.Server instance")
}
