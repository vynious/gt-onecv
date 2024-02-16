package students

import (
	"github.com/gin-gonic/gin"
	"github.com/vynious/gt-onecv/utils"
	"log"
	"net/http"
)

type StudentHandler struct {
	*StudentService
}

type SuspendStudentRequest struct {
	Student string `json:"student"`
}

func SpawnStudentHandler(ss *StudentService) *StudentHandler {
	return &StudentHandler{
		ss,
	}
}

func (sh *StudentHandler) SuspendStudent(c *gin.Context) {
	var requestBody SuspendStudentRequest
	if err := c.BindJSON(&requestBody); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, utils.NewAPIError(utils.ErrInvalidRequestBody))
		return
	}
	studentEmail := requestBody.Student
	if studentEmail == "" {
		log.Println("Error: No student email provided")
		c.JSON(http.StatusBadRequest, utils.NewAPIError(utils.ErrMissingStudentEmail))
		return
	}
	_, err := sh.ChangeStudentSuspensionStatus(c.Request.Context(), studentEmail, true)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, utils.NewAPIError(err))
		return
	}
	c.Status(http.StatusNoContent)
}
