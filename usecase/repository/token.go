package repository

// Token for getting token from activity stream to be use in notification
type Token interface {
	GetToken(userID string) string
	GetUnreadNotificationCount(userID string) int
}
