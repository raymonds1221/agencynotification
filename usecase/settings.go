package usecase

import (
	"github.com/Ubidy/Ubidy_AgencyNotificationAPI/domain/settings"
	"github.com/Ubidy/Ubidy_AgencyNotificationAPI/usecase/repository"
)

// SettingsInteractor usecase implementation for settings
type SettingsInteractor struct {
	repository repository.Settings
}

// NewSettingsInteractor create new instance of settings interactor
func NewSettingsInteractor(r repository.Settings) *SettingsInteractor {
	return &SettingsInteractor{
		repository: r,
	}
}

// GetSettingsByClientID retrieve the settings by clientID
func (si *SettingsInteractor) GetSettingsByClientID(clientID string) (settings.Settings, error) {
	s, err := si.repository.GetSettingsByClientID(clientID)

	if err != nil {
		return settings.Settings{}, err
	}

	return s, nil
}

// GetSettingsBySupplierID retrieve the settings by supplierID
func (si *SettingsInteractor) GetSettingsBySupplierID(supplierID string) (settings.Settings, error) {
	s, err := si.repository.GetSettingsBySupplierID(supplierID)

	if err != nil {
		return settings.Settings{}, err
	}

	return s, nil
}

// CreateOrUpdateSettings create or update a settings by tenantID
func (si *SettingsInteractor) CreateOrUpdateSettings(userID string, auctionsScheduling bool, prequalification bool, applications bool, clarifications bool, bidding bool, awarding bool, fulfillment bool, payments bool, ubidy bool, messages bool, users bool, activeSettings string) bool {
	result := si.repository.CreateOrUpdateSettings(userID, auctionsScheduling, prequalification, applications, clarifications, bidding, awarding, fulfillment, payments, ubidy, messages, users, activeSettings)

	return result
}
