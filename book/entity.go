package book

import "time"

type Book struct {
	ID          uint
	Title       string
	Description string
	Price       int
	Rating      int
	CreatedAt   time.Time
	UpdatetAt time.Time
}