package notifications

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type NotificationHandler struct {
	*NotificationService
}

type PublishNotificationRequest struct {
	teacher      string
	notification string
}

func SpawnNotificationHandler(ns *NotificationService) *NotificationHandler {
	return &NotificationHandler{
		ns,
	}
}

func (nh *NotificationHandler) PublishNotification(c *gin.Context) {
	var requestBody PublishNotificationRequest
	if err := c.BindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid request body"})
	}
	teacherEmail := requestBody.teacher
	content := requestBody.notification

	recipients, err := nh.GetNotifiableStudents(c.Request.Context(), teacherEmail, content)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
	}
	c.JSON(http.StatusOK, gin.H{"recipients": recipients})
}
