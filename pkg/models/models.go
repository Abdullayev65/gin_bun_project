package models

import "time"

type User struct {
	ID       int64  `bun:",pk,autoincrement"`
	Username string `bun:",unique,notnull"`
	Password string
}

type Post struct {
	ID          int64
	Description string
	UserID      int64
	CreatedAt   time.Time
}

type Comment struct {
	ID     int64
	Text   string
	UserID int64
	PostID int64
}
