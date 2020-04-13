package settings

// Item enum for different types of settings
type Item int

const (
	// AuctionsScheduling setting item for auction scheduling
	AuctionsScheduling Item = iota
	// Prequalification setting item for prequalification
	Prequalification
	// Applications setting item for applications
	Applications
	// Clarifications setting item for clarifications
	Clarifications
	// Bidding setting item for bidding
	Bidding
	// Awarding setting item for awarding
	Awarding
	// Fulfillment setting item for fulfillment
	Fulfillment
	// Payments setting item for payments
	Payments
	// Ubidy setting item for ubidy
	Ubidy
	// Messages setting item for messages
	Messages
	// Users setting item for users
	Users
)

func (s Item) String() string {
	switch s {
	case AuctionsScheduling:
		return "Auction Scheduling"
	case Prequalification:
		return "Prequalification"
	case Applications:
		return "Applications"
	case Clarifications:
		return "Clarifications"
	case Bidding:
		return "Bidding"
	case Awarding:
		return "Awarding"
	case Fulfillment:
		return "Fulfillment"
	case Payments:
		return "Payments"
	case Ubidy:
		return "Ubidy"
	case Messages:
		return "Messages"
	case Users:
		return "Users"
	default:
		return ""
	}
}

// Settings domain model for setting
type Settings struct {
	ActivityStreamSettingID string `json:"activityStreamSettingID"`
	AuctionsScheduling      bool   `json:"auctionsScheduling"`
	Prequalification        bool   `json:"prequalification"`
	Applications            bool   `json:"applications"`
	Clarifications          bool   `json:"clarifications"`
	Bidding                 bool   `json:"bidding"`
	Awarding                bool   `json:"awarding"`
	Fulfillment             bool   `json:"fulfillment"`
	Payments                bool   `json:"payments"`
	Ubidy                   bool   `json:"ubidy"`
	Messages                bool   `json:"messages"`
	Users                   bool   `json:"users"`
}

// NewSettings create new instance of Setting domain model
func NewSettings(
	activityStreamSettingID string,
	auctionsScheduling bool,
	prequalification bool,
	applications bool,
	clarifications bool,
	bidding bool,
	awarding bool,
	fulfillment bool,
	payments bool,
	ubidy bool,
	messages bool,
	users bool,
) Settings {
	return Settings{
		ActivityStreamSettingID: activityStreamSettingID,
		AuctionsScheduling:      auctionsScheduling,
		Prequalification:        prequalification,
		Applications:            applications,
		Clarifications:          clarifications,
		Bidding:                 bidding,
		Awarding:                awarding,
		Fulfillment:             fulfillment,
		Payments:                payments,
		Ubidy:                   ubidy,
		Messages:                messages,
		Users:                   users,
	}
}
