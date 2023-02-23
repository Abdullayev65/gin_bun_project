package service

import (
	"github.com/Abdullayev65/gin_bun_project/pkg/models"
)

func (s *Service) CommentAdd(comment *models.Comment) error {
	_, err := s.DB.NewInsert().Model(comment).Exec(s.ctx)
	if err != nil {
		return err
	}
	return nil
}
func (s *Service) CommentGet(id int64) (*models.Comment, error) {
	comment := &models.Comment{ID: id}
	err := s.DB.NewSelect().Model(comment).WherePK().Scan(s.ctx)
	if err != nil {
		return nil, err
	}
	return comment, nil
}
func (s *Service) CommentAll() []models.Comment {
	var comments []models.Comment
	s.DB.NewSelect().Model(&comments).Scan(s.ctx)
	return comments
}
func (s *Service) CommentUpdate(comment *models.Comment) error {
	_, err := s.DB.NewUpdate().Model(comment).
		Column("text").
		WherePK().Exec(s.ctx)
	if err != nil {
		return err
	}
	return nil
}
func (s *Service) CommentDelete(id int64) error {
	_, err := s.DB.NewDelete().Model((*models.Comment)(nil)).
		Where("id = ?", id).Exec(s.ctx)
	if err != nil {
		return err
	}
	return nil
}
