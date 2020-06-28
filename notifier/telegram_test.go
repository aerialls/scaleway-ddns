package notifier

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTelegramGoodTemplateSimple(t *testing.T) {
	tg := &Telegram{template: "{{ .RecordName }}"}
	data := &TelegramMessageData{
		RecordName: "foobar",
	}
	template, err := tg.getMessageFromTemplate(data)

	assert.Nil(t, err)
	assert.Equal(t, template, "foobar")
}

func TestTelegramGoodTemplateFull(t *testing.T) {
	tg := &Telegram{
		template: "{{ .RecordName }}{{ .RecordType }}{{ .Domain }}{{ .PreviousIP }}{{ .NewIP }}",
	}
	data := &TelegramMessageData{
		RecordName: "a",
		RecordType: "b",
		Domain:     "c",
		PreviousIP: "d",
		NewIP:      "e",
	}
	template, err := tg.getMessageFromTemplate(data)

	assert.Nil(t, err)
	assert.Equal(t, template, "abcde")
}

func TestTelegramWrongTemplate(t *testing.T) {
	tg := &Telegram{template: "{{ .Unknown }}"}
	data := &TelegramMessageData{
		RecordName: "foobar",
	}
	template, err := tg.getMessageFromTemplate(data)

	assert.NotNil(t, err)
	assert.Equal(t, template, "")
}
