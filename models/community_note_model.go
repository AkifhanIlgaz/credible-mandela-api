package models

import (
	"errors"
	"time"

	"github.com/AkifhanIlgaz/credible-mandela-api/utils/constants"
	"github.com/AkifhanIlgaz/credible-mandela-api/utils/message"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CommunityNote struct {
	Id                 primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Publisher          string             `json:"publisher" bson:"publisher"`
	CreatedAt          time.Time          `json:"createdAt" bson:"createdAt"`
	Title              string             `json:"title" bson:"title"`
	Content            string             `json:"content" bson:"content"`
	CoverImageIPFSHash string             `json:"coverImageIPFSHash" bson:"coverImageIPFSHash"`
	IsLiked            bool               `json:"isLiked" bson:"-"`
	LikeCount          uint               `json:"likeCount" bson:"-"`
}

type PublishCommunityNoteForm struct {
	Title              string `json:"title" binding:"required"`
	Content            string `json:"content" binding:"required"`
	CoverImageIPFSHash string `json:"coverImageIPFSHash"`
}

type PublishCommunityNoteResponse struct {
	CommunityNoteId string `json:"communityNoteId"`
}

func (form PublishCommunityNoteForm) Validate() error {
	titleLength := len(form.Title)

	if titleLength > constants.MaxTitleLength {
		return errors.New(message.TitleTooLong)
	}

	contentLength := len(form.Content)

	if contentLength > constants.MaxContentLength {
		return errors.New(message.ContentTooLong)
	}

	return nil
}

type Like struct {
	Id              primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	CommunityNoteId string             `json:"communityNoteId" bson:"communityNoteId"`
	Username        string             `json:"username" bson:"username"`
	CreatedAt       time.Time          `json:"createdAt" bson:"createdAt"`
}
