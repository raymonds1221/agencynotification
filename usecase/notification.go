package usecase

import (
	"github.com/Ubidy/Ubidy_AgencyNotificationAPI/usecase/repository"
	stream "gopkg.in/GetStream/stream-go2.v1"
)

// NotificationInteractor struct for notification interactor
type NotificationInteractor struct {
	repository repository.Notification
}

// NewNotificationInteractor create new instance of repository interactor
func NewNotificationInteractor(repository repository.Notification) *NotificationInteractor {
	return &NotificationInteractor{
		repository: repository,
	}
}

// GetNotifications implementation of get notification usecase
func (ni *NotificationInteractor) GetNotifications(userID string) ([]stream.Activity, error) {
	activities, err := ni.repository.GetNotifications(userID)

	if err != nil {
		return nil, err
	}

	return activities, nil
}

// UpdateNotificationArchive implementation of update notification archive
func (ni *NotificationInteractor) UpdateNotificationArchive(userID string, feedID string, isArchive bool) (stream.Activity, error) {
	activity, err := ni.repository.UpdateNotificationArchive(userID, feedID, isArchive)

	if err != nil {
		return stream.Activity{}, err
	}

	return activity, nil
}

// UpdateNotificationViewed implementation of update notification viewed
func (ni *NotificationInteractor) UpdateNotificationViewed(userID string, feedID string, isViewed bool) (stream.Activity, error) {
	activity, err := ni.repository.UpdateNotificationViewed(userID, feedID, isViewed)

	if err != nil {
		return stream.Activity{}, err
	}

	return activity, nil
}
