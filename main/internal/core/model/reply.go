package model

import (
	"gorm.io/gorm"
	"main/internal/pkg/objconv"
)

type Replies []Reply

type Reply struct {
	gorm.Model `json:"-"`
	CommentID  uint   `json:"comment_id"`
	UserID     uint   `json:"user_id"`
	Content    string `gorm:"type:text;column:content" json:"content"`
}

func NewReply(userID, commentID uint, content string) *Reply {
	return &Reply{
		UserID:    userID,
		CommentID: commentID,
		Content:   content,
	}
}

func (r *Reply) GetID() uint {
	return r.ID
}

func (r *Reply) GetOwnerID() uint {
	return r.UserID
}

func (r *Reply) ToMap() map[string]interface{} {
	return objconv.ToMap(r)
}

func (r *Reply) ToMapWithFields(fields []string) map[string]interface{} {
	return objconv.ToMapWithFields(r, fields)
}
