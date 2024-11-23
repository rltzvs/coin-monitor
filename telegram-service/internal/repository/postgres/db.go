package postgres

import (
	"context"
	"fmt"
	"telegram-service/internal/config"

	"github.com/jackc/pgx/v5"
)

type DB struct {
	conn *pgx.Conn
}

func NewDBConnection(cfg *config.DatabaseConfig) (*DB, error) {
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.Dbname)

	conn, err := pgx.Connect(context.Background(), connStr)
	if err != nil {
		return nil, fmt.Errorf("не удалось подключиться к базе данных: %v", err)
	}

	if err := conn.Ping(context.Background()); err != nil {
		return nil, fmt.Errorf("не удалось отправить сигнал подключения: %v", err)
	}

	return &DB{conn: conn}, nil
}

func (db *DB) Close() {
	db.conn.Close(context.Background())
}
