package repository

import (
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/Microsoft/ApplicationInsights-Go/appinsights"
	commentModel "github.com/Ubidy/Ubidy_AgencyNotificationAPI/domain/comment"
	stream "gopkg.in/GetStream/stream-go2.v1"
)

// Comment implementation of comment repository
type Comment struct {
	client          *stream.Client
	employerDB      *sql.DB
	auctionDB       *sql.DB
	telemetryClient appinsights.TelemetryClient
}

// NewCommentRepository create new instance of comment repository
func NewCommentRepository(client *stream.Client, employerDB *sql.DB, auctionDB *sql.DB, telemetryClient appinsights.TelemetryClient) *Comment {
	return &Comment{
		client:          client,
		employerDB:      employerDB,
		auctionDB:       auctionDB,
		telemetryClient: telemetryClient,
	}
}

// AddCommentToCandidateActivity create an activity when the agency add comment on the candidate viewer
func (c *Comment) AddCommentToCandidateActivity(comment commentModel.AuctionComment) (stream.Activity, error) {
	assignmentRepository, settingsRepository := c.createAssignmentAndSettingsRepository()
	assignments, _ := assignmentRepository.GetEmployerAssignmentsByAuctionID(comment.AuctionID)

	if assignmentRepository.IsApprovedAuctionStatus(comment.AuctionID) {
		for _, assignment := range assignments {
			userID := strings.ToLower(assignment.UserID)
			settings, _ := settingsRepository.GetSettingsByClientID(userID)

			if settings.Fulfillment {
				msg := "Please note, %s commented to candidate %s, for Talent Request #%s: %s, in Auction #%s (Competitive)."
				c.sendNotificationToEmployer(comment.Comment, userID, comment.AuctionID, comment.AuctionNumber, "auction", msg)
			}
		}
	}

	return stream.Activity{}, nil
}

// AddCommentToCandidateSuccessFeeActivity create an activity when the agency add comment on the candidate viewer for success fee
func (c *Comment) AddCommentToCandidateSuccessFeeActivity(comment commentModel.SuccessFeeComment) (stream.Activity, error) {
	assignmentRepository, settingsRepository := c.createAssignmentAndSettingsRepository()
	assignments, _ := assignmentRepository.GetEmployerAssignmentsBySuccessFeeID(comment.SuccessFeeID)

	if assignmentRepository.IsApprovedSuccessFeeStatus(comment.SuccessFeeID) {
		for _, assignment := range assignments {
			userID := strings.ToLower(assignment.UserID)
			settings, _ := settingsRepository.GetSettingsByClientID(userID)

			if settings.Fulfillment {
				msg := "%s commented on candidate %s, for %s, in Engagement#%s"
				c.sendNotificationToEmployer(comment.Comment, userID, comment.SuccessFeeID, comment.SuccessFeeNumber, "successFee", msg)
			}
		}
	}

	return stream.Activity{}, nil
}

func (c *Comment) createAssignmentAndSettingsRepository() (*Assignment, *Settings) {
	settingsRepository := NewSettingsRepository(c.employerDB, nil, c.telemetryClient)
	assignmentRepository := NewAssignmentRepository(c.auctionDB, c.telemetryClient)

	return assignmentRepository, settingsRepository
}

func (c *Comment) sendNotificationToEmployer(comment commentModel.Comment, userID string, id string, number string, t string, msg string) error {
	helper := NewHelper(c.telemetryClient)

	agencyData, _ := helper.createJSONMarshal(comment.AgencyUserID, comment.AgencyName, "supplier")
	candidateData, _ := helper.createJSONMarshal(comment.CandidateID, comment.CandidateName, "candidate")
	talentRequestData, _ := helper.createJSONMarshal(comment.TalentRequestID, comment.TalentRequestNumber, "talentRequest")
	data, _ := helper.createJSONMarshal(id, number, t)

	employerNotificationFeed := c.client.NotificationFeed("employernotification", userID)

	_, err := employerNotificationFeed.AddActivity(stream.Activity{
		Actor:     employerNotificationFeed.ID(),
		Verb:      "create",
		Object:    fmt.Sprintf("%s:%s", t, id),
		ForeignID: employerNotificationFeed.ID(),
		Time:      stream.Time{time.Now().UTC()},
		Extra: map[string]interface{}{
			"employer": fmt.Sprintf("employer:%s", userID),
			"agency":   fmt.Sprintf("agency:%s", comment.AgencyUserID),
			"content":  fmt.Sprintf(msg, agencyData, candidateData, comment.JobTitle, data),
			"category": "Fulfillment",
			"subcategory": map[string]string{
				"type":   "Candidate",
				"status": "New Comment",
				"data":   fmt.Sprintf("%s", talentRequestData),
			},
		},
	})

	if err != nil {
		c.telemetryClient.TrackException(err)
		return err
	}

	return nil
}
