package repository

import (
	"log"

	stream "gopkg.in/GetStream/stream-go2.v1"
)

// Notification implementation of notification repository
type Notification struct {
	client *stream.Client
}

// NewNotificationRepository create new instance of notification repository
func NewNotificationRepository(client *stream.Client) *Notification {
	return &Notification{
		client: client,
	}
}

// GetNotifications retrieve list of notifications
func (n *Notification) GetNotifications(userID string) ([]stream.Activity, error) {
	agencyNotification := n.client.NotificationFeed("agencynotification", userID)

	resp, err := agencyNotification.GetActivities(
		stream.WithActivitiesLimit(50),
		stream.WithNotificationsMarkRead(true),
	)

	if err != nil {
		return nil, err
	}

	var activities []stream.Activity

	for _, res := range resp.Results {
		for _, activity := range res.Activities {
			activities = append(activities, activity)
		}
	}

	for {
		if len(activities) >= 50 {
			break
		}

		resp, err = agencyNotification.GetNextPageActivities(resp)

		if resp == nil || err != nil {
			break
		}

		for _, res := range resp.Results {
			for _, activity := range res.Activities {
				activities = append(activities, activity)
			}
		}
	}

	return activities, nil
}

// UpdateNotificationArchive update the arhive of specific notification
func (n *Notification) UpdateNotificationArchive(userID string, feedID string, isArchive bool) (stream.Activity, error) {
	agencyNotification := n.client.NotificationFeed("agencynotification", userID)
	resp, err := agencyNotification.GetActivities()

	if err != nil {
		return stream.Activity{}, err
	}

	var activities []stream.Activity

	for _, res := range resp.Results {
		for _, activity := range res.Activities {
			activities = append(activities, activity)
		}
	}

	for {
		resp, err = agencyNotification.GetNextPageActivities(resp)

		if resp == nil || err != nil {
			break
		}

		for _, res := range resp.Results {
			for _, activity := range res.Activities {
				activities = append(activities, activity)
			}
		}
	}

	log.Printf("activities: %d, feedID: %s\n", len(activities), feedID)

	activity := n.findActivity(activities, func(activity stream.Activity) bool {
		return activity.ID == feedID
	})

	if activity != nil {
		activity.Extra["isArchive"] = isArchive
		n.client.UpdateActivities(*activity)
		return *activity, nil
	}

	return stream.Activity{}, nil
}

// UpdateNotificationViewed update the view of specific notification
func (n *Notification) UpdateNotificationViewed(userID string, feedID string, isViewed bool) (stream.Activity, error) {
	agencyNotification := n.client.NotificationFeed("agencynotification", userID)
	resp, err := agencyNotification.GetActivities()

	log.Printf("userID: %s, feedID: %s, isViewed: %v", userID, feedID, isViewed)

	if err != nil {
		return stream.Activity{}, err
	}

	var activities []stream.Activity

	for _, res := range resp.Results {
		for _, activity := range res.Activities {
			activities = append(activities, activity)
		}
	}

	for {
		resp, err = agencyNotification.GetNextPageActivities(resp)

		if resp == nil || err != nil {
			break
		}

		for _, res := range resp.Results {
			for _, activity := range res.Activities {
				activities = append(activities, activity)
			}
		}
	}

	activity := n.findActivity(activities, func(activity stream.Activity) bool {
		return activity.ID == feedID
	})

	if activity != nil {
		activity.Extra["isViewed"] = isViewed
		n.client.UpdateActivities(*activity)
		return *activity, nil
	}

	return stream.Activity{}, nil
}

func (n *Notification) findActivity(activities []stream.Activity, condition func(a stream.Activity) bool) *stream.Activity {
	for _, activity := range activities {
		if condition(activity) {
			return &activity
		}
	}
	return nil
}
