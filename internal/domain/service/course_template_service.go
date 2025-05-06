// internal/domain/service/course_template_service.go
package service

import (
	"sue_backend/internal/domain/model"
	"sue_backend/internal/domain/repository"
)

type CourseTemplateService struct {
	repo *repository.CourseTemplateRepository
}

func NewCourseTemplateService(r *repository.CourseTemplateRepository) *CourseTemplateService {
	return &CourseTemplateService{r}
}

func (s *CourseTemplateService) Create(ct *model.CourseTemplate) (int64, error) {
	ctCreated, err := s.repo.Create(ct)
	if err != nil {
		return 0, err
	}
	return ctCreated, nil
}

func (s *CourseTemplateService) Get(id int64) (*model.CourseTemplate, error) {
	return s.repo.FindByID(id)
}
func (s *CourseTemplateService) List(page, limit int) ([]*model.CourseTemplate, int, error) {
	return s.repo.FindAll(page, limit)
}
func (s *CourseTemplateService) Update(courseTemplate *model.CourseTemplate) error {

	return s.repo.Update(courseTemplate)
}
func (s *CourseTemplateService) Delete(id int64) error {
	return s.repo.Delete(id)
}
