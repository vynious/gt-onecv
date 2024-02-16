package registrations

import (
	"github.com/gin-gonic/gin"
	"github.com/vynious/gt-onecv/utils"
	"log"
	"net/http"
)

type RegistrationHandler struct {
	*RegistrationService
}

type RegisterStudentsRequest struct {
	Teacher  string   `json:"teacher"`
	Students []string `json:"students"`
}

func SpawnRegistrationHandler(rs *RegistrationService) *RegistrationHandler {
	return &RegistrationHandler{
		rs,
	}
}

func (rh *RegistrationHandler) RegisterStudents(c *gin.Context) {
	var requestBody RegisterStudentsRequest

	// Attempt to bind the incoming JSON payload to the struct
	if err := c.BindJSON(&requestBody); err != nil {
		log.Println("Error binding JSON:", err.Error())
		c.JSON(http.StatusBadRequest, utils.NewAPIError(utils.ErrInvalidRequestBody))
		return
	}

	// Check if the teacher's email is provided
	teacherEmail := requestBody.Teacher
	if teacherEmail == "" {
		log.Println("Error: No teacher email provided")
		c.JSON(http.StatusBadRequest, utils.NewAPIError(utils.ErrMissingTeacherEmail))
		return
	}

	// Check if the student emails list is provided and not empty
	studentEmails := requestBody.Students
	if len(studentEmails) == 0 {
		log.Println("Error: No student emails provided")
		c.JSON(http.StatusBadRequest, utils.NewAPIError(utils.ErrMissingStudentEmail))
		return
	}
	// Process each student email
	for _, studentEmail := range studentEmails {
		if err := rh.CreateRegistration(c.Request.Context(), teacherEmail, studentEmail); err != nil {
			c.JSON(http.StatusInternalServerError, utils.NewAPIError(err))
			return
		}
	}

	c.Status(http.StatusNoContent)
}

func (rh *RegistrationHandler) ViewCommonStudentsUnderTeachers(c *gin.Context) {
	teachers := c.QueryArray("teacher")

	if len(teachers) == 0 {
		c.JSON(http.StatusBadRequest, utils.NewAPIError(utils.ErrMissingTeacherEmail))
	}

	common, err := rh.GetCommonRegistrationsByTEmails(c.Request.Context(), teachers)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, utils.NewAPIError(err))
		return
	}

	c.JSON(http.StatusOK, gin.H{"students": common})
}
