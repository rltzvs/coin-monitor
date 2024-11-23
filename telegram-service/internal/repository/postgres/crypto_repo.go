package postgres

import (
	"context"
	"telegram-service/internal/entity"

	"github.com/jackc/pgx/v5"
)

type CryptoRepository struct {
	conn *pgx.Conn
}

func NewCryptoRepository(conn *pgx.Conn) *CryptoRepository {
	return &CryptoRepository{conn: conn}
}

func (r *CryptoRepository) CreateCrypto(ctx context.Context, crypto *entity.Crypto) error {
	query := `INSERT INTO cryptos (symbol) VALUES ($1)`

	_, err := r.conn.Exec(ctx, query, crypto.Symbol)
	return err
}
