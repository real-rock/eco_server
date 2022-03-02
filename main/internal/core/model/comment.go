package model

import (
	"gorm.io/gorm"
	"main/internal/pkg/objconv"
)

type Comments []Comment

type Comment struct {
	gorm.Model `json:"-" structs:"-"`
	UserID     uint    `json:"user_id"`
	QuantID    uint    `json:"quant_id"`
	Content    string  `gorm:"type:text;column:content" json:"content"`
	Replies    Replies `gorm:"constraint:OnDelete:CASCADE;" json:"replies"`
}

func NewComment(userID, quantID uint, content string) *Comment {
	return &Comment{
		UserID:  userID,
		QuantID: quantID,
		Content: content,
	}
}

func (c *Comment) GetID() uint {
	return c.ID
}

func (c *Comment) GetOwnerID() uint {
	return c.UserID
}

func (c *Comment) ToMap() map[string]interface{} {
	return objconv.ToMap(c)
}

func (c *Comment) ToMapWithFields(fields []string) map[string]interface{} {
	return objconv.ToMapWithFields(c, fields)
}
