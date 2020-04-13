package controller

import (
	"net/http"

	"github.com/Ubidy/Ubidy_AgencyNotificationAPI/delivery/usecase"
	"github.com/gin-gonic/gin"
)

// FulfillmentController implementation of fulfillment controller
type FulfillmentController struct{}

// AddNewCandidateActivity api for creating new activity when agency submitted a candidate
func (fc *FulfillmentController) AddNewCandidateActivity(fi usecase.FulfillmentInteractor) func(*gin.Context) {
	return func(c *gin.Context) {
		supplierID := c.PostForm("supplierID")
		supplierName := c.PostForm("supplierName")
		clientID := c.PostForm("clientID")
		clientName := c.PostForm("clientName")
		candidateID := c.PostForm("candidateID")
		candidateName := c.PostForm("candidateName")
		talentRequestID := c.PostForm("talentRequestID")
		talentRequestNumber := c.PostForm("talentRequestNumber")
		auctionID := c.PostForm("auctionID")
		auctionNumber := c.PostForm("auctionNumber")
		jobTitle := c.PostForm("jobTitle")

		activity, err := fi.AddNewCandidateActivity(supplierID, supplierName, clientID, clientName, candidateID, candidateName, talentRequestID, talentRequestNumber, auctionID, auctionNumber, jobTitle)

		if err != nil {
			c.JSON(http.StatusNoContent, gin.H{
				"error":           "Unable to create activity when agency submitted a candidate",
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

// AddUpdateCandidateActivity api for creating new activity when agency updated a candidate
func (fc *FulfillmentController) AddUpdateCandidateActivity(fi usecase.FulfillmentInteractor) func(*gin.Context) {
	return func(c *gin.Context) {
		supplierID := c.PostForm("supplierID")
		supplierName := c.PostForm("supplierName")
		clientID := c.PostForm("clientID")
		clientName := c.PostForm("clientName")
		candidateID := c.PostForm("candidateID")
		candidateName := c.PostForm("candidateName")
		talentRequestID := c.PostForm("talentRequestID")
		talentRequestNumber := c.PostForm("talentRequestNumber")
		auctionID := c.PostForm("auctionID")
		auctionNumber := c.PostForm("auctionNumber")
		jobTitle := c.PostForm("jobTitle")

		activity, err := fi.AddUpdateCandidateActivity(supplierID, supplierName, clientID, clientName, candidateID, candidateName, talentRequestID, talentRequestNumber, auctionID, auctionNumber, jobTitle)

		if err != nil {
			c.JSON(http.StatusNoContent, gin.H{
				"error":           "Unable to create activity when agency updated a candidate",
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

// AddCandidateSubmission3DaysIdleActivity api for creating new activity when agency is idle for 3 days
func (fc *FulfillmentController) AddCandidateSubmission3DaysIdleActivity(fi usecase.FulfillmentInteractor) func(*gin.Context) {
	return func(c *gin.Context) {
		supplierID := c.PostForm("supplierID")
		clientID := c.PostForm("clientID")
		clientName := c.PostForm("clientName")
		auctionID := c.PostForm("auctionID")
		auctionNumber := c.PostForm("auctionNumber")
		agencyTenantID := c.PostForm("agencyTenantID")
		employerTenantID := c.PostForm("employerTenantID")

		activity, err := fi.AddCandidateSubmission3DaysIdleActivity(supplierID, clientID, clientName, auctionID, auctionNumber, agencyTenantID, employerTenantID)

		if err != nil {
			c.JSON(http.StatusNoContent, gin.H{
				"error":           "Unable to create activity when agency is idle for 3 days",
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

// AddCandidateSubmission10DaysIdleActivity api for creating new activity when agency is idle for 10 days
func (fc *FulfillmentController) AddCandidateSubmission10DaysIdleActivity(fi usecase.FulfillmentInteractor) func(*gin.Context) {
	return func(c *gin.Context) {
		supplierID := c.PostForm("supplierID")
		clientID := c.PostForm("clientID")
		clientName := c.PostForm("clientName")
		auctionID := c.PostForm("auctionID")
		auctionNumber := c.PostForm("auctionNumber")
		agencyTenantID := c.PostForm("agencyTenantID")
		employerTenantID := c.PostForm("employerTenantID")

		activity, err := fi.AddCandidateSubmission10DaysIdleActivity(supplierID, clientID, clientName, auctionID, auctionNumber, agencyTenantID, employerTenantID)

		if err != nil {
			c.JSON(http.StatusNoContent, gin.H{
				"error":           "Unable to create activity when agency is idle for 10 days",
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

// AddCandidateSubmission14DaysIdleActivity api for creating new activity when agency is idle for 14 days
func (fc *FulfillmentController) AddCandidateSubmission14DaysIdleActivity(fi usecase.FulfillmentInteractor) func(*gin.Context) {
	return func(c *gin.Context) {
		supplierID := c.PostForm("supplierID")
		clientID := c.PostForm("clientID")
		clientName := c.PostForm("clientName")
		auctionID := c.PostForm("auctionID")
		auctionNumber := c.PostForm("auctionNumber")
		agencyTenantID := c.PostForm("agencyTenantID")
		employerTenantID := c.PostForm("employerTenantID")

		activity, err := fi.AddCandidateSubmission14DaysIdleActivity(supplierID, clientID, clientName, auctionID, auctionNumber, agencyTenantID, employerTenantID)

		if err != nil {
			c.JSON(http.StatusNoContent, gin.H{
				"error":           "Unable to create activity when agency is idle for 14 days",
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

// AddNewCandidateSuccessFeeActivity api for creating new activity when agency submitted a candidate
func (fc *FulfillmentController) AddNewCandidateSuccessFeeActivity(fi usecase.FulfillmentInteractor) func(*gin.Context) {
	return func(c *gin.Context) {
		supplierID := c.PostForm("supplierID")
		supplierName := c.PostForm("supplierName")
		clientID := c.PostForm("clientID")
		clientName := c.PostForm("clientName")
		candidateID := c.PostForm("candidateID")
		candidateName := c.PostForm("candidateName")
		talentRequestID := c.PostForm("talentRequestID")
		talentRequestNumber := c.PostForm("talentRequestNumber")
		successFeeID := c.PostForm("successFeeID")
		successFeeNumber := c.PostForm("successFeeNumber")
		jobTitle := c.PostForm("jobTitle")

		activity, err := fi.AddNewCandidateSuccessFeeActivity(supplierID, supplierName, clientID, clientName, candidateID, candidateName, talentRequestID, talentRequestNumber, successFeeID, successFeeNumber, jobTitle)

		if err != nil {
			c.JSON(http.StatusNoContent, gin.H{
				"error":           "Unable to create activity when agency submitted a candidate",
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

// AddUpdateCandidateSuccessFeeActivity api for creating new activity when agency updated a candidate
func (fc *FulfillmentController) AddUpdateCandidateSuccessFeeActivity(fi usecase.FulfillmentInteractor) func(*gin.Context) {
	return func(c *gin.Context) {
		supplierID := c.PostForm("supplierID")
		supplierName := c.PostForm("supplierName")
		clientID := c.PostForm("clientID")
		clientName := c.PostForm("clientName")
		candidateID := c.PostForm("candidateID")
		candidateName := c.PostForm("candidateName")
		talentRequestID := c.PostForm("talentRequestID")
		talentRequestNumber := c.PostForm("talentRequestNumber")
		successFeeID := c.PostForm("successFeeID")
		successFeeNumber := c.PostForm("successFeeNumber")
		jobTitle := c.PostForm("jobTitle")
		candidateStatus := c.PostForm("candidateStatus")

		activity, err := fi.AddUpdateCandidateSuccessFeeActivity(supplierID, supplierName, clientID, clientName, candidateID, candidateName, talentRequestID, talentRequestNumber, successFeeID, successFeeNumber, jobTitle, candidateStatus)

		if err != nil {
			c.JSON(http.StatusNoContent, gin.H{
				"error":           "Unable to create activity when agency updated a candidate",
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

// AddCandidateSubmission3DaysIdleSuccessFeeActivity api for creating new activity when agency is idle for 3 days
func (fc *FulfillmentController) AddCandidateSubmission3DaysIdleSuccessFeeActivity(fi usecase.FulfillmentInteractor) func(*gin.Context) {
	return func(c *gin.Context) {
		supplierID := c.PostForm("supplierID")
		clientID := c.PostForm("clientID")
		clientName := c.PostForm("clientName")
		successFeeID := c.PostForm("successFeeID")
		successFeeNumber := c.PostForm("successFeeNumber")
		agencyTenantID := c.PostForm("agencyTenantID")
		employerTenantID := c.PostForm("employerTenantID")

		activity, err := fi.AddCandidateSubmission3DaysIdleSuccessFeeActivity(supplierID, clientID, clientName, successFeeID, successFeeNumber, agencyTenantID, employerTenantID)

		if err != nil {
			c.JSON(http.StatusNoContent, gin.H{
				"error":           "Unable to create activity when agency is idle for 3 days",
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

// AddCandidateSubmission10DaysIdleSuccessFeeActivity api for creating new activity when agency is idle for 10 days
func (fc *FulfillmentController) AddCandidateSubmission10DaysIdleSuccessFeeActivity(fi usecase.FulfillmentInteractor) func(*gin.Context) {
	return func(c *gin.Context) {
		supplierID := c.PostForm("supplierID")
		clientID := c.PostForm("clientID")
		clientName := c.PostForm("clientName")
		successFeeID := c.PostForm("successFeeID")
		successFeeNumber := c.PostForm("successFeeNumber")
		agencyTenantID := c.PostForm("agencyTenantID")
		employerTenantID := c.PostForm("employerTenantID")

		activity, err := fi.AddCandidateSubmission10DaysIdleSuccessFeeActivity(supplierID, clientID, clientName, successFeeID, successFeeNumber, agencyTenantID, employerTenantID)

		if err != nil {
			c.JSON(http.StatusNoContent, gin.H{
				"error":           "Unable to create activity when agency is idle for 10 days",
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

// AddCandidateSubmission14DaysIdleSuccessFeeActivity api for creating new activity when agency is idle for 14 days
func (fc *FulfillmentController) AddCandidateSubmission14DaysIdleSuccessFeeActivity(fi usecase.FulfillmentInteractor) func(*gin.Context) {
	return func(c *gin.Context) {
		supplierID := c.PostForm("supplierID")
		clientID := c.PostForm("clientID")
		clientName := c.PostForm("clientName")
		successFeeID := c.PostForm("successFeeID")
		successFeeNumber := c.PostForm("successFeeNumber")
		agencyTenantID := c.PostForm("agencyTenantID")
		employerTenantID := c.PostForm("employerTenantID")

		activity, err := fi.AddCandidateSubmission14DaysIdleSuccessFeeActivity(supplierID, clientID, clientName, successFeeID, successFeeNumber, agencyTenantID, employerTenantID)

		if err != nil {
			c.JSON(http.StatusNoContent, gin.H{
				"error":           "Unable to create activity when agency is idle for 14 days",
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
