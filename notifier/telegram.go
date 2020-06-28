package notifier

import (
	"bytes"
	templating "text/template"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

// Telegram struct to be able to notify with Telegram messages
type Telegram struct {
	token    string
	chatID   int64
	template string
}

// TelegramMessageData holds information for the Telegram template message
type TelegramMessageData struct {
	Domain     string
	RecordName string
	RecordType string
	PreviousIP string
	NewIP      string
}

// NewTelegram returns a new Telegram notifier
func NewTelegram(
	token string,
	chatID int64,
	template string,
) *Telegram {
	return &Telegram{
		token:    token,
		chatID:   chatID,
		template: template,
	}
}

// Notify launches a new message on Telegram when the IP has changed
func (t *Telegram) Notify(
	domain string,
	recordName string,
	recordType string,
	previousIP string,
	newIP string,
) error {
	bot, err := tgbotapi.NewBotAPI(t.token)
	if err != nil {
		return err
	}

	if previousIP == "" {
		previousIP = "(empty)"
	}

	message, err := t.getMessageFromTemplate(&TelegramMessageData{
		Domain:     domain,
		RecordName: recordName,
		RecordType: recordType,
		PreviousIP: previousIP,
		NewIP:      newIP,
	})

	if err != nil {
		return err
	}

	msg := tgbotapi.NewMessage(t.chatID, message)
	msg.ParseMode = "markdown"

	_, err = bot.Send(msg)
	if err != nil {
		return err
	}

	return nil
}

func (t *Telegram) getMessageFromTemplate(data *TelegramMessageData) (string, error) {
	template, err := templating.New("telegram").Parse(t.template)
	if err != nil {
		return "", err
	}

	var message bytes.Buffer
	err = template.Execute(&message, data)

	if err != nil {
		return "", err
	}

	return message.String(), nil
}
