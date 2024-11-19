package postgres

import "telegram-service/internal/entity"

type PostgresUserCryptoRepository struct{}

func (r *PostgresUserCryptoRepository) AddUserCrypto(userID, cryptoID int) error {
	// Заглушка для добавления связи между пользователем и криптовалютой
	return nil
}

func (r *PostgresUserCryptoRepository) GetCryptosByUserID(userID int) ([]entity.Crypto, error) {
	// Заглушка для получения всех криптовалют, привязанных к пользователю
	return []entity.Crypto{}, nil
}

func (r *PostgresUserCryptoRepository) RemoveUserCrypto(userID, cryptoID int) error {
	// Заглушка для удаления связи между пользователем и криптовалютой
	return nil
}
