package repository

import (
	"context"
	"fmt"
	"sue_backend/internal/domain/model"
	"sue_backend/internal/infra/cache"
	"sue_backend/internal/infra/db"
	"time"
)

type CourseTemplateRepository struct {
	store db.Store
	cache cache.Store
}

func NewCourseTemplateRepository(store db.Store, cache cache.Store) *CourseTemplateRepository {
	return &CourseTemplateRepository{
		store: store,
		cache: cache,
	}
}

func (r *CourseTemplateRepository) Create(ct *model.CourseTemplate) (int64, error) {
	query := `
		INSERT INTO templates 
			(name, description, type, level, language, price, discount, duration, capacity, created_at, updated_at)
		VALUES 
			($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
		RETURNING id`

	now := time.Now()
	ct.CreatedAt = now
	ct.UpdatedAt = now

	rows, err := r.store.ExecQuery(
		context.Background(), query,
		ct.Name,
		ct.Description,
		ct.Type,
		ct.Level,
		ct.Language,
		ct.Price,
		ct.Discount,
		ct.Duration,
		ct.Capacity,
		ct.CreatedAt,
		ct.UpdatedAt,
	)
	if err != nil {
		return 0, err
	}
	if len(rows) == 0 || rows[0]["id"] == nil {
		return 0, fmt.Errorf("no ID returned")
	}

	id, ok := rows[0]["id"].(int64)
	if !ok {
		return 0, fmt.Errorf("unexpected ID type")
	}
	return id, nil
}

func (r *CourseTemplateRepository) FindAll(page, limit int) ([]*model.CourseTemplate, int, error) {
	offset := (page - 1) * limit
	query := `SELECT * FROM templates ORDER BY created_at DESC LIMIT $1 OFFSET $2`
	rows, err := r.store.ExecQuery(context.Background(), query, limit, offset)
	if err != nil {
		return nil, 0, err
	}

	var templates []*model.CourseTemplate
	for _, row := range rows {
		templates = append(templates, mapToCourseTemplate(row))
	}

	countQuery := `SELECT COUNT(*) FROM templates`
	countRows, err := r.store.ExecQuery(context.Background(), countQuery)
	if err != nil {
		return nil, 0, err
	}
	count := int64(0)
	if len(countRows) > 0 {
		count = countRows[0]["count"].(int64)
	}

	return templates, int(count), nil
}

func (r *CourseTemplateRepository) FindByID(id int64) (*model.CourseTemplate, error) {
	query := `SELECT * FROM templates WHERE id = $1`
	rows, err := r.store.ExecQuery(context.Background(), query, id)
	if err != nil {
		return nil, err
	}
	if len(rows) == 0 {
		return nil, nil
	}
	return mapToCourseTemplate(rows[0]), nil
}

func (r *CourseTemplateRepository) Update(ct *model.CourseTemplate) error {
	query := `
		UPDATE templates SET 
			name = $1,
			description = $2,
			type = $3,
			level = $4,
			language = $5,
			price = $6,
			discount = $7,
			duration = $8,
			capacity = $9,
			updated_at = $10
		WHERE id = $11`

	ct.UpdatedAt = time.Now()

	return r.store.Exec(
		context.Background(), query,
		ct.Name,
		ct.Description,
		ct.Type,
		ct.Level,
		ct.Language,
		ct.Price,
		ct.Discount,
		ct.Duration,
		ct.Capacity,
		ct.UpdatedAt,
		ct.ID,
	)
}

func (r *CourseTemplateRepository) Delete(id int64) error {
	// Set the status to "deleted" instead of actually deleting the record
	query := `UPDATE templates SET deleted_at = $1 WHERE id = $2`
	now := time.Now()
	return r.store.Exec(context.Background(), query, now, id)
}

func mapToCourseTemplate(row map[string]interface{}) *model.CourseTemplate {
	getStringPtr := func(val interface{}) *string {
		if s, ok := val.(string); ok {
			return &s
		}
		return nil
	}
	getFloatPtr := func(val interface{}) *float64 {
		if f, ok := val.(float64); ok {
			return &f
		}
		return nil
	}
	getIntPtr := func(val interface{}) *int {
		if i, ok := val.(int64); ok {
			x := int(i)
			return &x
		}
		return nil
	}
	getInt64Ptr := func(val interface{}) *int64 {
		if i, ok := val.(int64); ok {
			return &i
		}
		return nil
	}

	return &model.CourseTemplate{
		ID:          row["id"].(int64),
		Name:        row["name"].(string),
		Description: getStringPtr(row["description"]),
		Status:      row["status"].(string),
		CreatedBy:   getInt64Ptr(row["created_by"]),
		Type:        row["type"].(string),
		Level:       row["level"].(string),
		Language:    row["language"].(string),
		Image:       getStringPtr(row["image"]),
		Price:       getFloatPtr(row["price"]),
		Discount:    getFloatPtr(row["discount"]),
		Duration:    getStringPtr(row["duration"]),
		Capacity:    getIntPtr(row["capacity"]),
		CreatedAt:   row["created_at"].(time.Time),
		UpdatedAt:   row["updated_at"].(time.Time),
	}
}
