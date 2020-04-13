package usecase

import (
	commentModel "github.com/Ubidy/Ubidy_AgencyNotificationAPI/domain/comment"
	stream "gopkg.in/GetStream/stream-go2.v1"
)

// CommentInteractor interface for comment usecase
type CommentInteractor interface {
	AddCommentToCandidateActivity(comment commentModel.AuctionComment) (stream.Activity, error)
	AddCommentToCandidateSuccessFeeActivity(comment commentModel.SuccessFeeComment) (stream.Activity, error)
}
