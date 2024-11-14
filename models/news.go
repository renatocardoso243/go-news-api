package models

import "time"

type News struct {
	ID        	uint      `json:"id" gorm:"primary_key"`
	Author		string    `json:"author"`
	Title     	string    `json:"title"`
	URLToImage 	string    `json:"urlToImage"`
	PublishedAt time.Time `json:"publishedAt"`
	Content   	string    `json:"content"`
}