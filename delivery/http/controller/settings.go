package controller

import (
	"net/http"
	"strconv"

	"github.com/Ubidy/Ubidy_AgencyNotificationAPI/delivery/usecase"
	"github.com/gin-gonic/gin"
)

// SettingsController controller implementation for settings
type SettingsController struct{}

// GetSettingsByClientID api for retrieving settings by client id
func (sc *SettingsController) GetSettingsByClientID(si usecase.SettingsInteractor) func(*gin.Context) {
	return func(c *gin.Context) {
		clientID := c.Param("clientID")
		settings, err := si.GetSettingsByClientID(clientID)

		if err != nil {
			c.JSON(http.StatusNoContent, gin.H{
				"error":           "Unable to get settings by clientID",
				"internalMessage": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"status":   "ok",
			"settings": settings,
		})
	}
}

// GetSettingsBySupplierID api for retrieving settings by supplier id
func (sc *SettingsController) GetSettingsBySupplierID(si usecase.SettingsInteractor) func(*gin.Context) {
	return func(c *gin.Context) {
		supplierID := c.Param("supplierID")
		settings, err := si.GetSettingsBySupplierID(supplierID)

		if err != nil {
			c.JSON(http.StatusNoContent, gin.H{
				"error":           "Unable to get settings by clientID",
				"internalMessage": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"status":   "ok",
			"settings": settings,
		})
	}
}

// CreateOrUpdateSettings api for creating or updating settings
func (sc *SettingsController) CreateOrUpdateSettings(si usecase.SettingsInteractor) func(*gin.Context) {
	return func(c *gin.Context) {
		userID := c.PostForm("userID")
		var auctionsScheduling, prequalification, applications, clarifications, bidding,
			awarding, fulfillment, payments, ubidy, messages, users bool = true, true, true, true, true, true, true, true, true, true, true

		if c.PostForm("auctionsScheduling") != "" {
			auctionsScheduling, _ = strconv.ParseBool(c.PostForm("auctionsScheduling"))
		}

		if c.PostForm("prequalification") != "" {
			prequalification, _ = strconv.ParseBool(c.PostForm("prequalification"))
		}

		if c.PostForm("applications") != "" {
			applications, _ = strconv.ParseBool(c.PostForm("applications"))
		}

		if c.PostForm("clarifications") != "" {
			clarifications, _ = strconv.ParseBool(c.PostForm("clarifications"))
		}

		if c.PostForm("bidding") != "" {
			bidding, _ = strconv.ParseBool(c.PostForm("bidding"))
		}

		if c.PostForm("awarding") != "" {
			awarding, _ = strconv.ParseBool(c.PostForm("awarding"))
		}

		if c.PostForm("fulfillment") != "" {
			fulfillment, _ = strconv.ParseBool(c.PostForm("fulfillment"))
		}

		if c.PostForm("payments") != "" {
			payments, _ = strconv.ParseBool(c.PostForm("payments"))
		}

		if c.PostForm("ubidy") != "" {
			ubidy, _ = strconv.ParseBool(c.PostForm("ubidy"))
		}

		if c.PostForm("messages") != "" {
			messages, _ = strconv.ParseBool(c.PostForm("messages"))
		}

		if c.PostForm("users") != "" {
			users, _ = strconv.ParseBool(c.PostForm("users"))
		}

		activeSettings := c.PostForm("settings")

		result := si.CreateOrUpdateSettings(userID, auctionsScheduling, prequalification, applications, clarifications, bidding, awarding, fulfillment, payments, ubidy, messages, users, activeSettings)

		if !result {
			c.JSON(http.StatusNoContent, gin.H{
				"error":           "Unable to create or update settings",
				"internalMessage": "Unable to create or update settings",
			})
			return
		}

		c.JSON(http.StatusCreated, gin.H{
			"status": "ok",
		})
	}
}
