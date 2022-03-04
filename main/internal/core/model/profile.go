package model

import (
	"gorm.io/gorm"
	"main/internal/pkg/objconv"
)

type Profile struct {
	gorm.Model   `json:"-"`
	UserID       uint     `json:"user_id"`
	Nickname     string   `gorm:"column:nickname;not null" json:"nickname,omitempty"`
	ProfileImage string   `gorm:"column:profile_image;default:'photo/no_image.png'" json:"profile_image,omitempty"`
	Email        string   `gorm:"column:email" json:"email,omitempty"`
	Phone        string   `gorm:"column:phone" json:"phone,omitempty"`
	UserURL      string   `gorm:"column:user_url" json:"user_url,omitempty"`
	IntroMessage string   `gorm:"column:intro_message" json:"intro_message"`
	Location     Location `gorm:"embedded;embeddedPrefix:location_" json:"location,omitempty"`
}

type Location struct {
	Country string `json:"country"`
	City    string `json:"city"`
}

func NewProfile(userID uint, nickname string) *Profile {
	return &Profile{
		UserID:   userID,
		Nickname: nickname,
	}
}

func (p *Profile) GetID() uint {
	return p.ID
}

func (p *Profile) GetOwnerID() uint {
	return p.UserID
}

func (p *Profile) ToMap() map[string]interface{} {
	data := map[string]interface{}{
		"user_id":       p.UserID,
		"nickname":      p.Nickname,
		"profile_image": p.ProfileImage,
		"email":         p.Email,
		"user_url":      p.UserURL,
		"intro_message": p.IntroMessage,
		"location": map[string]string{
			"country": p.Location.Country,
			"city":    p.Location.City,
		},
	}
	return data
}

func (p *Profile) ToMapWithFields(fields []string) map[string]interface{} {
	return objconv.ToMapWithFields(p, fields)
}
