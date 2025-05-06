package repository

import (
	"context"
	"sue_backend/internal/domain/model"
	"sue_backend/internal/infra/cache"
	"sue_backend/internal/infra/db"
	"time"
)

type CourseRepository struct {
	store db.Store
	cache cache.Store
}

func NewCourseRepository(store db.Store, cache cache.Store) *CourseRepository {
	return &CourseRepository{
		store: store,
		cache: cache,
	}
}
func (r *CourseRepository) Create(course *model.Course) (int64, error) {
	query := `INSERT INTO courses (name, template_id, schedule, status, start_date, end_date, created_at, updated_at)
			  VALUES ($1, $2, $3, $4, $5, $6, NOW(), NOW()) RETURNING id`
	ctx := context.Background()
	rows, err := r.store.ExecQuery(ctx, query,
		course.Name,
		course.TemplateID,
		course.Schedule,
		course.Status,
		course.StartDate,
		course.EndDate,
		&course.ID,
	)
	if err != nil {
		return 0, err
	}
	if len(rows) == 0 {
		return 0, nil
	}
	return rows[0]["id"].(int64), err
}

func (r *CourseRepository) FindByID(id int64) (*model.Course, error) {
	query := `SELECT * FROM courses WHERE id = $1`
	rows, err := r.store.ExecQuery(context.Background(), query, id)
	if err != nil {
		return nil, err
	}
	if len(rows) == 0 {
		return nil, nil
	}
	row := rows[0]
	course := &model.Course{
		ID:         row["id"].(int64),
		Name:       row["name"].(string),
		TemplateID: row["template_id"].(int64),
		Schedule:   row["schedule"].(*string),
		Status:     row["status"].(string),
		StartDate:  row["start_date"].(*time.Time),
		EndDate:    row["end_date"].(*time.Time),
		CreatedAt:  row["created_at"].(time.Time),
		UpdatedAt:  row["updated_at"].(time.Time),
	}
	return course, nil
}

func (r *CourseRepository) FindAll(page, limit int) ([]*model.Course, int, error) {
	offset := (page - 1) * limit
	query := `SELECT * FROM courses ORDER BY created_at DESC LIMIT $1 OFFSET $2`
	rows, err := r.store.ExecQuery(context.Background(), query, limit, offset)
	if err != nil {
		return nil, 0, err
	}
	courses := make([]*model.Course, len(rows))
	for i, row := range rows {
		courses[i] = &model.Course{
			ID:         row["id"].(int64),
			Name:       row["name"].(string),
			TemplateID: row["template_id"].(int64),
			Schedule:   row["schedule"].(*string),
			Status:     row["status"].(string),
			StartDate:  row["start_date"].(*time.Time),
			EndDate:    row["end_date"].(*time.Time),
			CreatedAt:  row["created_at"].(time.Time),
			UpdatedAt:  row["updated_at"].(time.Time),
		}
	}
	countQuery := `SELECT COUNT(*) FROM courses`
	countRows, err := r.store.ExecQuery(context.Background(), countQuery)
	if err != nil || len(countRows) == 0 {
		return nil, 0, err
	}
	total := int(countRows[0]["count"].(int64))
	if total == 0 {
		return nil, 0, nil
	}
	return courses, total, nil
}

func (r *CourseRepository) Update(course *model.Course) error {
	// UPDATE query
	query := `UPDATE courses SET name = $1, template_id = $2, schedule = $3, status = $4, start_date = $5, end_date = $6, updated_at = NOW() WHERE id = $7`
	ctx := context.Background()
	return r.store.Exec(ctx, query,
		course.Name,
		course.TemplateID,
		course.Schedule,
		course.Status,
		course.StartDate,
		course.EndDate,
		course.ID,
	)

}

func (r *CourseRepository) Delete(id int64) error {
	/// Delete is set status to "deleted"
	query := `UPDATE courses SET status = 'deleted', updated_at = NOW() WHERE id = $1`
	ctx := context.Background()
	return r.store.Exec(ctx, query, id)
}
