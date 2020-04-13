package repository

import (
	stream "gopkg.in/GetStream/stream-go2.v1"
)

// Activity implementation of activity repository
type Activity struct {
	client *stream.Client
}

// NewActivityRepository create new instance of activity repository
func NewActivityRepository(client *stream.Client) *Activity {
	return &Activity{
		client: client,
	}
}

// GetActivities retrieve list of activities
func (a *Activity) GetActivities(supplierID string) ([]stream.Activity, error) {
	agencyFeed := a.client.FlatFeed("agency", supplierID)

	resp, err := agencyFeed.GetActivities()

	if err != nil {
		return nil, err
	}

	return resp.Results, nil
}
