package models

import (
	"github.com/uptrace/bun"
	"time"
)

type User struct {
	ID       int    `bun:",pk,autoincrement"`
	Username string `bun:",unique,notnull"`
	Password string
}

type Post struct {
	bun.BaseModel `bun:"table:post,alias:p"`
	ID            int `bun:",pk,autoincrement"`
	Description   string
	UserID        int
	User          *User        `bun:"rel:belongs-to,join:user_id=id" json:"-"`
	CreatedAt     time.Time    `bun:",nullzero,notnull,default:current_timestamp"`
	Attachments   []Attachment `bun:"m2m:post_attachment,join:Post=Attachment"`
}

type Comment struct {
	bun.BaseModel `bun:"table:comment,alias:c"`
	ID            int    `bun:",pk,autoincrement"`
	Text          string `bun:",nullzero,notnull"`
	UserID        int
	User          *User `bun:"rel:belongs-to,join:user_id=id" json:"-"`
	PostID        int
	Post          *Post `bun:"rel:belongs-to,join:post_id=id" json:"-"`
}

type Attachment struct {
	bun.BaseModel `bun:"table:attachment"`
	ID            int       `bun:",pk,autoincrement" json:"id"`
	Path          string    `bun:",nullzero,notnull,unique" json:"-"`
	FileName      string    `bun:",nullzero" json:"name"`
	CreateAt      time.Time `bun:",nullzero,notnull,default:current_timestamp" json:"create_at"`
}

//type Order struct {
//	ID int64 `bun:",pk"`
//	// Order and Item in join:Order=Item are fields in OrderToItem model
//	Items []Item `bun:"m2m:order_to_items,join:Order=Item"`
//}
//
//type Item struct {
//	ID int64 `bun:",pk"`
//}

type PostAttachment struct {
	bun.BaseModel `bun:"table:post_attachment"`
	Post          *Post       `bun:"rel:belongs-to,join:post_id=id"`
	PostID        int         `bun:",pk"`
	Attachment    *Attachment `bun:"rel:belongs-to,join:attachment_id=id"`
	AttachmentID  int         `bun:",pk"`
}

//type OrderToItem struct {
//	OrderID int64  `bun:",pk"`
//	Order   *Order `bun:"rel:belongs-to,join:order_id=id"`
//	ItemID  int64  `bun:",pk"`
//	Item    *Item  `bun:"rel:belongs-to,join:item_id=id"`
//}
