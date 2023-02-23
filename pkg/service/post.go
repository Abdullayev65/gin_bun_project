package service

import (
	"github.com/Abdullayev65/gin_bun_project/pkg/models"
)

func (s *Service) PostAdd(post *models.Post) error {
	_, err := s.DB.NewInsert().Model(post).Exec(s.ctx)
	if err != nil {
		return err
	}
	return nil
}
func (s *Service) PostGet(id int64) (*models.Post, error) {
	post := &models.Post{ID: id}
	err := s.DB.NewSelect().Model(post).WherePK().Scan(s.ctx)
	if err != nil {
		return nil, err
	}
	return post, nil
}
func (s *Service) PostAll() []models.Post {
	var posts []models.Post
	s.DB.NewSelect().Model(&posts).Scan(s.ctx)
	return posts
}

func (s *Service) PostUpdate(post *models.Post) error {
	_, err := s.DB.NewUpdate().Model(post).
		Column("description").
		WherePK().Exec(s.ctx)
	if err != nil {
		return err
	}
	return nil
}
func (s *Service) PostDelete(id int64) error {
	_, err := s.DB.NewDelete().Model((*models.Post)(nil)).
		Where("id = ?", id).Exec(s.ctx)
	if err != nil {
		return err
	}
	return nil
}
