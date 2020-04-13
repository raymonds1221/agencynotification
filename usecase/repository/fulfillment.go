package repository

import stream "gopkg.in/GetStream/stream-go2.v1"

// Fulfillment interface for fulfillment repository
type Fulfillment interface {
	AddNewCandidateActivity(supplierID string, supplierName string, clientID string, clientName string, candidateID string, candidateName string, talentRequestID string, talentRequestNumber string, auctionID string, auctionNumber string, jobTitle string) (stream.Activity, error)
	AddUpdateCandidateActivity(supplierID string, supplierName string, clientID string, clientName string, candidateID string, candidateName string, talentRequestID string, talentRequestNumber string, auctionID string, auctionNumber string, jobTitle string) (stream.Activity, error)
	AddCandidateSubmission3DaysIdleActivity(supplierID string, clientID string, clientName string, auctionID string, auctionNumber string, agencyTenantID string, employerTenantID string) (stream.Activity, error)
	AddCandidateSubmission10DaysIdleActivity(supplierID string, clientID string, clientName string, auctionID string, auctionNumber string, agencyTenantID string, employerTenantID string) (stream.Activity, error)
	AddCandidateSubmission14DaysIdleActivity(supplierID string, clientID string, clientName string, auctionID string, auctionNumber string, agencyTenantID string, employerTenantID string) (stream.Activity, error)

	AddNewCandidateSuccessFeeActivity(supplierID string, supplierName string, clientID string, clientName string, candidateID string, candidateName string, talentRequestID string, talentRequestNumber string, successFeeID string, successFeeNumber string, jobTitle string) (stream.Activity, error)
	AddUpdateCandidateSuccessFeeActivity(supplierID string, supplierName string, clientID string, clientName string, candidateID string, candidateName string, talentRequestID string, talentRequestNumber string, successFeeID string, successFeeNumber string, jobTitle string, candidateStatus string) (stream.Activity, error)
	AddCandidateSubmission3DaysIdleSuccessFeeActivity(supplierID string, clientID string, clientName string, successFeeID string, successFeeNumber string, agencyTenantID string, employerTenantID string) (stream.Activity, error)
	AddCandidateSubmission10DaysIdleSuccessFeeActivity(supplierID string, clientID string, clientName string, successFeeID string, successFeeNumber string, agencyTenantID string, employerTenantID string) (stream.Activity, error)
	AddCandidateSubmission14DaysIdleSuccessFeeActivity(supplierID string, clientID string, clientName string, successFeeID string, successFeeNumber string, agencyTenantID string, employerTenantID string) (stream.Activity, error)
}
