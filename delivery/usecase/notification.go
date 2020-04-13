package usecase

import (
	stream "gopkg.in/GetStream/stream-go2.v1"
)

// NotificationInteractor interface for notification interactor
type NotificationInteractor interface {
	GetNotifications(userID string) ([]stream.Activity, error)
	UpdateNotificationArchive(userID string, feedID string, isArchive bool) (stream.Activity, error)
	UpdateNotificationViewed(userID string, feedID string, isViewed bool) (stream.Activity, error)
}
