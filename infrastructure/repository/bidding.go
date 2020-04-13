package repository

import (
	"database/sql"
	"fmt"

	"github.com/Microsoft/ApplicationInsights-Go/appinsights"
	stream "gopkg.in/GetStream/stream-go2.v1"
)

// Bidding implementation of bidding repository
type Bidding struct {
	client          *stream.Client
	employerDB      *sql.DB
	auctionDB       *sql.DB
	telemetryClient appinsights.TelemetryClient
}

// NewBiddingRepository create new instance of bidding repository
func NewBiddingRepository(client *stream.Client, employerDB *sql.DB, auctionDB *sql.DB, telemetryClient appinsights.TelemetryClient) *Bidding {
	return &Bidding{
		client:          client,
		employerDB:      employerDB,
		auctionDB:       auctionDB,
		telemetryClient: telemetryClient,
	}
}

// AddPlaceBidActivity create an activity when agency place a bid
func (b *Bidding) AddPlaceBidActivity(supplierID string, clientID string, auctionID string) (stream.Activity, error) {
	agencyFeed := b.client.FlatFeed("agency", supplierID)

	resp, err := agencyFeed.AddActivity(stream.Activity{
		Actor:  agencyFeed.ID(),
		Verb:   "place",
		Target: fmt.Sprintf("auction:%s", auctionID),
		To:     []string{fmt.Sprintf("employer:%s", clientID)},
	})

	if err != nil {
		return stream.Activity{}, err
	}

	return resp.Activity, nil
}
