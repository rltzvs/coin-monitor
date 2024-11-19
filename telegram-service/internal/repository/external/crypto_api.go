package external

import "telegram-service/internal/entity"

type CryptoRepository interface {
	GetAllCryptos() ([]entity.Crypto, error)
}

type MockCryptoRepository struct{}

func (m *MockCryptoRepository) GetAllCryptos() ([]entity.Crypto, error) {
	return []entity.Crypto{}, nil
}
