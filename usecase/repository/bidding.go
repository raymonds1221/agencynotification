package repository

import stream "gopkg.in/GetStream/stream-go2.v1"

// Bidding interface for bidding repository
type Bidding interface {
	AddPlaceBidActivity(supplierID string, clientID string, auctionID string) (stream.Activity, error)
}
