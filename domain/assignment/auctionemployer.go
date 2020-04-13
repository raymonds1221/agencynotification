package assignment

// AuctionEmployer model for auction employer assignment
type AuctionEmployer struct {
	ID        string `json:"id"`
	AuctionID string `json:"auctionID"`
	ClientID  string `json:"clientID"`
	UserID    string `json:"userID"`
}

// NewAuctionEmployer create new instance of auction employer assignment
func NewAuctionEmployer(id string, auctionID string, clientID string, userID string) *AuctionEmployer {
	return &AuctionEmployer{
		ID:        id,
		AuctionID: auctionID,
		ClientID:  clientID,
		UserID:    userID,
	}
}
