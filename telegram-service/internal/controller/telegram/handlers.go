package telegram

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type CommandHandler interface {
	HandleCommand(message *tgbotapi.Message)
}

type TestCommandHandler struct{}

func (h *TestCommandHandler) HandleCommand(message *tgbotapi.Message) {
	switch message.Command() {
	case "start":
		h.handleStart(message)
	case "help":
		h.handleHelp(message)
	default:
		h.handleUnknown(message)
	}
}

func (h *TestCommandHandler) handleStart(message *tgbotapi.Message) {
	log.Printf("Received /start from user %d", message.From.ID)
}

func (h *TestCommandHandler) handleHelp(message *tgbotapi.Message) {
	log.Printf("Received /help from user %d", message.From.ID)
}

func (h *TestCommandHandler) handleUnknown(message *tgbotapi.Message) {
	log.Printf("Received unknown command from user %d: %s", message.From.ID, message.Text)
}
