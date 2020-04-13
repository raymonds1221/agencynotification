package repository

import (
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/Microsoft/ApplicationInsights-Go/appinsights"
	"github.com/google/uuid"
	stream "gopkg.in/GetStream/stream-go2.v1"
)

// Clarification implementation of clarification repository
type Clarification struct {
	client          *stream.Client
	employerDB      *sql.DB
	auctionDB       *sql.DB
	telemetryClient appinsights.TelemetryClient
}

// NewClarificationRepository create new instance of clarification repository
func NewClarificationRepository(client *stream.Client, employerDB *sql.DB, auctionDB *sql.DB, telemetryClient appinsights.TelemetryClient) *Clarification {
	return &Clarification{
		client:          client,
		employerDB:      employerDB,
		auctionDB:       auctionDB,
		telemetryClient: telemetryClient,
	}
}

// AddPostTopicActivity create an activity when agency posted a new topic to clarification
func (c *Clarification) AddPostTopicActivity(supplierID string, supplierName string, clientID string, clientName string, auctionID string, auctionNumber string) (stream.Activity, error) {
	helper := NewHelper(c.telemetryClient)

	supplierData, _ := helper.createJSONMarshal(supplierID, supplierName, "supplier")
	auctionData, _ := helper.createJSONMarshal(auctionID, auctionNumber, "auction")

	assignmentRepository, settingsRepository := c.createAssignmentAndSettingsRepository()
	assignments, _ := assignmentRepository.GetEmployerAssignmentsByAuctionID(auctionID)

	for _, assignment := range assignments {
		userID := strings.ToLower(assignment.UserID)
		settings, _ := settingsRepository.GetSettingsByClientID(userID)

		if settings.Clarifications {
			verb := "apply"
			object := fmt.Sprintf("auction:%s", auctionID)
			content := fmt.Sprintf("%s created a new clarification topic for your Auction #%s (Competitive).", supplierData, auctionData)
			category := "Clarification"
			subcategory := map[string]string{
				"type":   "Clarification",
				"status": "Topic",
			}

			c.sendNotificationToEmployer(userID, verb, object, supplierID, content, category, subcategory)
		}
	}

	return stream.Activity{}, nil
}

// AddPostQuestionActivity create new activity when agency posted a question to a clarification discussion
func (c *Clarification) AddPostQuestionActivity(supplierID string, supplierName string, clientID string, clientName string, auctionID string, auctionNumber string) (stream.Activity, error) {
	helper := NewHelper(c.telemetryClient)

	supplierData, _ := helper.createJSONMarshal(supplierID, supplierName, "supplier")
	auctionData, _ := helper.createJSONMarshal(auctionID, auctionNumber, "auction")

	assignmentRepository, settingsRepository := c.createAssignmentAndSettingsRepository()
	assignments, _ := assignmentRepository.GetEmployerAssignmentsByAuctionID(auctionID)

	for _, assignment := range assignments {
		userID := strings.ToLower(assignment.UserID)
		settings, _ := settingsRepository.GetSettingsByClientID(userID)

		if settings.Clarifications {
			verb := "apply"
			object := fmt.Sprintf("auction:%s", auctionID)
			content := fmt.Sprintf("Please note, Agency %s created a new clarification reply for your Auction #%s (Competitive).", supplierData, auctionData)
			category := "Clarification"
			subcategory := map[string]string{
				"type":   "Clarification",
				"status": "Replied",
			}

			c.sendNotificationToEmployer(userID, verb, object, supplierID, content, category, subcategory)
		}
	}

	return stream.Activity{}, nil
}

// AddPostTopicSuccessFeeActivity create an activity when agency posted a new topic to clarification
func (c *Clarification) AddPostTopicSuccessFeeActivity(supplierID string, supplierName string, clientID string, clientName string, successFeeID string, successFeeNumber string) (stream.Activity, error) {
	helper := NewHelper(c.telemetryClient)

	supplierData, _ := helper.createJSONMarshal(supplierID, supplierName, "supplier")
	successFeeData, _ := helper.createJSONMarshal(successFeeID, successFeeNumber, "successFee")

	assignmentRepository, settingsRepository := c.createAssignmentAndSettingsRepository()
	assignments, _ := assignmentRepository.GetEmployerAssignmentsBySuccessFeeID(successFeeID)

	for _, assignment := range assignments {
		userID := strings.ToLower(assignment.UserID)
		settings, _ := settingsRepository.GetSettingsByClientID(userID)

		if settings.Clarifications {
			verb := "apply"
			object := fmt.Sprintf("successfee:%s", successFeeID)
			content := fmt.Sprintf("%s created a clarification topic for Engagement #%s", supplierData, successFeeData)
			category := "Clarification"
			subcategory := map[string]string{
				"type":   "Clarification",
				"status": "Topic",
			}

			c.sendNotificationToEmployer(userID, verb, object, supplierID, content, category, subcategory)
		}
	}

	return stream.Activity{}, nil
}

// AddPostQuestionSuccessFeeActivity create new activity when agency posted a question to a clarification discussion
func (c *Clarification) AddPostQuestionSuccessFeeActivity(supplierID string, supplierName string, clientID string, clientName string, successFeeID string, successFeeNumber string) (stream.Activity, error) {
	helper := NewHelper(c.telemetryClient)

	supplierData, _ := helper.createJSONMarshal(supplierID, supplierName, "supplier")
	successFeeData, _ := helper.createJSONMarshal(successFeeID, successFeeNumber, "successFee")

	assignmentRepository, settingsRepository := c.createAssignmentAndSettingsRepository()
	assignments, _ := assignmentRepository.GetEmployerAssignmentsBySuccessFeeID(successFeeID)

	for _, assignment := range assignments {
		userID := strings.ToLower(assignment.UserID)
		settings, _ := settingsRepository.GetSettingsByClientID(userID)

		if settings.Clarifications {
			verb := "apply"
			object := fmt.Sprintf("successfee:%s", successFeeID)
			content := fmt.Sprintf("%s replied to a clarification topic in Engagement #%s", supplierData, successFeeData)
			category := "Clarification"
			subcategory := map[string]string{
				"type":   "Clarification",
				"status": "Replied",
			}

			c.sendNotificationToEmployer(userID, verb, object, supplierID, content, category, subcategory)
		}
	}

	return stream.Activity{}, nil
}

func (c *Clarification) createAssignmentAndSettingsRepository() (*Assignment, *Settings) {
	assignmentRepository := NewAssignmentRepository(c.auctionDB, c.telemetryClient)
	settingsRepository := NewSettingsRepository(c.employerDB, nil, c.telemetryClient)

	return assignmentRepository, settingsRepository
}

func (c *Clarification) sendNotificationToEmployer(userID string, verb string, object string, supplierID string, content string, category string, subcategory map[string]string) {
	agencyFeed := c.client.FlatFeed("agency", supplierID)
	employerNotificationFeed := c.client.NotificationFeed("employernotification", userID)

	_, err := employerNotificationFeed.AddActivity(stream.Activity{
		Actor:     agencyFeed.ID(),
		Verb:      verb,
		Object:    object,
		ForeignID: uuid.New().String(),
		Time:      stream.Time{time.Now().UTC()},
		Extra: map[string]interface{}{
			"employer":    fmt.Sprintf("employer:%s", userID),
			"agency":      fmt.Sprintf("agency:%s", supplierID),
			"content":     content,
			"category":    category,
			"subcategory": subcategory,
		},
	})

	if err != nil {
		c.telemetryClient.TrackException(err)
		panic(err)
	}
}
