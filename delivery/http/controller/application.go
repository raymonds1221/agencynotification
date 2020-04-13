package controller

import (
	"net/http"

	"github.com/Ubidy/Ubidy_AgencyNotificationAPI/delivery/usecase"
	"github.com/gin-gonic/gin"
)

// ApplicationController implementation of application controller
type ApplicationController struct{}

// AddSubmitApplicationActivity api for creating activity when agency submitted application
func (ac *ApplicationController) AddSubmitApplicationActivity(ai usecase.ApplicationInteractor) func(*gin.Context) {
	return func(c *gin.Context) {
		supplierID := c.PostForm("supplierID")
		supplierName := c.PostForm("supplierName")
		clientID := c.PostForm("clientID")
		clientName := c.PostForm("clientName")
		auctionID := c.PostForm("auctionID")
		auctionNumber := c.PostForm("auctionNumber")

		activity, err := ai.AddSubmitApplicationActivity(supplierID, supplierName, clientID, clientName, auctionID, auctionNumber)

		if err != nil {
			c.JSON(http.StatusNoContent, gin.H{
				"error":           "Unable to create activity for agency submitted an application",
				"internalMessage": err.Error(),
			})
			return
		}

		c.JSON(http.StatusCreated, gin.H{
			"status":   "ok",
			"activity": activity,
		})
	}
}

// AddWithdrawApplicationActivity api for creating activity when agency withdraw his/her application
func (ac *ApplicationController) AddWithdrawApplicationActivity(ai usecase.ApplicationInteractor) func(*gin.Context) {
	return func(c *gin.Context) {
		supplierID := c.PostForm("supplierID")
		supplierName := c.PostForm("supplierName")
		clientID := c.PostForm("clientID")
		clientName := c.PostForm("clientName")
		auctionID := c.PostForm("auctionID")
		auctionNumber := c.PostForm("auctionNumber")

		activity, err := ai.AddWithdrawApplicationActivity(supplierID, supplierName, clientID, clientName, auctionID, auctionNumber)

		if err != nil {
			c.JSON(http.StatusNoContent, gin.H{
				"error":           "Unable to create activity for agency withdraw his/her application",
				"internalMessage": err.Error(),
			})
			return
		}

		c.JSON(http.StatusCreated, gin.H{
			"status":   "ok",
			"activity": activity,
		})
	}
}

// AddSubmitApplicationSuccessFeeActivity api for creating activity when agency submitted application
func (ac *ApplicationController) AddSubmitApplicationSuccessFeeActivity(ai usecase.ApplicationInteractor) func(*gin.Context) {
	return func(c *gin.Context) {
		supplierID := c.PostForm("supplierID")
		supplierName := c.PostForm("supplierName")
		clientID := c.PostForm("clientID")
		clientName := c.PostForm("clientName")
		successFeeID := c.PostForm("successFeeID")
		successFeeNumber := c.PostForm("successFeeNumber")

		activity, err := ai.AddSubmitApplicationSuccessFeeActivity(supplierID, supplierName, clientID, clientName, successFeeID, successFeeNumber)

		if err != nil {
			c.JSON(http.StatusNoContent, gin.H{
				"error":           "Unable to create activity for agency submitted an application",
				"internalMessage": err.Error(),
			})
			return
		}

		c.JSON(http.StatusCreated, gin.H{
			"status":   "ok",
			"activity": activity,
		})
	}
}

// AddWithdrawApplicationSuccessFeeActivity api for creating activity when agency withdraw his/her application
func (ac *ApplicationController) AddWithdrawApplicationSuccessFeeActivity(ai usecase.ApplicationInteractor) func(*gin.Context) {
	return func(c *gin.Context) {
		supplierID := c.PostForm("supplierID")
		supplierName := c.PostForm("supplierName")
		clientID := c.PostForm("clientID")
		clientName := c.PostForm("clientName")
		successFeeID := c.PostForm("successFeeID")
		successFeeNumber := c.PostForm("successFeeNumber")

		activity, err := ai.AddWithdrawApplicationSuccessFeeActivity(supplierID, supplierName, clientID, clientName, successFeeID, successFeeNumber)

		if err != nil {
			c.JSON(http.StatusNoContent, gin.H{
				"error":           "Unable to create activity for agency withdraw his/her application",
				"internalMessage": err.Error(),
			})
			return
		}

		c.JSON(http.StatusCreated, gin.H{
			"status":   "ok",
			"activity": activity,
		})
	}
}
