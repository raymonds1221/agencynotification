package repository

import (
	stream "gopkg.in/GetStream/stream-go2.v1"
)

// Application interface for application repository
type Application interface {
	AddSubmitApplicationActivity(supplierID string, supplierName string, clientID string, clientName string, auctionID string, auctionNumber string) (stream.Activity, error)
	AddWithdrawApplicationActivity(supplierID string, supplierName string, clientID string, clientName string, auctionID string, auctionNumber string) (stream.Activity, error)

	AddSubmitApplicationSuccessFeeActivity(supplierID string, supplierName string, clientID string, clientName string, successFeeID string, successFeeNumber string) (stream.Activity, error)
	AddWithdrawApplicationSuccessFeeActivity(supplierID string, supplierName string, clientID string, clientName string, successFeeID string, successFeeNumber string) (stream.Activity, error)
}
