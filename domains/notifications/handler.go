package notifications

import (
	"github.com/gin-gonic/gin"
	"github.com/vynious/gt-onecv/utils"
	"log"
	"net/http"
)

type NotificationHandler struct {
	*NotificationService
}

func SpawnNotificationHandler(ns *NotificationService) *NotificationHandler {
	return &NotificationHandler{
		ns,
	}
}

func (nh *NotificationHandler) RetrieveStudentsForNotification(c *gin.Context) {

	var requestBody struct {
		Teacher      string `json:"teacher"`
		Notification string `json:"notification"`
	}

	if err := c.BindJSON(&requestBody); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, utils.NewAPIError(utils.ErrInvalidRequestBody))
		return
	}
	teacherEmail := requestBody.Teacher
	content := requestBody.Notification

	recipients, err := nh.GetNotifiableStudents(c.Request.Context(), teacherEmail, content)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, utils.NewAPIError(err))
		return
	}
	c.JSON(http.StatusOK, gin.H{"recipients": recipients})
}
