package usecase

// TokenInteractor usecase declaration for token
type TokenInteractor interface {
	GetToken(userID string) string
	GetUnreadNotificationCount(userID string) int
}
