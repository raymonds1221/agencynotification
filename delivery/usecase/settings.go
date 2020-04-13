package usecase

import "github.com/Ubidy/Ubidy_AgencyNotificationAPI/domain/settings"

// SettingsInteractor usecase declaration for settings
type SettingsInteractor interface {
	GetSettingsByClientID(clientID string) (settings.Settings, error)
	GetSettingsBySupplierID(supplierID string) (settings.Settings, error)
	CreateOrUpdateSettings(userID string, auctionsScheduling bool, prequalification bool, applications bool, clarifications bool, bidding bool, awarding bool, fulfillment bool, payments bool, ubidy bool, messages bool, users bool, activeSettings string) bool
}
