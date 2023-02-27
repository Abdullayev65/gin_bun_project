package io

import (
	"github.com/Abdullayev65/gin_bun_project/pkg/models"
	"strconv"
	"time"
)

type PostOutput struct {
	ID          int                `json:"id"`
	Description string             `json:"description"`
	UserID      int                `json:"userID"`
	CreatedAt   time.Time          `json:"createdAt"`
	Attachments []AttachmentOutput `json:"attachments"`
}
type AttachmentOutput struct {
	ID  int    `json:"id"`
	URL string `json:"url"`
}

func NewPostOutput(p *models.Post) *PostOutput {
	po := PostOutput{ID: p.ID, Description: p.Description,
		UserID: p.UserID, CreatedAt: p.CreatedAt}
	if len(p.Attachments) > 0 {
		atachs := make([]AttachmentOutput, 0, len(p.Attachments))
		for _, a := range p.Attachments {
			atachs = append(atachs, AttachmentOutput{ID: a.ID, URL: "http://localhost:8080/attachment/" + strconv.Itoa(a.ID)})
		}
		po.Attachments = atachs
	}
	return &po
}
