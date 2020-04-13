package repository

import (
	"database/sql"

	"github.com/Ubidy/Ubidy_AgencyNotificationAPI/domain/auction"
)

// Auction implementation of Auction repository
type Auction struct {
	db *sql.DB
}

// NewAuctionRepository create new instance of auction repository
func NewAuctionRepository(auctionDB *sql.DB) *Auction {
	return &Auction{
		db: auctionDB,
	}
}

func (a *Auction) getAuctionByID(auctionID string) *auction.Auction {
	query := "SELECT ClientID, ClientName, AuctionNumber, AuctionStatusId, AuctionStatus FROM Auctions WHERE AuctionId=?"

	a.db.Ping()

	rows, err := a.db.Query(query, auctionID)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	if n := rows.Next(); !n {
		return nil
	}

	var auctionNumber, auctionStatusID int64
	var clientID, clientName, auctionStatus string

	err = rows.Scan(&clientID, &clientName, &auctionNumber, &auctionStatusID, &auctionStatus)

	if err != nil {
		panic(err)
	}

	return auction.NewAuction(auctionID, clientID, clientName, auctionNumber, auctionStatusID, auctionStatus)
}

func (a *Auction) getSuccessFeeByID(successFeeID string) *auction.SuccessFee {
	query := "SELECT ClientID, ClientName, SuccessFeeNumber, SuccessFeeStatusId, SuccessFeeStatus FROM SuccessFees where SuccessFeeId=?"

	a.db.Ping()

	rows, err := a.db.Query(query, successFeeID)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	if n := rows.Next(); !n {
		return nil
	}

	var successFeeNumber, successFeeStatusID int64
	var clientID, clientName, successFeeStatus string

	err = rows.Scan(&clientID, &clientName, &successFeeNumber, &successFeeStatusID, &successFeeStatus)

	if err != nil {
		panic(err)
	}

	return auction.NewSuccessFee(successFeeID, clientID, clientName, successFeeNumber, successFeeStatusID, successFeeStatus)
}
