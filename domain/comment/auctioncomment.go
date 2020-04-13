package comment

// AuctionComment model for auction comment
type AuctionComment struct {
	Comment
	AuctionID     string
	AuctionNumber string
}
