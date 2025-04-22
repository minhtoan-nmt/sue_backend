package service

import (
	"sue_backend/internal/domain/model"
	"sue_backend/internal/domain/repository"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(r *repository.UserRepository) *UserService {
	return &UserService{repo: r}
}

func (s *UserService) Create(user *model.User) (*model.User, error) {
	if err := s.repo.Create(user); err != nil {
		return nil, err
	}
	return user, nil
}

func (s *UserService) GetByID(id int64) (*model.User, error) {
	return s.repo.FindByID(id)
}

func (s *UserService) List() ([]*model.User, error) {
	return s.repo.FindAll()
}

func (s *UserService) Update(user *model.User) error {
	return s.repo.Update(user)
}

func (s *UserService) Delete(id int64) error {
	return s.repo.Delete(id)
}

func (s *UserService) GetPaginatedUsers(page, limit int) ([]*model.User, int, error) {
	return s.repo.FindPaginated(page, limit)
}
