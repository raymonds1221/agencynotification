package usecase

import (
	"github.com/Ubidy/Ubidy_AgencyNotificationAPI/usecase/repository"
	stream "gopkg.in/GetStream/stream-go2.v1"
)

// ApplicationInteractor implementation of application usecase
type ApplicationInteractor struct {
	repository repository.Application
}

// NewApplicationInteractor create new instance of application interactor
func NewApplicationInteractor(r repository.Application) *ApplicationInteractor {
	return &ApplicationInteractor{
		repository: r,
	}
}

// AddSubmitApplicationActivity create an activity when agency submitted application
func (ai *ApplicationInteractor) AddSubmitApplicationActivity(supplierID string, supplierName string, clientID string, clientName string, auctionID string, auctionNumber string) (stream.Activity, error) {
	activity, err := ai.repository.AddSubmitApplicationActivity(supplierID, supplierName, clientID, clientName, auctionID, auctionNumber)

	if err != nil {
		return stream.Activity{}, err
	}

	return activity, nil
}

// AddWithdrawApplicationActivity create an activity when agency withdraw his/her application
func (ai *ApplicationInteractor) AddWithdrawApplicationActivity(supplierID string, supplierName string, clientID string, clientName string, auctionID string, auctionNumber string) (stream.Activity, error) {
	activity, err := ai.repository.AddWithdrawApplicationActivity(supplierID, supplierName, clientID, clientName, auctionID, auctionNumber)

	if err != nil {
		return stream.Activity{}, err
	}

	return activity, nil
}

// AddSubmitApplicationSuccessFeeActivity create an activity when agency submitted application
func (ai *ApplicationInteractor) AddSubmitApplicationSuccessFeeActivity(supplierID string, supplierName string, clientID string, clientName string, successFeeID string, successFeeNumber string) (stream.Activity, error) {
	activity, err := ai.repository.AddSubmitApplicationSuccessFeeActivity(supplierID, supplierName, clientID, clientName, successFeeID, successFeeNumber)

	if err != nil {
		return stream.Activity{}, err
	}

	return activity, nil
}

// AddWithdrawApplicationSuccessFeeActivity create an activity when agency withdraw his/her application
func (ai *ApplicationInteractor) AddWithdrawApplicationSuccessFeeActivity(supplierID string, supplierName string, clientID string, clientName string, successFeeID string, successFeeNumber string) (stream.Activity, error) {
	activity, err := ai.repository.AddWithdrawApplicationSuccessFeeActivity(supplierID, supplierName, clientID, clientName, successFeeID, successFeeNumber)

	if err != nil {
		return stream.Activity{}, err
	}

	return activity, nil
}
