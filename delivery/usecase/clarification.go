package usecase

import stream "gopkg.in/GetStream/stream-go2.v1"

// ClarificationInteractor interface for clarification usecase
type ClarificationInteractor interface {
	AddPostTopicActivity(supplierID string, supplierName string, clientID string, clientName string, auctionID string, auctionNumber string) (stream.Activity, error)
	AddPostQuestionActivity(supplierID string, supplierName string, clientID string, clientName string, auctionID string, auctionNumber string) (stream.Activity, error)

	AddPostTopicSuccessFeeActivity(supplierID string, supplierName string, clientID string, clientName string, successFeeID string, successFeeNumber string) (stream.Activity, error)
	AddPostQuestionSuccessFeeActivity(supplierID string, supplierName string, clientID string, clientName string, successFeeID string, successFeeNumber string) (stream.Activity, error)
}
