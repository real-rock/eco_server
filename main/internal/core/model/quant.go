package model

import (
	"gorm.io/gorm"
	"main/internal/pkg/objconv"
)

type Quants []Quant

type Quant struct {
	gorm.Model  `json:"-" swaggerignore:"true"`
	UserID      uint        `json:"user_id"`
	QuantOption QuantOption `gorm:"constraint:OnDelete:CASCADE;foreignKey:QuantID;references:ID" json:"-" swaggerignore:"true"`
	ResDataID   string      `gorm:"column:data_id" json:"-" swaggerignore:"true"`
	Active      bool        `gorm:"column:active;default:false" json:"-" swaggerignore:"true"`
	LikedUsers  []*User     `gorm:"many2many:user_favorite_quants;" json:"-" swaggerignore:"true"`
	Name        string      `gorm:"column:name;not null;unique" json:"name" example:"quant model name"`
	Description string      `gorm:"column:description" json:"description" example:"quant model description"`
	Private     bool        `gorm:"column:private;default:false" json:"-" swaggerignore:"true"`
	Comments    Comments    `gorm:"constraint:OnDelete:CASCADE;" json:"-" swaggerignore:"true"`
}

func NewQuant(userID uint, name string) *Quant {
	return &Quant{
		UserID: userID,
		Name:   name,
	}
}

func (q *Quant) GetID() uint {
	return q.ID
}

func (q *Quant) GetOwnerID() uint {
	return q.UserID
}

func (q *Quant) ToMap() map[string]interface{} {
	return objconv.ToMap(q)
}

func (q *Quant) ToMapWithFields(fields []string) map[string]interface{} {
	return objconv.ToMapWithFields(q, fields)
}
