package models

import (
	"time"
)

type Publisher struct {
	ID 			int	      		`gorm:"AUTO_INCREMENT,primary_key"json:"id"`
	Name 			string        		`json:"name"`
	Email 			string         		`json:"email"`
	Address 		string                  `json:"address"`
	Phone			int                     `json:"phone"`
	Book			[]Books
	CreatedAt 		time.Time           	`json:"createdAt"`
	UpdatedAt 		time.Time
}
