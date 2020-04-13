package repository

import (
	"errors"
	"fmt"
	"time"

	"database/sql"
	"net/url"
	"strings"

	"github.com/Microsoft/ApplicationInsights-Go/appinsights"
	"github.com/google/uuid"
	stream "gopkg.in/GetStream/stream-go2.v1"
)

const (
	candidatePending           = "/api/v1.0.0/employer/candidatepending"
	candidatePendingSuccessFee = "/api/v1.0.0/employer/candidatepending/successfee"
)

const (
	approvedApplicationStatusID = 2
	approvedApplicationStatus   = "Approved"
)

// Fulfillment implementation of fulfillment repository
type Fulfillment struct {
	client          *stream.Client
	employerDB      *sql.DB
	auctionDB       *sql.DB
	telemetryClient appinsights.TelemetryClient
}

// NewFulfillmentRepository create new instance of fulfillment repository
func NewFulfillmentRepository(client *stream.Client, employerDB *sql.DB, auctionDB *sql.DB, telemetryClient appinsights.TelemetryClient) *Fulfillment {
	return &Fulfillment{
		client:          client,
		employerDB:      employerDB,
		auctionDB:       auctionDB,
		telemetryClient: telemetryClient,
	}
}

// AddNewCandidateActivity create a new activity when the agency submitted candidate
func (f *Fulfillment) AddNewCandidateActivity(supplierID string, supplierName string, clientID string, clientName string, candidateID string, candidateName string, talentRequestID string, talentRequestNumber string, auctionID string, auctionNumber string, jobTitle string) (stream.Activity, error) {
	helper := NewHelper(f.telemetryClient)

	supplierData, _ := helper.createJSONMarshal(supplierID, supplierName, "supplier")
	candidateData, _ := helper.createJSONMarshal(candidateID, candidateName, "candidate")
	talentRequestData, _ := helper.createJSONMarshal(talentRequestID, fmt.Sprintf("%06s", talentRequestNumber), "talentRequest")
	auctionData, _ := helper.createJSONMarshal(auctionID, auctionNumber, "auction")

	assignmentRepository, settingsRepository := f.createAssignmentAndSettingsRepository()
	assignments, _ := assignmentRepository.GetEmployerAssignmentsByAuctionID(auctionID)

	if assignmentRepository.IsApprovedAuctionStatus(auctionID) {
		for _, assignment := range assignments {
			userID := strings.ToLower(assignment.UserID)
			settings, _ := settingsRepository.GetSettingsByClientID(userID)

			if settings.Fulfillment {
				verb := "submit"
				target := fmt.Sprintf("auction:%s", auctionID)
				content := fmt.Sprintf("Congratulations, %s submitted a new candidate %s, for Talent Request #%s: %s, for Auction #%s (Competitive).", string(supplierData), string(candidateData), string(talentRequestData), jobTitle, string(auctionData))
				category := "Fulfillment"
				subcategory := map[string]string{
					"type":   "Candidate",
					"status": "New",
				}

				f.sendNotificationtoEmployer(userID, supplierID, candidateID, verb, target, content, category, subcategory)
			}
		}

		formData := url.Values{
			"supplierID":          {supplierID},
			"supplierName":        {supplierName},
			"clientID":            {clientID},
			"clientName":          {clientName},
			"candidateID":         {candidateID},
			"candidateName":       {candidateName},
			"talentRequestID":     {talentRequestID},
			"talentRequestNumber": {talentRequestNumber},
			"auctionID":           {auctionID},
			"auctionNumber":       {auctionNumber},
			"jobTitle":            {jobTitle},
		}

		helper.post(candidatePending, formData)
	}

	return stream.Activity{}, nil
}

// AddUpdateCandidateActivity create an activity when agency updated candidate
func (f *Fulfillment) AddUpdateCandidateActivity(supplierID string, supplierName string, clientID string, clientName string, candidateID string, candidateName string, talentRequestID string, talentRequestNumber string, auctionID string, auctionNumber string, jobTitle string) (stream.Activity, error) {
	helper := NewHelper(f.telemetryClient)

	supplierData, _ := helper.createJSONMarshal(supplierID, supplierName, "supplier")
	candidateData, _ := helper.createJSONMarshal(candidateID, candidateName, "candidate")
	talentRequestData, _ := helper.createJSONMarshal(talentRequestID, fmt.Sprintf("%06s", talentRequestNumber), "talentRequest")
	auctionData, _ := helper.createJSONMarshal(auctionID, auctionNumber, "auction")

	assignmentRepository, settingsRepository := f.createAssignmentAndSettingsRepository()
	assignments, _ := assignmentRepository.GetEmployerAssignmentsByAuctionID(auctionID)

	if assignmentRepository.IsApprovedAuctionStatus(auctionID) {
		for _, assignment := range assignments {
			userID := strings.ToLower(assignment.UserID)
			settings, _ := settingsRepository.GetSettingsByClientID(userID)

			if settings.Fulfillment {
				verb := "update"
				target := fmt.Sprintf("auction:%s", auctionID)
				content := fmt.Sprintf("Please note, %s updated candidate %s, for Talent Request #%s: %s, for Auction #%s (Competitive).", string(supplierData), string(candidateData), string(talentRequestData), jobTitle, string(auctionData))
				category := "Fulfillment"
				subcategory := map[string]string{
					"type":   "Candidate",
					"status": "New",
				}

				f.sendNotificationtoEmployer(userID, supplierID, candidateID, verb, target, content, category, subcategory)
			}
		}
	}

	return stream.Activity{}, nil
}

// AddCandidateSubmission3DaysIdleActivity create a new activity when the agency is idle for 3 days
func (f *Fulfillment) AddCandidateSubmission3DaysIdleActivity(supplierID string, clientID string, clientName string, auctionID string, auctionNumber string, agencyTenantID string, employerTenantID string) (stream.Activity, error) {
	helper := NewHelper(f.telemetryClient)

	applicationStatusID, applicationStatus := f.getAgencyStatusByAuctionID(auctionID, agencyTenantID)

	if applicationStatusID != approvedApplicationStatusID && applicationStatus != approvedApplicationStatus {
		return stream.Activity{}, errors.New("AddCandidateSubmission3DaysIdleActivity: agency is not allowed!")
	}

	clientData1, _ := helper.createJSONMarshalWithUUID(clientID, clientName, "client", uuid.New().String())
	clientData2, _ := helper.createJSONMarshalWithUUID(clientID, clientName, "client", uuid.New().String())
	auctionData, _ := helper.createJSONMarshal(auctionID, auctionNumber, "auction")

	assignmentRepository := NewAssignmentRepository(f.auctionDB, f.telemetryClient)
	assignments, _ := assignmentRepository.GetAgencyAssignmentsByAuctionID(auctionID, agencyTenantID)

	if assignmentRepository.IsApprovedAuctionStatus(auctionID) {
		for _, assignment := range assignments {
			userID := strings.ToLower(assignment.UserID)

			verb := "submit"
			target := fmt.Sprintf("auction:%s", auctionID)
			content := fmt.Sprintf("Competitive Auction#%s, for %s, is waiting for a candidate submission. If you have interviewed a candidate or have a Candidate ready, please remember to submit so %s can start their review!", string(auctionData), string(clientData1), string(clientData2))
			category := "Fulfillment"
			subcategory := map[string]string{
				"type":   "Candidate",
				"status": "Required",
			}

			f.sendNotificationToAgency(userID, supplierID, clientID, verb, target, content, category, subcategory, agencyTenantID, employerTenantID)
		}
	}

	return stream.Activity{}, nil
}

// AddCandidateSubmission10DaysIdleActivity create a new activity when the agency is idle for 10 days
func (f *Fulfillment) AddCandidateSubmission10DaysIdleActivity(supplierID string, clientID string, clientName string, auctionID string, auctionNumber string, agencyTenantID string, employerTenantID string) (stream.Activity, error) {
	helper := NewHelper(f.telemetryClient)

	applicationStatusID, applicationStatus := f.getAgencyStatusByAuctionID(auctionID, agencyTenantID)

	if applicationStatusID != approvedApplicationStatusID && applicationStatus != approvedApplicationStatus {
		return stream.Activity{}, errors.New("AddCandidateSubmission10DaysIdleActivity: agency is not allowed!")
	}

	clientData, _ := helper.createJSONMarshalWithUUID(clientID, clientName, "client", uuid.New().String())
	auctionData, _ := helper.createJSONMarshal(auctionID, auctionNumber, "auction")

	assignmentRepository := NewAssignmentRepository(f.auctionDB, f.telemetryClient)
	assignments, _ := assignmentRepository.GetAgencyAssignmentsByAuctionID(auctionID, agencyTenantID)

	if assignmentRepository.IsApprovedAuctionStatus(auctionID) {
		for _, assignment := range assignments {
			userID := strings.ToLower(assignment.UserID)

			verb := "submit"
			target := fmt.Sprintf("auction:%s", auctionID)
			content := fmt.Sprintf("Competitive Auction#%s, for %s, is waiting for a candidate submission. Are you still working on this Auction? Consider sending the employer an update on your progress. If you are unable to submit candidates soon, you may be removed from this Auction by the employer.", string(auctionData), string(clientData))
			category := "Fulfillment"
			subcategory := map[string]string{
				"type":   "Candidate",
				"status": "Required",
			}

			f.sendNotificationToAgency(userID, supplierID, clientID, verb, target, content, category, subcategory, agencyTenantID, employerTenantID)
		}
	}

	return stream.Activity{}, nil
}

// AddCandidateSubmission14DaysIdleActivity create a new activity when the agency is idle for 14 days
func (f *Fulfillment) AddCandidateSubmission14DaysIdleActivity(supplierID string, clientID string, clientName string, auctionID string, auctionNumber string, agencyTenantID string, employerTenantID string) (stream.Activity, error) {
	helper := NewHelper(f.telemetryClient)

	applicationStatusID, applicationStatus := f.getAgencyStatusByAuctionID(auctionID, agencyTenantID)

	if applicationStatusID != approvedApplicationStatusID && applicationStatus != approvedApplicationStatus {
		return stream.Activity{}, errors.New("AddCandidateSubmission14DaysIdleActivity: agency is not allowed!")
	}

	clientData, _ := helper.createJSONMarshalWithUUID(clientID, clientName, "client", uuid.New().String())
	auctionData, _ := helper.createJSONMarshal(auctionID, auctionNumber, "auction")

	assignmentRepository := NewAssignmentRepository(f.auctionDB, f.telemetryClient)
	assignments, _ := assignmentRepository.GetAgencyAssignmentsByAuctionID(auctionID, agencyTenantID)

	if assignmentRepository.IsApprovedAuctionStatus(auctionID) {
		for _, assignment := range assignments {
			userID := strings.ToLower(assignment.UserID)

			verb := "submit"
			target := fmt.Sprintf("auction:%s", auctionID)
			content := fmt.Sprintf("Competitve Auction#%s, for %s, is still waiting for a candidate submission. Please note, due to no candidate submission for this Auction, your approval is under review for revocation.", string(auctionData), string(clientData))
			category := "Fulfillment"
			subcategory := map[string]string{
				"type":   "Candidate",
				"status": "Required",
			}

			f.sendNotificationToAgency(userID, supplierID, clientID, verb, target, content, category, subcategory, agencyTenantID, employerTenantID)
		}
	}

	return stream.Activity{}, nil
}

// AddNewCandidateSuccessFeeActivity create a new activity when the agency submitted candidate
func (f *Fulfillment) AddNewCandidateSuccessFeeActivity(supplierID string, supplierName string, clientID string, clientName string, candidateID string, candidateName string, talentRequestID string, talentRequestNumber string, successFeeID string, successFeeNumber string, jobTitle string) (stream.Activity, error) {
	helper := NewHelper(f.telemetryClient)

	supplierData, _ := helper.createJSONMarshal(supplierID, supplierName, "supplier")
	candidateData, _ := helper.createJSONMarshal(candidateID, candidateName, "candidate")
	talentRequestData, _ := helper.createJSONMarshal(talentRequestID, fmt.Sprintf("%06s", talentRequestNumber), "talentRequest")
	successFeeData, _ := helper.createJSONMarshal(successFeeID, successFeeNumber, "successFee")

	assignmentRepository, settingsRepository := f.createAssignmentAndSettingsRepository()
	assignments, _ := assignmentRepository.GetEmployerAssignmentsBySuccessFeeID(successFeeID)

	if assignmentRepository.IsApprovedSuccessFeeStatus(successFeeID) {
		for _, assignment := range assignments {
			userID := strings.ToLower(assignment.UserID)
			settings, _ := settingsRepository.GetSettingsByClientID(userID)

			if settings.Fulfillment {
				verb := "submit"
				target := fmt.Sprintf("successfee:%s", successFeeID)
				content := fmt.Sprintf("%s submitted NEW candidate %s, for TR #%s: %s", string(supplierData), string(candidateData), string(talentRequestData), jobTitle)
				category := "Fulfillment"
				subcategory := map[string]string{
					"type":   "Candidate",
					"status": "New",
					"data":   fmt.Sprintf("%s", successFeeData),
				}

				f.sendNotificationtoEmployer(userID, supplierID, candidateID, verb, target, content, category, subcategory)
			}
		}

		formData := url.Values{
			"supplierID":          {supplierID},
			"supplierName":        {supplierName},
			"clientID":            {clientID},
			"clientName":          {clientName},
			"candidateID":         {candidateID},
			"candidateName":       {candidateName},
			"talentRequestID":     {talentRequestID},
			"talentRequestNumber": {talentRequestNumber},
			"successFeeID":        {successFeeID},
			"successFeeNumber":    {successFeeNumber},
			"jobTitle":            {jobTitle},
		}

		helper.post(candidatePendingSuccessFee, formData)
	}

	return stream.Activity{}, nil
}

// AddUpdateCandidateSuccessFeeActivity create an activity when agency updated candidate
func (f *Fulfillment) AddUpdateCandidateSuccessFeeActivity(supplierID string, supplierName string, clientID string, clientName string, candidateID string, candidateName string, talentRequestID string, talentRequestNumber string, successFeeID string, successFeeNumber string, jobTitle string, candidateStatus string) (stream.Activity, error) {
	helper := NewHelper(f.telemetryClient)

	supplierData, _ := helper.createJSONMarshal(supplierID, supplierName, "supplier")
	candidateData, _ := helper.createJSONMarshal(candidateID, candidateName, "candidate")
	talentRequestData, _ := helper.createJSONMarshal(talentRequestID, fmt.Sprintf("%06s", talentRequestNumber), "talentRequest")
	successFeeData, _ := helper.createJSONMarshal(successFeeID, successFeeNumber, "successFee")

	assignmentRepository, settingsRepository := f.createAssignmentAndSettingsRepository()
	assignments, _ := assignmentRepository.GetEmployerAssignmentsBySuccessFeeID(successFeeID)

	if assignmentRepository.IsApprovedSuccessFeeStatus(successFeeID) {
		for _, assignment := range assignments {
			userID := strings.ToLower(assignment.UserID)
			settings, _ := settingsRepository.GetSettingsByClientID(userID)

			if settings.Fulfillment {
				verb := "update"
				target := fmt.Sprintf("successfee:%s", successFeeID)
				content := fmt.Sprintf("%s updated candidate %s, for TR #%s: %s", string(supplierData), string(candidateData), string(talentRequestData), jobTitle)
				category := "Fulfillment"
				subcategory := map[string]string{
					"type":   "Candidate",
					"status": strings.Title(candidateStatus),
					"data":   fmt.Sprintf("%s", successFeeData),
				}

				f.sendNotificationtoEmployer(userID, supplierID, candidateID, verb, target, content, category, subcategory)
			}
		}
	}

	return stream.Activity{}, nil
}

// AddCandidateSubmission3DaysIdleSuccessFeeActivity create a new activity when the agency is idle for 3 days
func (f *Fulfillment) AddCandidateSubmission3DaysIdleSuccessFeeActivity(supplierID string, clientID string, clientName string, successFeeID string, successFeeNumber string, agencyTenantID string, employerTenantID string) (stream.Activity, error) {
	helper := NewHelper(f.telemetryClient)

	applicationStatusID, applicationStatus := f.getAgencyStatusBySuccessFeeID(successFeeID, agencyTenantID)

	if applicationStatusID != approvedApplicationStatusID && applicationStatus != approvedApplicationStatus {
		return stream.Activity{}, errors.New("AddCandidateSubmission3DaysIdleSuccessFeeActivity: agency is not allowed!")
	}

	clientData1, _ := helper.createJSONMarshalWithUUID(clientID, clientName, "client", uuid.New().String())
	clientData2, _ := helper.createJSONMarshalWithUUID(clientID, clientName, "client", uuid.New().String())
	successFeeData, _ := helper.createJSONMarshal(successFeeID, successFeeNumber, "successFee")

	assignmentRepository := NewAssignmentRepository(f.auctionDB, f.telemetryClient)
	assignments, _ := assignmentRepository.GetAgencyAssignmentsBySuccessFeeID(successFeeID, agencyTenantID)

	if assignmentRepository.IsApprovedSuccessFeeStatus(successFeeID) {
		for _, assignment := range assignments {
			userID := strings.ToLower(assignment.UserID)

			verb := "submit"
			target := fmt.Sprintf("successfee:%s", successFeeID)
			content := fmt.Sprintf("Engagement #%s, for %s, is waiting for a candidate submission. If you have interviewed a candidate or have a Candidate ready, please remember to submit so %s can start their review!", string(successFeeData), string(clientData1), string(clientData2))
			category := "Fulfillment"
			subcategory := map[string]string{
				"type":   "Candidate",
				"status": "Required",
			}

			f.sendNotificationToAgency(userID, supplierID, clientID, verb, target, content, category, subcategory, agencyTenantID, employerTenantID)
		}
	}

	return stream.Activity{}, nil
}

// AddCandidateSubmission10DaysIdleSuccessFeeActivity create a new activity when the agency is idle for 10 days
func (f *Fulfillment) AddCandidateSubmission10DaysIdleSuccessFeeActivity(supplierID string, clientID string, clientName string, successFeeID string, successFeeNumber string, agencyTenantID string, employerTenantID string) (stream.Activity, error) {
	helper := NewHelper(f.telemetryClient)

	applicationStatusID, applicationStatus := f.getAgencyStatusBySuccessFeeID(successFeeID, agencyTenantID)

	if applicationStatusID != approvedApplicationStatusID && applicationStatus != approvedApplicationStatus {
		return stream.Activity{}, errors.New("AddCandidateSubmission10DaysIdleSuccessFeeActivity: agency is not allowed!")
	}

	clientData, _ := helper.createJSONMarshalWithUUID(clientID, clientName, "client", uuid.New().String())
	successFeeData, _ := helper.createJSONMarshal(successFeeID, successFeeNumber, "successFee")

	assignmentRepository := NewAssignmentRepository(f.auctionDB, f.telemetryClient)
	assignments, _ := assignmentRepository.GetAgencyAssignmentsBySuccessFeeID(successFeeID, agencyTenantID)

	if assignmentRepository.IsApprovedSuccessFeeStatus(successFeeID) {
		for _, assignment := range assignments {
			userID := strings.ToLower(assignment.UserID)

			verb := "submit"
			target := fmt.Sprintf("successfee:%s", successFeeID)
			content := fmt.Sprintf("Engagement #%s, for %s, is waiting for a candidate submission. Are you still working on this Engagement? Consider sending the employer an update on your progress. If you are unable to submit candidates soon, you may be removed from this Auction by the employer.", string(successFeeData), string(clientData))
			category := "Fulfillment"
			subcategory := map[string]string{
				"type":   "Candidate",
				"status": "Required",
			}

			f.sendNotificationToAgency(userID, supplierID, clientID, verb, target, content, category, subcategory, agencyTenantID, employerTenantID)
		}
	}

	return stream.Activity{}, nil
}

// AddCandidateSubmission14DaysIdleSuccessFeeActivity create a new activity when the agency is idle for 14 days
func (f *Fulfillment) AddCandidateSubmission14DaysIdleSuccessFeeActivity(supplierID string, clientID string, clientName string, successFeeID string, successFeeNumber string, agencyTenantID string, employerTenantID string) (stream.Activity, error) {
	helper := NewHelper(f.telemetryClient)
	applicationStatusID, applicationStatus := f.getAgencyStatusBySuccessFeeID(successFeeID, agencyTenantID)

	if applicationStatusID != approvedApplicationStatusID && applicationStatus != approvedApplicationStatus {
		return stream.Activity{}, errors.New("AddCandidateSubmission14DaysIdleSuccessFeeActivity: agency is not allowed!")
	}

	clientData, _ := helper.createJSONMarshalWithUUID(clientID, clientName, "client", uuid.New().String())
	successFeeData, _ := helper.createJSONMarshal(successFeeID, successFeeNumber, "successFee")

	assignmentRepository := NewAssignmentRepository(f.auctionDB, f.telemetryClient)
	assignments, _ := assignmentRepository.GetAgencyAssignmentsBySuccessFeeID(successFeeID, agencyTenantID)

	if assignmentRepository.IsApprovedSuccessFeeStatus(successFeeID) {
		for _, assignment := range assignments {
			userID := strings.ToLower(assignment.UserID)

			verb := "submit"
			target := fmt.Sprintf("successfee:%s", successFeeID)
			content := fmt.Sprintf("Engagement #%s, for %s, is still waiting for a candidate submission. Please note, due to no candidate submission for this Engagement, your approval is under review for revocation.", string(successFeeData), string(clientData))
			category := "Fulfillment"
			subcategory := map[string]string{
				"type":   "Candidate",
				"status": "Required",
			}

			f.sendNotificationToAgency(userID, supplierID, clientID, verb, target, content, category, subcategory, agencyTenantID, employerTenantID)
		}
	}

	return stream.Activity{}, nil
}

func (f *Fulfillment) getAgencyStatusByAuctionID(auctionID string, agencyTenantID string) (int, string) {
	query := "select ApplicationStatusId, ApplicationStatus from Applications where AuctionId=? and SupplierId=? and IsDeleted=0"

	err := f.auctionDB.Ping()

	if err != nil {
		f.telemetryClient.TrackException(err)
		panic(err)
	}

	rows, err := f.auctionDB.Query(query, auctionID, agencyTenantID)

	var applicationStatusID int
	var applicationStatus string

	if rows.Next() {
		rows.Scan(&applicationStatusID, applicationStatus)
	}

	return applicationStatusID, applicationStatus
}

func (f *Fulfillment) getAgencyStatusBySuccessFeeID(successFeeID string, agencyTenantID string) (int, string) {
	query := "select SuccessFeeApplicationStatusId, SuccessFeeApplicationStatus from SuccessFeeApplications where SuccessFeeId=? and SupplierId=? and IsDeleted=0"

	err := f.auctionDB.Ping()

	if err != nil {
		f.telemetryClient.TrackException(err)
		panic(err)
	}

	rows, err := f.auctionDB.Query(query, successFeeID, agencyTenantID)

	var successFeeApplicationStatusID int
	var successFeeApplicationStatus string

	if rows.Next() {
		rows.Scan(&successFeeApplicationStatusID, &successFeeApplicationStatus)
	}

	return successFeeApplicationStatusID, successFeeApplicationStatus
}

func (f *Fulfillment) createAssignmentAndSettingsRepository() (*Assignment, *Settings) {
	assignmentRepository := NewAssignmentRepository(f.auctionDB, f.telemetryClient)
	settingsRepository := NewSettingsRepository(f.employerDB, nil, f.telemetryClient)

	return assignmentRepository, settingsRepository
}

func (f *Fulfillment) sendNotificationToAgency(userID string, supplierID string, clientID string, verb string, target string, content string, category string, subcategory map[string]string, agencyTenantID string, employerTenantID string) {
	agencyFeed := f.client.FlatFeed("agency", supplierID)
	employerFeed := f.client.FlatFeed("candidate", clientID)

	agencyNotificationFeed := f.client.NotificationFeed("agencynotification", userID)
	_, err := agencyNotificationFeed.AddActivity(stream.Activity{
		Actor:     agencyFeed.ID(),
		Verb:      verb,
		Object:    employerFeed.ID(),
		Target:    target,
		ForeignID: uuid.New().String(),
		Time:      stream.Time{time.Now().UTC()},
		Extra: map[string]interface{}{
			"employer":         fmt.Sprintf("employer:%s", clientID),
			"agency":           fmt.Sprintf("agency:%s", userID),
			"content":          content,
			"category":         category,
			"subcategory":      subcategory,
			"agencyTenantID":   agencyTenantID,
			"employerTenantID": employerTenantID,
		},
	})

	if err != nil {
		f.telemetryClient.TrackException(err)
		panic(err)
	}
}

func (f *Fulfillment) sendNotificationtoEmployer(userID string, supplierID string, candidateID string, verb string, target string, content string, category string, subcategory map[string]string) {
	employerNotificationFeed := f.client.NotificationFeed("employernotification", userID)
	agencyFeed := f.client.FlatFeed("agency", supplierID)
	candidateFeed := f.client.FlatFeed("candidate", candidateID)

	_, err := employerNotificationFeed.AddActivity(stream.Activity{
		Actor:     agencyFeed.ID(),
		Verb:      verb,
		Object:    candidateFeed.ID(),
		Target:    target,
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
		f.telemetryClient.TrackException(err)
		panic(err)
	}
}
