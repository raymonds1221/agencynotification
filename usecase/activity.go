package usecase

import (
	"github.com/Ubidy/Ubidy_AgencyNotificationAPI/usecase/repository"
	stream "gopkg.in/GetStream/stream-go2.v1"
)

// ActivityInteractor implementation of activity usecase
type ActivityInteractor struct {
	repository repository.Activity
}

// NewActivityInteractor create new instance of activity interactor
func NewActivityInteractor(repository repository.Activity) *ActivityInteractor {
	return &ActivityInteractor{
		repository: repository,
	}
}

// GetActivities get all activities using the supplierID
func (ai *ActivityInteractor) GetActivities(supplierID string) ([]stream.Activity, error) {
	activities, err := ai.repository.GetActivities(supplierID)

	if err != nil {
		return nil, err
	}

	return activities, nil
}
