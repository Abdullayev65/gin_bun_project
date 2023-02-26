package service

import (
	"github.com/Abdullayev65/gin_bun_project/pkg/models"
)

func (s *Service) AttachmentAddAll(attachments []models.Attachment) error {
	_, err := s.DB.NewInsert().
		Model(&attachments).Exec(s.ctx)
	if err != nil {
		return err
	}
	return nil
}
func (s *Service) AttachmentGet(id int) (*models.Attachment, error) {
	post := &models.Attachment{ID: id}
	err := s.DB.NewSelect().Model(post).WherePK().Scan(s.ctx)
	if err != nil {
		return nil, err
	}
	return post, nil
}
