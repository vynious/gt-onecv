package registration

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type RegistrationHandler struct {
	*RegistrationService
}

type RegisterStudentsRequest struct {
	teacher  string
	students []string
}

func SpawnRegistrationHandler(rs *RegistrationService) *RegistrationHandler {
	return &RegistrationHandler{
		rs,
	}
}

func (rh *RegistrationHandler) RegisterStudents(c *gin.Context) {
	var requestBody RegisterStudentsRequest

	if err := c.BindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid request body"})
	}
	studentEmails := requestBody.students
	teacherEmail := requestBody.teacher

	for i := range studentEmails {
		_, err := rh.CreateRegistration(c.Request.Context(), teacherEmail, studentEmails[i])
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		}
	}

	c.JSON(http.StatusNoContent, gin.H{"message": "successfully registered students"})
}

func (rh *RegistrationHandler) ViewCommonStudentsUnderTeachers(c *gin.Context) {
	teachers := c.QueryArray("teacher")

	common, err := rh.GetCommonRegistrationsByTEmails(c.Request.Context(), teachers)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{"students": common})
}
