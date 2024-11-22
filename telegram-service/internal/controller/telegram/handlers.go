package telegram

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type TelegramHandler struct {
	Bot *tgbotapi.BotAPI
}

func (h *TelegramHandler) HandleCommand(update tgbotapi.Update) {
	message := update.Message
	fmt.Println(message.Command())
	switch message.Command() {
	case "start":
		h.HandleStart(message)
	case "rates":
		h.HandleRates(message)
	case "start_auto":
		h.HandleStartAuto(message)
	case "stop_auto":
		h.HandleStopAuto(message)
	default:
		h.sendMessage(message.Chat.ID, "Команда не распознана.")
	}
}

func (h *TelegramHandler) HandleStart(message *tgbotapi.Message) {
	// Добавляем пользователя в базу данных

	h.sendMessage(message.Chat.ID, "Добро пожаловать! Доступные команды: /rates, /start-auto, /stop-auto.")
}

func (h *TelegramHandler) HandleRates(message *tgbotapi.Message) {
	args := strings.Fields(message.CommandArguments())

	if len(args) == 0 {
		// Получаем все курсы

		h.sendMessage(message.Chat.ID, "Текущие курсы:\nBTC: 100 000$\nETH: 4 000$\nSOL: 50$")
		return
	}

	crypto := args[0]

	// Получаем курс

	h.sendMessage(message.Chat.ID, fmt.Sprintf("Текущий курс для %s: 50 000$", crypto))
}

func (h *TelegramHandler) HandleStartAuto(message *tgbotapi.Message) {
	args := strings.Fields(message.CommandArguments())
	if len(args) == 0 {
		h.sendMessage(message.Chat.ID, "Укажите интервал обновлений в минутах.")
		return
	}
	interval, err := strconv.Atoi(args[0])
	if err != nil || interval <= 0 {
		h.sendMessage(message.Chat.ID, "Некорректный интервал. Укажите число больше 0.")
		return
	}

	// Запускаем автообновление

	h.sendMessage(message.Chat.ID, fmt.Sprintf("Автообновление установлено каждые %d минут.", interval))
}

func (h *TelegramHandler) HandleStopAuto(message *tgbotapi.Message) {
	// Останавливаем автообновление

	h.sendMessage(message.Chat.ID, "Автообновление отключено.")
}

func (h *TelegramHandler) sendMessage(chatID int64, text string) {
	msg := tgbotapi.NewMessage(chatID, text)
	_, err := h.Bot.Send(msg)
	if err != nil {
		log.Printf("Ошибка отправки сообщения: %v", err)
	}
}
