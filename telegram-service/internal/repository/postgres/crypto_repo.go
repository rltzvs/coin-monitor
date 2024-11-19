package postgres

import "telegram-service/internal/entity"

type PostgresCryptoRepository struct{}

func (r *PostgresCryptoRepository) CreateCrypto(crypto entity.Crypto) error {
	// Заглушка для вставки новой записи в таблицу crypto
	return nil
}

func (r *PostgresCryptoRepository) GetAllCryptos() ([]entity.Crypto, error) {
	// Заглушка для получения всех криптовалют
	return []entity.Crypto{}, nil
}

func (r *PostgresCryptoRepository) GetCryptoByID(id int) (*entity.Crypto, error) {
	// Заглушка для получения криптовалюты по ID
	return &entity.Crypto{}, nil
}
