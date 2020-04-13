package auction

const (
	// ApprovedAuctionStatusID status id for approved auctions
	ApprovedAuctionStatusID = 2
	// ApprovedAuctionStatus status for approved auctions
	ApprovedAuctionStatus = "Approved"
)

// Auction model for auction
type Auction struct {
	AuctionID       string `json:"auctionID"`
	ClientID        string `json:"clientID"`
	ClientName      string `json:"clientName"`
	AuctionNumber   int64  `json:"auctionNumber"`
	AuctionStatusID int64  `json:"auctionStatusID"`
	AuctionStatus   string `json:"auctionStatus"`
}

// NewAuction create new instance of auction model
func NewAuction(auctionID string, clientID string, clientName string, auctionNumber int64, auctionStatusID int64, auctionStatus string) *Auction {
	return &Auction{
		AuctionID:       auctionID,
		ClientID:        clientID,
		ClientName:      clientName,
		AuctionNumber:   auctionNumber,
		AuctionStatusID: auctionStatusID,
		AuctionStatus:   auctionStatus,
	}
}

// IsApprovedAuction check if the auction is approved
func (a *Auction) IsApprovedAuction() bool {
	isApprovedAuctionStatus := a != nil &&
		a.AuctionStatusID == ApprovedAuctionStatusID &&
		a.AuctionStatus == ApprovedAuctionStatus

	if isApprovedAuctionStatus {
		return true
	}
	return false
}
