package models

import (
	"github.com/uptrace/bun"
	"time"
)

type User struct {
	ID       int64  `bun:",pk,autoincrement"`
	Username string `bun:",unique,notnull"`
	Password string
}

type Post struct {
	bun.BaseModel `bun:"table:post,alias:p"`
	ID            int64 `bun:",pk,autoincrement"`
	Description   string
	UserID        int64
	User          *User     `bun:"rel:belongs-to,join:user_id=id" json:"-"`
	CreatedAt     time.Time `bun:",nullzero,notnull,default:current_timestamp"`
}

type Comment struct {
	bun.BaseModel `bun:"table:comment,alias:c"`
	ID            int64  `bun:",pk,autoincrement"`
	Text          string `bun:",nullzero,notnull"`
	UserID        int64
	User          *User `bun:"rel:belongs-to,join:user_id=id" json:"-"`
	PostID        int64
	Post          *Post `bun:"rel:belongs-to,join:post_id=id" json:"-"`
}
