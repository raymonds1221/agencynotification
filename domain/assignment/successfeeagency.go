package assignment

// SuccessFeeAgency model for success fee agency assignment
type SuccessFeeAgency struct {
	ID           string `json:"id"`
	SuccessFeeID string `json:"successFeeID"`
	SupplierID   string `json:"supplierID"`
	UserID       string `json:"userID"`
}

// NewSuccessFeeAgency create new instance of success fee agency assignment
func NewSuccessFeeAgency(id string, successFeeID string, supplierID string, userID string) *SuccessFeeAgency {
	return &SuccessFeeAgency{
		ID:           id,
		SuccessFeeID: successFeeID,
		SupplierID:   supplierID,
		UserID:       userID,
	}
}
