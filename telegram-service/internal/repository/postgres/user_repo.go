package postgres

import "telegram-service/internal/entity"

type PostgresUserRepository struct{}

func (r *PostgresUserRepository) CreateUser(user entity.User) error {
	// Заглушка
	return nil
}

func (r *PostgresUserRepository) GetUserByTelegramID(telegramID int64) (*entity.User, error) {
	// Заглушка
	return &entity.User{}, nil
}
