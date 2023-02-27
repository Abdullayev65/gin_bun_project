package service

import (
	"github.com/Abdullayev65/gin_bun_project/pkg/models"
)

func (s *Service) PostAdd(post *models.Post, attachmentIDs []int) error {
	_, err := s.DB.NewInsert().Model(post).Exec(s.ctx)
	if err != nil {
		return err
	}
	if len(attachmentIDs) > 0 {
		postAttachs := make([]models.PostAttachment, 0, len(attachmentIDs))
		for _, aID := range attachmentIDs {
			postAttachs = append(postAttachs, models.PostAttachment{PostID: post.ID, AttachmentID: aID})
		}
		_, err = s.DB.NewInsert().Model(&postAttachs).Exec(s.ctx)
		if err != nil {
			return err
		}
		for _, aID := range attachmentIDs {
			post.Attachments = append(post.Attachments, models.Attachment{ID: aID})
		}
	}
	return nil
}
func (s *Service) PostGet(id int) (*models.Post, error) {
	post := &models.Post{ID: id}
	err := s.DB.NewSelect().Model(post).
		Relation("Attachments").
		WherePK().Scan(s.ctx)
	if err != nil {
		return nil, err
	}
	return post, nil
}
func (s *Service) PostAll() []models.Post {
	var posts []models.Post
	s.DB.NewSelect().Model(&posts).
		Relation("Attachments").Scan(s.ctx)
	return posts
}
func (s *Service) PostAllByUserID(userID int) []models.Post {
	var posts []models.Post
	s.DB.NewSelect().Model(&posts).
		Relation("Attachments").
		Where("user_id = ", userID).Scan(s.ctx)
	return posts
}

func (s *Service) PostUpdate(post *models.Post, attachmentIDs []int) error {
	_, err := s.DB.NewUpdate().Model(post).
		Column("description").
		WherePK().Exec(s.ctx)

	if len(attachmentIDs) > 0 {
		postAttachs := make([]models.PostAttachment, 0, len(attachmentIDs))
		for _, aID := range attachmentIDs {
			postAttachs = append(postAttachs, models.PostAttachment{PostID: post.ID, AttachmentID: aID})
		}
		_, err = s.DB.NewInsert().Model(postAttachs).Exec(s.ctx)
		if err != nil {
			return err
		}
	}

	if err != nil {
		return err
	}

	return nil
}
func (s *Service) PostDelete(id int) error {
	_, err := s.DB.NewDelete().Model((*models.Post)(nil)).
		Where("id = ?", id).Exec(s.ctx)
	if err != nil {
		return err
	}
	return nil
}
