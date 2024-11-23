package postgres

import (
	"context"
	"telegram-service/internal/entity"

	"github.com/jackc/pgx/v5"
)

type UserRepository struct {
	conn *pgx.Conn
}

func NewUserRepository(conn *pgx.Conn) *UserRepository {
	return &UserRepository{conn: conn}
}

func (r *UserRepository) CreateUser(ctx context.Context, user *entity.User) error {
	query := `INSERT INTO users (telegram_id, is_auto_update, auto_update_interval)
	          VALUES ($1, $2, $3)`

	_, err := r.conn.Exec(ctx, query, user.TelegramID, user.IsAutoUpdate, user.AutoUpdateInterval)
	return err
}

func (r *UserRepository) GetUserByTelegramID(ctx context.Context, telegramID int) (*entity.User, error) {
	query := `SELECT telegram_id, is_auto_update, auto_update_interval
	          FROM users WHERE telegram_id = $1`

	var user entity.User
	err := r.conn.QueryRow(ctx, query, telegramID).Scan(&user.TelegramID, &user.IsAutoUpdate, &user.AutoUpdateInterval)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &user, err
}

func (r *UserRepository) UpdateUser(ctx context.Context, user *entity.User) error {
	query := `UPDATE users
	          SET is_auto_update = $1, auto_update_interval = $2
	          WHERE telegram_id = $3`

	_, err := r.conn.Exec(ctx, query, user.IsAutoUpdate, user.AutoUpdateInterval, user.TelegramID)
	return err
}
