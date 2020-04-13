package usecase

import (
	commentModel "github.com/Ubidy/Ubidy_AgencyNotificationAPI/domain/comment"
	"github.com/Ubidy/Ubidy_AgencyNotificationAPI/usecase/repository"
	stream "gopkg.in/GetStream/stream-go2.v1"
)

// CommentInteractor implementation of comment usecase
type CommentInteractor struct {
	repository repository.Comment
}

// NewCommentInteractor create new instance of comment interactor
func NewCommentInteractor(repository repository.Comment) *CommentInteractor {
	return &CommentInteractor{
		repository: repository,
	}
}

// AddCommentToCandidateActivity create an activity to employer when employer commented on the candidate
func (ci *CommentInteractor) AddCommentToCandidateActivity(comment commentModel.AuctionComment) (stream.Activity, error) {
	activity, err := ci.repository.AddCommentToCandidateActivity(comment)

	if err != nil {
		return stream.Activity{}, err
	}

	return activity, nil
}

// AddCommentToCandidateSuccessFeeActivity create an activity to employer when employer commented on the candidate for success fee
func (ci *CommentInteractor) AddCommentToCandidateSuccessFeeActivity(comment commentModel.SuccessFeeComment) (stream.Activity, error) {
	activity, err := ci.repository.AddCommentToCandidateSuccessFeeActivity(comment)

	if err != nil {
		return stream.Activity{}, err
	}

	return activity, nil
}
