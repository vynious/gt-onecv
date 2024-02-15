package students

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type StudentHandler struct {
	*StudentService
}

type SuspendStudentRequest struct {
	student string
}

func SpawnStudentHandler(ss *StudentService) *StudentHandler {
	return &StudentHandler{
		ss,
	}
}

func (sh *StudentHandler) SuspendStudent(c *gin.Context) {
	var requestBody SuspendStudentRequest
	if err := c.BindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid request body"})
	}
	studentEmail := requestBody.student
	_, err := sh.ChangeStudentSuspensionStatus(c.Request.Context(), studentEmail, true)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
	}
	c.JSON(http.StatusNoContent, gin.H{"message": "successfully suspended student"})
}
