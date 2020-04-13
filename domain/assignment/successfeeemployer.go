package assignment

// SuccessFeeEmployer model for success fee employer assignment
type SuccessFeeEmployer struct {
	ID           string `json:"id"`
	SuccessFeeID string `json:"successFeeID"`
	ClientID     string `json:"clientID"`
	UserID       string `json:"userID"`
}

// NewSuccessFeeEmployer create new instance of success fee employer assignment
func NewSuccessFeeEmployer(id string, successFeeID string, clientID string, userID string) *SuccessFeeEmployer {
	return &SuccessFeeEmployer{
		ID:           id,
		SuccessFeeID: successFeeID,
		ClientID:     clientID,
		UserID:       userID,
	}
}
