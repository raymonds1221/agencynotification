package controller

import (
	"net/http"

	"github.com/Ubidy/Ubidy_AgencyNotificationAPI/delivery/usecase"
	"github.com/gin-gonic/gin"
)

// ClarificationController implementation of clarification controller
type ClarificationController struct{}

// AddPostTopicActivity api for creating activity when agency post a new topic to clarification
func (cc *ClarificationController) AddPostTopicActivity(ci usecase.ClarificationInteractor) func(*gin.Context) {
	return func(c *gin.Context) {
		supplierID := c.PostForm("supplierID")
		supplierName := c.PostForm("supplierName")
		clientID := c.PostForm("clientID")
		clientName := c.PostForm("clientName")
		auctionID := c.PostForm("auctionID")
		auctionNumber := c.PostForm("auctionNumber")

		activity, err := ci.AddPostTopicActivity(supplierID, supplierName, clientID, clientName, auctionID, auctionNumber)

		if err != nil {
			c.JSON(http.StatusNoContent, gin.H{
				"error":           "Unable to create activity for agency creating new topic",
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

// AddPostQuestionActivity api for creating activity when agency added question to the clarification discussion
func (cc *ClarificationController) AddPostQuestionActivity(ci usecase.ClarificationInteractor) func(*gin.Context) {
	return func(c *gin.Context) {
		supplierID := c.PostForm("supplierID")
		supplierName := c.PostForm("supplierName")
		clientID := c.PostForm("clientID")
		clientName := c.PostForm("clientName")
		auctionID := c.PostForm("auctionID")
		auctionNumber := c.PostForm("auctionNumber")

		activity, err := ci.AddPostQuestionActivity(supplierID, supplierName, clientID, clientName, auctionID, auctionNumber)

		if err != nil {
			c.JSON(http.StatusNoContent, gin.H{
				"error":           "Unable to create activity for agency posted a question to discussion",
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

// AddPostTopicSuccessFeeActivity api for creating activity when agency post a new topic to clarification
func (cc *ClarificationController) AddPostTopicSuccessFeeActivity(ci usecase.ClarificationInteractor) func(*gin.Context) {
	return func(c *gin.Context) {
		supplierID := c.PostForm("supplierID")
		supplierName := c.PostForm("supplierName")
		clientID := c.PostForm("clientID")
		clientName := c.PostForm("clientName")
		successFeeID := c.PostForm("successFeeID")
		successFeeNumber := c.PostForm("successFeeNumber")

		activity, err := ci.AddPostTopicSuccessFeeActivity(supplierID, supplierName, clientID, clientName, successFeeID, successFeeNumber)

		if err != nil {
			c.JSON(http.StatusNoContent, gin.H{
				"error":           "Unable to create activity for agency creating new topic",
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

// AddPostQuestionSuccessFeeActivity api for creating activity when agency added question to the clarification discussion
func (cc *ClarificationController) AddPostQuestionSuccessFeeActivity(ci usecase.ClarificationInteractor) func(*gin.Context) {
	return func(c *gin.Context) {
		supplierID := c.PostForm("supplierID")
		supplierName := c.PostForm("supplierName")
		clientID := c.PostForm("clientID")
		clientName := c.PostForm("clientName")
		successFeeID := c.PostForm("successFeeID")
		successFeeNumber := c.PostForm("successFeeNumber")

		activity, err := ci.AddPostQuestionSuccessFeeActivity(supplierID, supplierName, clientID, clientName, successFeeID, successFeeNumber)

		if err != nil {
			c.JSON(http.StatusNoContent, gin.H{
				"error":           "Unable to create activity for agency posted a question to discussion",
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
