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
	user := &models.Post{ID: id}
	err := s.DB.NewSelect().Model(user).WherePK().Scan(s.ctx)
	if err != nil {
		return nil, err
	}
	return user, nil
}
func (s *Service) PostAll() []models.Post {
	var users []models.Post
	s.DB.NewSelect().Model(&users).Scan(s.ctx)
	return users
}

func (s *Service) PostUpdate(user *models.Post) error {
	_, err := s.DB.NewUpdate().Model(user).WherePK().Exec(s.ctx)
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
