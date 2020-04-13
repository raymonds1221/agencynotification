package controller

import (
	"net/http"

	"github.com/Ubidy/Ubidy_AgencyNotificationAPI/delivery/usecase"
	"github.com/gin-gonic/gin"
)

// BiddingController implementation of bidding controller
type BiddingController struct{}

// AddPlaceBidActivity api for creating activity when the agency place a bid
func (bc *BiddingController) AddPlaceBidActivity(bi usecase.BiddingInteractor) func(*gin.Context) {
	return func(c *gin.Context) {
		supplierID := c.PostForm("supplierID")
		clientID := c.PostForm("clientID")
		auctionID := c.PostForm("auctionID")

		activity, err := bi.AddPlaceBidActivity(supplierID, clientID, auctionID)

		if err != nil {
			c.JSON(http.StatusNoContent, gin.H{
				"error":           "Unable to create activity for agency placed a bid",
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
