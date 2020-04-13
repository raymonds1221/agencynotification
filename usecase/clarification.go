package usecase

import (
	"github.com/Ubidy/Ubidy_AgencyNotificationAPI/usecase/repository"
	stream "gopkg.in/GetStream/stream-go2.v1"
)

// ClarificationInteractor implementation of clarification usecase
type ClarificationInteractor struct {
	repository repository.Clarification
}

// NewClarificationInteractor create new instance clarification interactor
func NewClarificationInteractor(r repository.Clarification) *ClarificationInteractor {
	return &ClarificationInteractor{
		repository: r,
	}
}

// AddPostTopicActivity create an activity when agency posted a new topic to clarification
func (ci *ClarificationInteractor) AddPostTopicActivity(supplierID string, supplierName string, clientID string, clientName string, auctionID string, auctionNumber string) (stream.Activity, error) {
	activity, err := ci.repository.AddPostTopicActivity(supplierID, supplierName, clientID, clientName, auctionID, auctionNumber)

	if err != nil {
		return stream.Activity{}, err
	}

	return activity, nil
}

// AddPostQuestionActivity create an activity when agency posted a question on a discussion
func (ci *ClarificationInteractor) AddPostQuestionActivity(supplierID string, supplierName string, clientID string, clientName string, auctionID string, auctionNumber string) (stream.Activity, error) {
	activity, err := ci.repository.AddPostQuestionActivity(supplierID, supplierName, clientID, clientName, auctionID, auctionNumber)

	if err != nil {
		return stream.Activity{}, err
	}

	return activity, nil
}

// AddPostTopicSuccessFeeActivity create an activity when agency posted a new topic to clarification
func (ci *ClarificationInteractor) AddPostTopicSuccessFeeActivity(supplierID string, supplierName string, clientID string, clientName string, successFeeID string, successFeeNumber string) (stream.Activity, error) {
	activity, err := ci.repository.AddPostTopicSuccessFeeActivity(supplierID, supplierName, clientID, clientName, successFeeID, successFeeNumber)

	if err != nil {
		return stream.Activity{}, err
	}

	return activity, nil
}

// AddPostQuestionSuccessFeeActivity create an activity when agency posted a question on a discussion
func (ci *ClarificationInteractor) AddPostQuestionSuccessFeeActivity(supplierID string, supplierName string, clientID string, clientName string, successFeeID string, successFeeNumber string) (stream.Activity, error) {
	activity, err := ci.repository.AddPostQuestionSuccessFeeActivity(supplierID, supplierName, clientID, clientName, successFeeID, successFeeNumber)

	if err != nil {
		return stream.Activity{}, err
	}

	return activity, nil
}
