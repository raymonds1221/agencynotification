package repository

import (
	"fmt"
	"time"

	"database/sql"
	"strings"

	"github.com/Microsoft/ApplicationInsights-Go/appinsights"
	"github.com/google/uuid"
	stream "gopkg.in/GetStream/stream-go2.v1"
)

// Application implementation for application repository
type Application struct {
	client          *stream.Client
	employerDB      *sql.DB
	auctionDB       *sql.DB
	telemetryClient appinsights.TelemetryClient
}

// NewApplicationRepository create new instance of application repository
func NewApplicationRepository(client *stream.Client, employerDB *sql.DB, auctionDB *sql.DB, telemetryClient appinsights.TelemetryClient) *Application {
	return &Application{
		client:          client,
		employerDB:      employerDB,
		auctionDB:       auctionDB,
		telemetryClient: telemetryClient,
	}
}

// AddSubmitApplicationActivity create an activity when agency submitted application
func (a *Application) AddSubmitApplicationActivity(supplierID string, supplierName string, clientID string, clientName string, auctionID string, auctionNumber string) (stream.Activity, error) {
	helper := NewHelper(a.telemetryClient)

	supplierData, _ := helper.createJSONMarshal(supplierID, supplierName, "supplier")
	auctionData, _ := helper.createJSONMarshal(auctionID, auctionNumber, "auction")

	assignmentRepository, settingsRepository := a.createAssignmentAndSettingsRepository()
	assignments, _ := assignmentRepository.GetEmployerAssignmentsByAuctionID(auctionID)

	if assignmentRepository.IsApprovedAuctionStatus(auctionID) {
		for _, assignment := range assignments {
			userID := strings.ToLower(assignment.UserID)
			settings, _ := settingsRepository.GetSettingsByClientID(userID)

			if settings.Applications {
				verb := "apply"
				object := fmt.Sprintf("auction:%s", auctionID)
				content := fmt.Sprintf("Congratulations, Agency %s applied to your Auction #%s (Competitive).", supplierData, auctionData)
				category := "Application"
				subcategory := map[string]string{
					"type":   "Application",
					"status": "Pending",
				}

				a.sendNotificationToEmployer(userID, verb, object, supplierID, content, category, subcategory)
			}
		}
	}

	return stream.Activity{}, nil
}

// AddWithdrawApplicationActivity create an activity when ageny withdraw his/her application
func (a *Application) AddWithdrawApplicationActivity(supplierID string, supplierName string, clientID string, clientName string, auctionID string, auctionNumber string) (stream.Activity, error) {
	helper := NewHelper(a.telemetryClient)

	supplierData, _ := helper.createJSONMarshal(supplierID, supplierName, "supplier")
	auctionData, _ := helper.createJSONMarshal(auctionID, auctionNumber, "auction")

	assignmentRepository, settingsRepository := a.createAssignmentAndSettingsRepository()
	assignments, _ := assignmentRepository.GetEmployerAssignmentsByAuctionID(auctionID)

	if assignmentRepository.IsApprovedAuctionStatus(auctionID) {
		for _, assignment := range assignments {
			userID := strings.ToLower(assignment.UserID)
			settings, _ := settingsRepository.GetSettingsByClientID(userID)

			if settings.Applications {
				verb := "withdraw"
				object := fmt.Sprintf("auction:%s", auctionID)
				content := fmt.Sprintf("Please note, Agency %s withdrew to your Auction #%s (Competitive).", supplierData, auctionData)
				category := "Application"
				subcategory := map[string]string{
					"type":   "Application",
					"status": "Withdrew",
				}

				a.sendNotificationToEmployer(userID, verb, object, supplierID, content, category, subcategory)
			}
		}
	}

	return stream.Activity{}, nil
}

// AddSubmitApplicationSuccessFeeActivity create an activity when agency submitted application
func (a *Application) AddSubmitApplicationSuccessFeeActivity(supplierID string, supplierName string, clientID string, clientName string, successFeeID string, successFeeNumber string) (stream.Activity, error) {
	helper := NewHelper(a.telemetryClient)

	supplierData, _ := helper.createJSONMarshal(supplierID, supplierName, "supplier")
	successFeeData, _ := helper.createJSONMarshal(successFeeID, successFeeNumber, "successFee")

	assignmentRepository, settingsRepository := a.createAssignmentAndSettingsRepository()
	assignments, _ := assignmentRepository.GetEmployerAssignmentsBySuccessFeeID(successFeeID)

	if assignmentRepository.IsApprovedSuccessFeeStatus(successFeeID) {
		for _, assignment := range assignments {
			userID := strings.ToLower(assignment.UserID)
			settings, _ := settingsRepository.GetSettingsByClientID(userID)

			if settings.Applications {
				verb := "apply"
				object := fmt.Sprintf("successfee:%s", successFeeID)
				content := fmt.Sprintf("%s applied to your Engagement #%s", supplierData, successFeeData)
				category := "Application"
				subcategory := map[string]string{
					"type":   "Application",
					"status": "Pending",
				}

				a.sendNotificationToEmployer(userID, verb, object, supplierID, content, category, subcategory)
			}
		}
	}

	return stream.Activity{}, nil
}

// AddWithdrawApplicationSuccessFeeActivity create an activity when ageny withdraw his/her application
func (a *Application) AddWithdrawApplicationSuccessFeeActivity(supplierID string, supplierName string, clientID string, clientName string, successFeeID string, successFeeNumber string) (stream.Activity, error) {
	helper := NewHelper(a.telemetryClient)

	supplierData, _ := helper.createJSONMarshal(supplierID, supplierName, "supplier")
	successFeeData, _ := helper.createJSONMarshal(successFeeID, successFeeNumber, "successFee")

	assignmentRepository, settingsRepository := a.createAssignmentAndSettingsRepository()
	assignments, _ := assignmentRepository.GetEmployerAssignmentsBySuccessFeeID(successFeeID)

	if assignmentRepository.IsApprovedSuccessFeeStatus(successFeeID) {
		for _, assignment := range assignments {
			userID := strings.ToLower(assignment.UserID)
			settings, _ := settingsRepository.GetSettingsByClientID(userID)

			if settings.Applications {
				verb := "withdraw"
				object := fmt.Sprintf("successfee:%s", successFeeID)
				content := fmt.Sprintf("%s their application to Engagement #%s", supplierData, successFeeData)
				category := "Application"
				subcategory := map[string]string{
					"type":   "Application",
					"status": "Withdrew",
				}

				a.sendNotificationToEmployer(userID, verb, object, supplierID, content, category, subcategory)
			}
		}
	}

	return stream.Activity{}, nil
}

func (a *Application) createAssignmentAndSettingsRepository() (*Assignment, *Settings) {
	assignmentRepository := NewAssignmentRepository(a.auctionDB, a.telemetryClient)
	settingsRepository := NewSettingsRepository(a.employerDB, nil, a.telemetryClient)

	return assignmentRepository, settingsRepository
}

func (a *Application) sendNotificationToEmployer(userID string, verb string, object string, supplierID string, content string, category string, subcategory map[string]string) {
	agencyFeed := a.client.FlatFeed("agency", supplierID)
	employerNotificationFeed := a.client.NotificationFeed("employernotification", userID)
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
		a.telemetryClient.TrackException(err)
		panic(err)
	}
}
