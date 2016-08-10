package models

import "time"

type User struct {
	ID 			int	      		`gorm:"AUTO_INCREMENT,primary_key"json:"id"`
	Name 			string        		`json:"name"`
	Email 			string         		`json:"email"`
	Password 		string         		`json:"password"`
	CreatedAt 		time.Time         	`json:"createdAt"`
	UpdatedAt 		time.Time
}
