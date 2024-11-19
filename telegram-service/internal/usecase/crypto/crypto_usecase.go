package crypto

import "telegram-service/internal/entity"

type CryptoUsecaseImpl struct{}

func (u *CryptoUsecaseImpl) GetRates() ([]entity.Crypto, error) {
	// Заглушка
	return []entity.Crypto{}, nil
}
