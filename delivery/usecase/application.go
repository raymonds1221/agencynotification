package usecase

import (
	stream "gopkg.in/GetStream/stream-go2.v1"
)

// ApplicationInteractor interface for application usecase
type ApplicationInteractor interface {
	AddSubmitApplicationActivity(supplierID string, supplierName string, clientID string, clientName string, auctionID string, auctionNumber string) (stream.Activity, error)
	AddWithdrawApplicationActivity(supplierID string, supplierName string, clientID string, clientName string, auctionID string, auctionNumber string) (stream.Activity, error)

	AddSubmitApplicationSuccessFeeActivity(supplierID string, supplierName string, clientID string, clientName string, successFeeID string, successFeeNumber string) (stream.Activity, error)
	AddWithdrawApplicationSuccessFeeActivity(supplierID string, supplierName string, clientID string, clientName string, successFeeID string, successFeeNumber string) (stream.Activity, error)
}
