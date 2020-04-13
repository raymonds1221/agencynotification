package usecase

import (
	"github.com/Ubidy/Ubidy_AgencyNotificationAPI/usecase/repository"
	stream "gopkg.in/GetStream/stream-go2.v1"
)

// BiddingInteractor implementation of bidding usecase
type BiddingInteractor struct {
	repository repository.Bidding
}

// NewBiddingInteractor create new instance of bidding interactor
func NewBiddingInteractor(r repository.Bidding) *BiddingInteractor {
	return &BiddingInteractor{
		repository: r,
	}
}

// AddPlaceBidActivity create an activity when agency place a bid
func (bi *BiddingInteractor) AddPlaceBidActivity(supplierID string, clientID string, auctionID string) (stream.Activity, error) {
	activity, err := bi.repository.AddPlaceBidActivity(supplierID, clientID, auctionID)

	if err != nil {
		return stream.Activity{}, err
	}

	return activity, nil
}
