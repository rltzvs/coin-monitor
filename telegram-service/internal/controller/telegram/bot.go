package telegram

import "log/slog"

type TelegramBot struct {
}

func NewTelegramBot(token string, logger *slog.Logger) *TelegramBot {
	return &TelegramBot{}
}

func (b *TelegramBot) Start() error {
	return nil
}
