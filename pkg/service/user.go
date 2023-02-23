package service

import (
	"fmt"
	"github.com/Abdullayev65/gin_bun_project/pkg/models"
)

func (s *Service) UserAdd(user *models.User) error {
	_, err := s.DB.NewInsert().Model(user).Exec(s.ctx)
	if err != nil {
		return err
	}
	return nil
}
func (s *Service) UserGet(id int64) (*models.User, error) {
	user := &models.User{ID: id}
	err := s.DB.NewSelect().Model(user).WherePK().Scan(s.ctx)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *Service) UserAll() []models.User {
	var users []models.User
	s.DB.NewSelect().Model(&users).Scan(s.ctx)
	return users
}

func (s *Service) UserUpdate(user *models.User) error {
	res, err := s.DB.NewUpdate().Model(user).WherePK().Exec(s.ctx)
	fmt.Println(res)
	if err != nil {
		return err
	}
	return nil
}
func (s *Service) UserDelete(id int64) error {
	_, err := s.DB.NewDelete().Model((*models.User)(nil)).
		Where("id = ?", id).Exec(s.ctx)
	if err != nil {
		return err
	}
	return nil
}
func (s *Service) UserByUsername(username string) (*models.User, error) {
	user := new(models.User)
	err := s.DB.NewSelect().Model(user).Where("username = ?", username).Scan(s.ctx)
	if err != nil {
		return nil, err
	}
	return user, nil
}
