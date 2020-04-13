package usecase

import (
	stream "gopkg.in/GetStream/stream-go2.v1"
)

// ActivityInteractor interface for activity usecase
type ActivityInteractor interface {
	GetActivities(supplierID string) ([]stream.Activity, error)
}
