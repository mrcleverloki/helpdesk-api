package repository

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	"helpdesk-api/app/model"
)

var (
	ErrNotFound  = errors.New("data tidak ditemukan")
	ErrDuplicate = errors.New("data duplikat")
)

type UserRepository interface {
	FindAll(ctx context.Context, q model.ListQuery) ([]model.User, int, error)
	FindByID(ctx context.Context, id int) (model.User, error)
	FindByUsername(ctx context.Context, username string) (model.User, error)
	Create(ctx context.Context, u model.User) (model.User, error)
	Update(ctx context.Context, u model.User) (model.User, error)
	Delete(ctx context.Context, id int) error
}

type userPostgresRepository struct {
	pool *pgxpool.Pool
}

func NewUserRepository(pool *pgxpool.Pool) UserRepository {
	return &userPostgresRepository{pool: pool}
}

func (r *userPostgresRepository) FindAll(ctx context.Context, q model.ListQuery) ([]model.User, int, error) {
	whereClauses := []string{"1=1"}
	args := []any{}
	argCount := 1

	if q.Search != "" {
		whereClauses = append(whereClauses, fmt.Sprintf("(username ILIKE $%d OR email ILIKE $%d)", argCount, argCount))
		args = append(args, "%"+q.Search+"%")
		argCount++
	}

	if q.IsActive != nil {
		whereClauses = append(whereClauses, fmt.Sprintf("is_active = $%d", argCount))
		args = append(args, *q.IsActive)
		argCount++
	}

	whereStr := strings.Join(whereClauses, " AND ")

	countQuery := fmt.Sprintf("SELECT COUNT(*) FROM users WHERE %s", whereStr)
	var total int
	if err := r.pool.QueryRow(ctx, countQuery, args...).Scan(&total); err != nil {
		return nil, 0, fmt.Errorf("hitung total user: %w", err)
	}

	allowedCols := map[string]string{
		"id":         "id",
		"username":   "username",
		"email":      "email",
		"created_at": "created_at",
	}
	sortCol := allowedCols[q.Sort]
	if sortCol == "" {
		sortCol = "id"
	}

	orderDir := "ASC"
	if strings.ToUpper(q.Order) == "DESC" {
		orderDir = "DESC"
	}

	offset := (q.Page - 1) * q.Limit
	dataQuery := fmt.Sprintf(
		"SELECT id, username, email, password, role, is_active, created_at FROM users WHERE %s ORDER BY %s %s LIMIT $%d OFFSET $%d",
		whereStr, sortCol, orderDir, argCount, argCount+1,
	)
	args = append(args, q.Limit, offset)

	rows, err := r.pool.Query(ctx, dataQuery, args...)
	if err != nil {
		return nil, 0, fmt.Errorf("ambil daftar user: %w", err)
	}
	defer rows.Close()

	users := []model.User{}
	for rows.Next() {
		var u model.User
		if err := rows.Scan(&u.ID, &u.Username, &u.Email, &u.Password, &u.Role, &u.IsActive, &u.CreatedAt); err != nil {
			return nil, 0, fmt.Errorf("scan user: %w", err)
		}
		users = append(users, u)
	}

	return users, total, nil
}

func (r *userPostgresRepository) FindByID(ctx context.Context, id int) (model.User, error) {
	var u model.User
	err := r.pool.QueryRow(ctx,
		`SELECT id, username, email, password, role, is_active, created_at FROM users WHERE id = $1`, id,
	).Scan(&u.ID, &u.Username, &u.Email, &u.Password, &u.Role, &u.IsActive, &u.CreatedAt)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return model.User{}, ErrNotFound
		}
		return model.User{}, fmt.Errorf("ambil user berdasarkan id: %w", err)
	}
	return u, nil
}

func (r *userPostgresRepository) FindByUsername(ctx context.Context, username string) (model.User, error) {
	var u model.User
	err := r.pool.QueryRow(ctx,
		`SELECT id, username, email, password, role, is_active, created_at FROM users WHERE LOWER(username) = LOWER($1)`, username,
	).Scan(&u.ID, &u.Username, &u.Email, &u.Password, &u.Role, &u.IsActive, &u.CreatedAt)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return model.User{}, ErrNotFound
		}
		return model.User{}, fmt.Errorf("ambil user berdasarkan username: %w", err)
	}
	return u, nil
}

func (r *userPostgresRepository) Create(ctx context.Context, u model.User) (model.User, error) {
	err := r.pool.QueryRow(ctx,
		`INSERT INTO users (username, email, password, role, is_active) VALUES ($1, $2, $3, $4, $5) RETURNING id, created_at`,
		u.Username, u.Email, u.Password, u.Role, u.IsActive,
	).Scan(&u.ID, &u.CreatedAt)

	if err != nil {
		if strings.Contains(err.Error(), "users_username_key") || strings.Contains(err.Error(), "duplicate key") {
			return model.User{}, ErrDuplicate
		}
		return model.User{}, fmt.Errorf("simpan user: %w", err)
	}
	return u, nil
}

func (r *userPostgresRepository) Update(ctx context.Context, u model.User) (model.User, error) {
	err := r.pool.QueryRow(ctx,
		`UPDATE users SET username = $1, email = $2, is_active = $3 WHERE id = $4 RETURNING created_at, role`,
		u.Username, u.Email, u.IsActive, u.ID,
	).Scan(&u.CreatedAt, &u.Role)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return model.User{}, ErrNotFound
		}
		if strings.Contains(err.Error(), "users_username_key") || strings.Contains(err.Error(), "duplicate key") {
			return model.User{}, ErrDuplicate
		}
		return model.User{}, fmt.Errorf("perbarui user: %w", err)
	}
	return u, nil
}

func (r *userPostgresRepository) Delete(ctx context.Context, id int) error {
	tag, err := r.pool.Exec(ctx, `DELETE FROM users WHERE id = $1`, id)
	if err != nil {
		return fmt.Errorf("hapus user: %w", err)
	}
	if tag.RowsAffected() == 0 {
		return ErrNotFound
	}
	return nil
}
