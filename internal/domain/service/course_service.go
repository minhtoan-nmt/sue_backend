package service

import (
	"sue_backend/internal/domain/model"
	"sue_backend/internal/domain/repository"
)

type CourseService struct {
	repo *repository.CourseRepository
}

func NewCourseService(repo *repository.CourseRepository) *CourseService {
	return &CourseService{repo: repo}
}

func (s *CourseService) Create(course *model.Course) (int64, error) {
	courseCreated, err := s.repo.Create(course)
	if err != nil {
		return 0, err
	}
	return courseCreated, nil
}

func (s *CourseService) Get(id int64) (*model.Course, error) {
	return s.repo.FindByID(id)
}

func (s *CourseService) List(page, limit int) ([]*model.Course, int, error) {
	return s.repo.FindAll(page, limit)
}

func (s *CourseService) Update(course *model.Course) error {
	return s.repo.Update(course)
}

func (s *CourseService) Delete(id int64) error {
	return s.repo.Delete(id)
}
