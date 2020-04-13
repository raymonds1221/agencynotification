package repository

import (
	commentModel "github.com/Ubidy/Ubidy_AgencyNotificationAPI/domain/comment"
	stream "gopkg.in/GetStream/stream-go2.v1"
)

// Comment interface for comment repository
type Comment interface {
	AddCommentToCandidateActivity(comment commentModel.AuctionComment) (stream.Activity, error)
	AddCommentToCandidateSuccessFeeActivity(comment commentModel.SuccessFeeComment) (stream.Activity, error)
}
