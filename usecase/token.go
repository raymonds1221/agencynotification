package usecase

import "github.com/Ubidy/Ubidy_AgencyNotificationAPI/usecase/repository"

// TokenInteractor implementation for token usecase
type TokenInteractor struct {
	repository repository.Token
}

// NewTokenInteractor create new instance of token repository
func NewTokenInteractor(r repository.Token) *TokenInteractor {
	return &TokenInteractor{
		repository: r,
	}
}

// GetToken retrieve token from activity stream
func (t *TokenInteractor) GetToken(userID string) string {
	token := t.repository.GetToken(userID)

	return token
}

// GetUnreadNotificationCount unread activities from activity stream
func (t *TokenInteractor) GetUnreadNotificationCount(userID string) int {
	unread := t.repository.GetUnreadNotificationCount(userID)

	return unread
}
