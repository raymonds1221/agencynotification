package repository

import "github.com/Ubidy/Ubidy_AgencyNotificationAPI/domain/settings"

// Settings repository for getting settings
type Settings interface {
	GetSettingsByClientID(clientID string) (settings.Settings, error)
	GetSettingsBySupplierID(supplierID string) (settings.Settings, error)
	CreateOrUpdateSettings(userID string, auctionsScheduling bool, prequalification bool, applications bool, clarifications bool, bidding bool, awarding bool, fulfillment bool, payments bool, ubidy bool, messages bool, users bool, activeSettings string) bool
}
