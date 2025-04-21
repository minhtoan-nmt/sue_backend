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
		ID:         row["id"].(int),
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

func toStrPtr(val interface{}) *string {
	if val == nil {
		return nil
	}
	s := val.(string)
	return &s
}
