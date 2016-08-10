package models

import "time"

type Books struct {
	ID 			int	      		`gorm:"AUTO_INCREMENT,primary_key"json:"id"`
	Name 			string        		`json:"name"`
	AuthorID		int64
	RiviewID		int64
	PublisherID		int64
	Riview			Riview
	Author			Author
	Publisher		Publisher
	CreatedAt 		time.Time         	`json:"createdAt"`
	UpdatedAt 		time.Time
}
