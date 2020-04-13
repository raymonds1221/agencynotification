package usecase

import (
	"github.com/Ubidy/Ubidy_AgencyNotificationAPI/usecase/repository"
	stream "gopkg.in/GetStream/stream-go2.v1"
)

// FulfillmentInteractor implementation of fulfillment interactor
type FulfillmentInteractor struct {
	repository repository.Fulfillment
}

// NewFulfillmentInteractor create new instance of fulfillment interactor
func NewFulfillmentInteractor(r repository.Fulfillment) *FulfillmentInteractor {
	return &FulfillmentInteractor{
		repository: r,
	}
}

// AddNewCandidateActivity create new activity when agency submitted new candidate
func (fi *FulfillmentInteractor) AddNewCandidateActivity(supplierID string, supplierName string, clientID string, clientName string, candidateID string, candidateName string, talentRequestID string, talentRequestNumber string, auctionID string, auctionNumber string, jobTitle string) (stream.Activity, error) {
	activity, err := fi.repository.AddNewCandidateActivity(supplierID, supplierName, clientID, clientName, candidateID, candidateName, talentRequestID, talentRequestNumber, auctionID, auctionNumber, jobTitle)

	if err != nil {
		return stream.Activity{}, err
	}

	return activity, nil
}

// AddUpdateCandidateActivity create new activity when agency updated candidate
func (fi *FulfillmentInteractor) AddUpdateCandidateActivity(supplierID string, supplierName string, clientID string, clientName string, candidateID string, candidateName string, talentRequestID string, talentRequestNumber string, auctionID string, auctionNumber string, jobTitle string) (stream.Activity, error) {
	activity, err := fi.repository.AddUpdateCandidateActivity(supplierID, supplierName, clientID, clientName, candidateID, candidateName, talentRequestID, talentRequestNumber, auctionID, auctionNumber, jobTitle)

	if err != nil {
		return stream.Activity{}, err
	}

	return activity, nil
}

// AddCandidateSubmission3DaysIdleActivity create new activity when the agency is idle for 3 days
func (fi *FulfillmentInteractor) AddCandidateSubmission3DaysIdleActivity(supplierID string, clientID string, clientName string, auctionID string, auctionNumber string, agencyTenantID string, employerTenantID string) (stream.Activity, error) {
	activity, err := fi.repository.AddCandidateSubmission3DaysIdleActivity(supplierID, clientID, clientName, auctionID, auctionNumber, agencyTenantID, employerTenantID)

	if err != nil {
		return stream.Activity{}, err
	}

	return activity, nil
}

// AddCandidateSubmission10DaysIdleActivity create new activity when the agency is idle for 10 days
func (fi *FulfillmentInteractor) AddCandidateSubmission10DaysIdleActivity(supplierID string, clientID string, clientName string, auctionID string, auctionNumber string, agencyTenantID string, employerTenantID string) (stream.Activity, error) {
	activity, err := fi.repository.AddCandidateSubmission10DaysIdleActivity(supplierID, clientID, clientName, auctionID, auctionNumber, agencyTenantID, employerTenantID)

	if err != nil {
		return stream.Activity{}, err
	}

	return activity, nil
}

// AddCandidateSubmission14DaysIdleActivity create new activity when the agency is idle for 14 days
func (fi *FulfillmentInteractor) AddCandidateSubmission14DaysIdleActivity(supplierID string, clientID string, clientName string, auctionID string, auctionNumber string, agencyTenantID string, employerTenantID string) (stream.Activity, error) {
	activity, err := fi.repository.AddCandidateSubmission14DaysIdleActivity(supplierID, clientID, clientName, auctionID, auctionNumber, agencyTenantID, employerTenantID)

	if err != nil {
		return stream.Activity{}, err
	}

	return activity, nil
}

// AddNewCandidateSuccessFeeActivity create new activity when agency submitted new candidate
func (fi *FulfillmentInteractor) AddNewCandidateSuccessFeeActivity(supplierID string, supplierName string, clientID string, clientName string, candidateID string, candidateName string, talentRequestID string, talentRequestNumber string, successFeeID string, successFeeNumber string, jobTitle string) (stream.Activity, error) {
	activity, err := fi.repository.AddNewCandidateSuccessFeeActivity(supplierID, supplierName, clientID, clientName, candidateID, candidateName, talentRequestID, talentRequestNumber, successFeeID, successFeeNumber, jobTitle)

	if err != nil {
		return stream.Activity{}, err
	}

	return activity, nil
}

// AddUpdateCandidateSuccessFeeActivity creat new activity when agency updated candidate
func (fi *FulfillmentInteractor) AddUpdateCandidateSuccessFeeActivity(supplierID string, supplierName string, clientID string, clientName string, candidateID string, candidateName string, talentRequestID string, talentRequestNumber string, successFeeID string, successFeeNumber string, jobTitle string, candidateStatus string) (stream.Activity, error) {
	activity, err := fi.repository.AddUpdateCandidateSuccessFeeActivity(supplierID, supplierName, clientID, clientName, candidateID, candidateName, talentRequestID, talentRequestNumber, successFeeID, successFeeNumber, jobTitle, candidateStatus)

	if err != nil {
		return stream.Activity{}, err
	}

	return activity, nil
}

// AddCandidateSubmission3DaysIdleSuccessFeeActivity create new activity when the agency is idle for 3 days
func (fi *FulfillmentInteractor) AddCandidateSubmission3DaysIdleSuccessFeeActivity(supplierID string, clientID string, clientName string, successFeeID string, successFeeNumber string, agencyTenantID string, employerTenantID string) (stream.Activity, error) {
	activity, err := fi.repository.AddCandidateSubmission3DaysIdleSuccessFeeActivity(supplierID, clientID, clientName, successFeeID, successFeeNumber, agencyTenantID, employerTenantID)

	if err != nil {
		return stream.Activity{}, err
	}

	return activity, nil
}

// AddCandidateSubmission10DaysIdleSuccessFeeActivity create new activity when the agency is idle for 10 days
func (fi *FulfillmentInteractor) AddCandidateSubmission10DaysIdleSuccessFeeActivity(supplierID string, clientID string, clientName string, successFeeID string, successFeeNumber string, agencyTenantID string, employerTenantID string) (stream.Activity, error) {
	activity, err := fi.repository.AddCandidateSubmission10DaysIdleSuccessFeeActivity(supplierID, clientID, clientName, successFeeID, successFeeNumber, agencyTenantID, employerTenantID)

	if err != nil {
		return stream.Activity{}, err
	}

	return activity, nil
}

// AddCandidateSubmission14DaysIdleSuccessFeeActivity create new activity when the agency is idle for 14 days
func (fi *FulfillmentInteractor) AddCandidateSubmission14DaysIdleSuccessFeeActivity(supplierID string, clientID string, clientName string, successFeeID string, successFeeNumber string, agencyTenantID string, employerTenantID string) (stream.Activity, error) {
	activity, err := fi.repository.AddCandidateSubmission14DaysIdleSuccessFeeActivity(supplierID, clientID, clientName, successFeeID, successFeeNumber, agencyTenantID, employerTenantID)

	if err != nil {
		return stream.Activity{}, err
	}

	return activity, nil
}
