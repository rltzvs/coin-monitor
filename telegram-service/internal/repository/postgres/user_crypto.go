package postgres

import (
	"context"
	"telegram-service/internal/entity"

	"github.com/jackc/pgx/v5"
)

type UserCryptoRepository struct {
	conn *pgx.Conn
}

func NewUserCryptoRepository(conn *pgx.Conn) *UserCryptoRepository {
	return &UserCryptoRepository{conn: conn}
}

func (r *UserCryptoRepository) AddUserCrypto(ctx context.Context, userCrypto *entity.UserCrypto) error {
	query := `INSERT INTO user_cryptos (user_id, crypto_id)
	          VALUES ($1, $2) ON CONFLICT DO NOTHING`

	_, err := r.conn.Exec(ctx, query, userCrypto.UserId, userCrypto.CryptoId)
	return err
}
func (r *UserCryptoRepository) RemoveUserCrypto(ctx context.Context, userID int, cryptoID int) error {
	query := `DELETE FROM user_cryptos WHERE user_id = $1 AND crypto_id = $2`

	_, err := r.conn.Exec(ctx, query, userID, cryptoID)
	return err
}

func (r *UserCryptoRepository) ListUserCryptos(ctx context.Context, userID int) ([]*entity.Crypto, error) {
	query := `
	SELECT c.id, c.symbol
	FROM user_cryptos uc
	INNER JOIN cryptos c ON uc.crypto_id = c.id
	WHERE uc.user_id = $1
	`

	rows, err := r.conn.Query(ctx, query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var cryptos []*entity.Crypto
	for rows.Next() {
		var crypto entity.Crypto
		err := rows.Scan(&crypto.ID, &crypto.Symbol)
		if err != nil {
			return nil, err
		}
		cryptos = append(cryptos, &crypto)
	}
	return cryptos, nil
}
