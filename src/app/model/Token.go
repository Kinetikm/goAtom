package model

import (
	"github.com/jinzhu/gorm"
	"github.com/satori/go.uuid"
)

type Token struct {
	gorm.Model

	Token         	 uuid.UUID	`sql:"not null`
	UserID			 uint		`sql:"not null`
	User             User 		`gorm:"ForeignKey:UserID"`
}

func (Token) TableName() string {
	return "token"
}