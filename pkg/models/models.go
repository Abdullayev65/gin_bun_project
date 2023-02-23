package models

import "time"

type User struct {
	ID       int64  `bun:",pk,autoincrement"`
	Username string `bun:",unique,notnull"`
	Password string
}

type Post struct {
	ID          int64 `bun:",pk,autoincrement"`
	Description string
	UserID      int64
	CreatedAt   time.Time `bun:",nullzero,notnull,default:current_timestamp"`
}

type Comment struct {
	ID     int64  `bun:",pk,autoincrement"`
	Text   string `bun:",nullzero,notnull"`
	UserID int64
	PostID int64
}
