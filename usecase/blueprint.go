package usecase

import (
	"github.com/Ubidy/Ubidy_AgencyNotificationAPI/domain/blueprint"
	"github.com/Ubidy/Ubidy_AgencyNotificationAPI/usecase/repository"
)

// BlueprintInteractor implementation of Blueprint usecase
type BlueprintInteractor struct {
	repository repository.Blueprint
}

// NewBlueprintInteractor create new instance of blueprint usecase
func NewBlueprintInteractor(br repository.Blueprint) *BlueprintInteractor {
	return &BlueprintInteractor{
		repository: br,
	}
}

// GetBlueprint retrieve instance of blueprint from usecase
func (bi *BlueprintInteractor) GetBlueprint() (blueprint.Blueprint, error) {
	b, err := bi.repository.GetBlueprint()

	if err != nil {
		return blueprint.Blueprint{}, err
	}

	return b, nil
}
