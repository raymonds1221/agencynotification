package assignment

// AuctionAgency model for auction agency assignment
type AuctionAgency struct {
	ID         string `json:"id"`
	AuctionID  string `json:"auctionID"`
	SupplierID string `json:"supplierID"`
	UserID     string `json:"userID"`
}

// NewAuctionAgency create new instance of auction agency assignment
func NewAuctionAgency(id string, auctionID string, supplierID string, userID string) *AuctionAgency {
	return &AuctionAgency{
		ID:         id,
		AuctionID:  auctionID,
		SupplierID: supplierID,
		UserID:     userID,
	}
}
