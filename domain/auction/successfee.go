package auction

const (
	// ApprovedSuccessFeeStatusID status id for approved success fee
	ApprovedSuccessFeeStatusID = 2
	// ApprovedSuccessFeeStatus status id for approved success fee
	ApprovedSuccessFeeStatus = "Approved"
)

// SuccessFee model for success fee
type SuccessFee struct {
	SuccessFeeID       string `json:"successFeeID"`
	ClientID           string `json:"clientID"`
	ClientName         string `json:"clientName"`
	SuccessFeeNumber   int64  `json:"successFeeNumber"`
	SuccessFeeStatusID int64  `json:"successFeeStatusID"`
	SuccessFeeStatus   string `json:"successFeeStatus"`
}

// NewSuccessFee create new instanc of success fee model
func NewSuccessFee(successFeeID string, clientID string, clientName string, successFeeNumber int64, successFeeStatusID int64, successFeeStatus string) *SuccessFee {
	return &SuccessFee{
		SuccessFeeID:       successFeeID,
		ClientID:           clientID,
		ClientName:         clientName,
		SuccessFeeNumber:   successFeeNumber,
		SuccessFeeStatusID: successFeeStatusID,
		SuccessFeeStatus:   successFeeStatus,
	}
}

// IsApprovedSuccessFee check if the success fee is approved
func (s *SuccessFee) IsApprovedSuccessFee() bool {
	isApprovedSuccessFeeStatus := s != nil &&
		s.SuccessFeeStatusID == ApprovedSuccessFeeStatusID &&
		s.SuccessFeeStatus == ApprovedSuccessFeeStatus

	if isApprovedSuccessFeeStatus {
		return true
	}
	return false
}
