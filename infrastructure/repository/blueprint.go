package repository

import "github.com/Ubidy/Ubidy_AgencyNotificationAPI/domain/blueprint"

// Blueprint implementation of Blueprint repository
type Blueprint struct{}

// NewBlueprintRepository create new instance of Blueprint repository
func NewBlueprintRepository() *Blueprint {
	return &Blueprint{}
}

// GetBlueprint retrieve new instance of blueprint domain
func (br *Blueprint) GetBlueprint() (blueprint.Blueprint, error) {
	b := blueprint.Blueprint{
		ID:   "blueprind_id_1",
		Name: "blueprint_name_1",
		Desc: "blueprint_desc_1",
	}

	return b, nil
}
