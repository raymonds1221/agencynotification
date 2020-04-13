package controller

import (
	"net/http"

	"github.com/Ubidy/Ubidy_AgencyNotificationAPI/delivery/usecase"
	"github.com/gin-gonic/gin"
)

// TokenController implementation of token controller
type TokenController struct{}

// GetToken api for retrieving token
func (tc *TokenController) GetToken(ti usecase.TokenInteractor) func(*gin.Context) {
	return func(c *gin.Context) {
		userID := c.Param("userID")
		token := ti.GetToken(userID)
		unread := ti.GetUnreadNotificationCount(userID)

		c.JSON(http.StatusOK, gin.H{
			"status":           "ok",
			"token":            token,
			"unreadActivities": unread,
		})
	}
}
