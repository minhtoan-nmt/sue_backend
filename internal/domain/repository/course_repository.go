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
	query := `
	SELECT 
		c.id AS course_id, c.name, c.template_id, c.schedule, c.status, c.start_date, c.end_date, 
		c.created_at AS course_created_at, c.updated_at AS course_updated_at,
		ta.id AS ta_id, ta.teacher_id, ta.role, ta.status AS ta_status, 
		ta.start_date AS ta_start_date, ta.end_date AS ta_end_date, 
		ta.created_at AS ta_created_at, ta.updated_at AS ta_updated_at
	FROM courses c
	LEFT JOIN teacher_assignments ta ON ta.course_id = c.id AND ta.status != 'deleted'
	WHERE c.id = $1
	`
	rows, err := r.store.ExecQuery(context.Background(), query, id)
	if err != nil || len(rows) == 0 {
		return nil, err
	}

	// Parse course info (from first row)
	row := rows[0]
	course := &model.Course{
		ID:         row["course_id"].(int64),
		Name:       row["name"].(string),
		TemplateID: row["template_id"].(int64),
		Schedule:   toStrPtr(row["schedule"]),
		Status:     row["status"].(string),
		StartDate:  toTimePtr(row["start_date"]),
		EndDate:    toTimePtr(row["end_date"]),
		CreatedAt:  row["course_created_at"].(time.Time),
		UpdatedAt:  row["course_updated_at"].(time.Time),
	}

	// Parse teacher assignments
	assignments := make([]*model.TeacherAssignment, 0)
	for _, row := range rows {
		if row["ta_id"] == nil {
			continue
		}
		ta := &model.TeacherAssignment{
			ID:        row["ta_id"].(int64),
			TeacherID: row["teacher_id"].(int64),
			CourseID:  id,
			Role:      model.TeacherAssignmentRole(row["role"].(string)),
			Status:    model.TeacherAssignmentStatus(row["ta_status"].(string)),
			StartDate: toTimePtr(row["ta_start_date"]),
			EndDate:   toTimePtr(row["ta_end_date"]),
			CreatedAt: row["ta_created_at"].(time.Time),
			UpdatedAt: row["ta_updated_at"].(time.Time),
		}
		assignments = append(assignments, ta)
	}
	if assignments == nil {
		assignments = make([]*model.TeacherAssignment, 0)
	}
	course.TeacherAssignments = assignments

	return course, nil
}
func (r *CourseRepository) FindAll(page, limit int) ([]*model.Course, int, error) {
	offset := (page - 1) * limit

	query := `
	SELECT 
		c.id AS course_id, c.name, c.template_id, c.schedule, c.status, c.start_date, c.end_date,
		c.created_at AS course_created_at, c.updated_at AS course_updated_at,
		ta.id AS ta_id, ta.teacher_id, ta.role, ta.status AS ta_status,
		ta.start_date AS ta_start_date, ta.end_date AS ta_end_date,
		ta.created_at AS ta_created_at, ta.updated_at AS ta_updated_at
	FROM courses c
	LEFT JOIN teacher_assignments ta ON ta.course_id = c.id AND ta.status != 'deleted'
	ORDER BY c.created_at DESC
	LIMIT $1 OFFSET $2
	`

	rows, err := r.store.ExecQuery(context.Background(), query, limit, offset)
	if err != nil {
		return nil, 0, err
	}

	courseMap := make(map[int64]*model.Course)
	for _, row := range rows {
		courseID := row["course_id"].(int64)
		course, exists := courseMap[courseID]
		if !exists {
			course = &model.Course{
				ID:                 courseID,
				Name:               row["name"].(string),
				TemplateID:         row["template_id"].(int64),
				Schedule:           toStrPtr(row["schedule"]),
				Status:             row["status"].(string),
				StartDate:          toTimePtr(row["start_date"]),
				EndDate:            toTimePtr(row["end_date"]),
				CreatedAt:          row["course_created_at"].(time.Time),
				UpdatedAt:          row["course_updated_at"].(time.Time),
				TeacherAssignments: make([]*model.TeacherAssignment, 0),
			}
			courseMap[courseID] = course
		}

		if row["ta_id"] != nil {
			ta := &model.TeacherAssignment{
				ID:        row["ta_id"].(int64),
				TeacherID: row["teacher_id"].(int64),
				CourseID:  courseID,
				Role:      model.TeacherAssignmentRole(row["role"].(string)),
				Status:    model.TeacherAssignmentStatus(row["ta_status"].(string)),
				StartDate: toTimePtr(row["ta_start_date"]),
				EndDate:   toTimePtr(row["ta_end_date"]),
				CreatedAt: row["ta_created_at"].(time.Time),
				UpdatedAt: row["ta_updated_at"].(time.Time),
			}
			course.TeacherAssignments = append(course.TeacherAssignments, ta)
		}
	}

	// Convert map to slice
	courses := make([]*model.Course, 0, len(courseMap))
	for _, c := range courseMap {
		courses = append(courses, c)
	}

	// Count total
	countQuery := `SELECT COUNT(*) FROM courses`
	countRows, err := r.store.ExecQuery(context.Background(), countQuery)
	if err != nil || len(countRows) == 0 {
		return nil, 0, err
	}
	total := int(countRows[0]["count"].(int64))

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

func toTimePtr(val interface{}) *time.Time {
	if val == nil {
		return nil
	}
	t := val.(time.Time)
	return &t
}
