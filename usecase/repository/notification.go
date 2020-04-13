package repository

import (
	stream "gopkg.in/GetStream/stream-go2.v1"
)

// Notification interface for notification repository
type Notification interface {
	GetNotifications(userID string) ([]stream.Activity, error)
	UpdateNotificationArchive(userID string, feedID string, isArchive bool) (stream.Activity, error)
	UpdateNotificationViewed(userID string, feedID string, isViewed bool) (stream.Activity, error)
}
