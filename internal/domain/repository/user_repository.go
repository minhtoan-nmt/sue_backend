package repository

import (
	"context"
	"sue_backend/internal/domain/model"
	"sue_backend/internal/infra/cache"
	"sue_backend/internal/infra/db"
	"time"
)

type UserRepository struct {
	store db.Store
	cache cache.Store
}

func NewUserRepository(store db.Store, cache cache.Store) *UserRepository {
	return &UserRepository{
		store: store,
		cache: cache,
	}
}
func (r *UserRepository) FindByEmail(email string) (*model.User, error) {
	query := `SELECT id, first_name, last_name, email, password, salt, role, phone, status, created_at, updated_at FROM users WHERE email = $1`
	ctx := context.Background()
	rows, err := r.store.ExecQuery(ctx, query, email)
	if err != nil {
		return nil, err
	}
	if len(rows) == 0 {
		return nil, nil
	}
	row := rows[0]

	u := &model.User{
		ID:         row["id"].(int64),
		First_name: toStrPtr(row["first_name"]),
		Last_name:  toStrPtr(row["last_name"]),
		Email:      toStrPtr(row["email"]),
		Password:   toStrPtr(row["password"]),
		Salt:       toStrPtr(row["salt"]),
		Role:       model.UserRole(row["role"].(string)),
		Phone:      toStrPtr(row["phone"]),
		Status:     row["status"].(string),
		CreatedAt:  row["created_at"].(time.Time),
		UpdatedAt:  row["updated_at"].(time.Time),
	}
	return u, nil
}

func (r *UserRepository) Create(user *model.User) error {
	query := `INSERT INTO users (first_name, last_name, email, password, salt, phone, role, status, created_at, updated_at)
				VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10)`
	ctx := context.Background()
	return r.store.Exec(ctx, query,
		user.First_name,
		user.Last_name,
		user.Email,
		user.Password,
		user.Salt,
		user.Phone,
		user.Role,
		user.Status,
		user.CreatedAt,
		user.UpdatedAt,
	)
}
func (r *UserRepository) FindByID(id int64) (*model.User, error) {
	query := `SELECT id, first_name, last_name, email, password, salt, role, phone, status, created_at, updated_at FROM users WHERE id = $1`
	ctx := context.Background()
	rows, err := r.store.ExecQuery(ctx, query, id)
	if err != nil || len(rows) == 0 {
		return nil, err
	}
	return mapRowToUser(rows[0]), nil
}

func (r *UserRepository) FindPaginated(page, limit int) ([]*model.User, int, error) {
	offset := (page - 1) * limit
	ctx := context.Background()

	// 1. Query paginated data
	query := `SELECT id, first_name, last_name, email, role, phone, status, created_at, updated_at
	          FROM users ORDER BY id LIMIT $1 OFFSET $2`
	rows, err := r.store.ExecQuery(ctx, query, limit, offset)
	if err != nil {
		return nil, 0, err
	}

	users := make([]*model.User, 0, len(rows))
	for _, row := range rows {
		users = append(users, &model.User{
			ID:         row["id"].(int64),
			First_name: toStrPtr(row["first_name"]),
			Last_name:  toStrPtr(row["last_name"]),
			Email:      toStrPtr(row["email"]),
			Role:       model.UserRole(row["role"].(string)),
			Phone:      toStrPtr(row["phone"]),
			Status:     row["status"].(string),
			CreatedAt:  row["created_at"].(time.Time),
			UpdatedAt:  row["updated_at"].(time.Time),
		})
	}

	// 2. Count total
	countQuery := `SELECT COUNT(*) as total FROM users`
	countRows, err := r.store.ExecQuery(ctx, countQuery)
	if err != nil || len(countRows) == 0 {
		return nil, 0, err
	}
	total := int(countRows[0]["total"].(int64))

	return users, total, nil
}

func (r *UserRepository) FindAll() ([]*model.User, error) {
	query := `SELECT id, first_name, last_name, email, role, phone, status, created_at, updated_at FROM users`
	ctx := context.Background()
	rows, err := r.store.ExecQuery(ctx, query)
	if err != nil {
		return nil, err
	}
	users := make([]*model.User, 0, len(rows))
	for _, row := range rows {
		users = append(users, mapRowToUser(row))
	}
	return users, nil
}
func (r *UserRepository) Update(user *model.User) error {
	query := `UPDATE users SET first_name=$1, last_name=$2, phone=$3, status=$4, updated_at=$5 WHERE id=$6`
	ctx := context.Background()
	return r.store.Exec(ctx, query,
		user.First_name, user.Last_name,
		user.Phone, user.Status, time.Now(), user.ID,
	)
}

func (r *UserRepository) Delete(id int64) error {
	query := `DELETE FROM users WHERE id = $1`
	ctx := context.Background()
	return r.store.Exec(ctx, query, id)
}

func mapRowToUser(row map[string]interface{}) *model.User {
	return &model.User{
		ID:         row["id"].(int64),
		First_name: toStrPtr(row["first_name"]),
		Last_name:  toStrPtr(row["last_name"]),
		Email:      toStrPtr(row["email"]),
		Password:   toStrPtr(row["password"]),
		Salt:       toStrPtr(row["salt"]),
		Role:       model.UserRole(row["role"].(string)),
		Phone:      toStrPtr(row["phone"]),
		Status:     row["status"].(string),
		CreatedAt:  row["created_at"].(time.Time),
		UpdatedAt:  row["updated_at"].(time.Time),
	}
}
func toStrPtr(val interface{}) *string {
	if val == nil {
		return nil
	}
	s := val.(string)
	return &s
}
