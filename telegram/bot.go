package telegram

import (
	"log"
	"time"

	tb "gopkg.in/tucnak/telebot.v2"
)

// NewBot telegram bot
func NewBot(token string) *tb.Bot {
	b, err := tb.NewBot(tb.Settings{
		Token:  token,
		Poller: &tb.LongPoller{Timeout: 10 * time.Second},
	})

	if err != nil {
		log.Fatal(err)
	}

	arena := tb.ReplyButton{Text: "🗡  Arena"}
	challenge := tb.ReplyButton{Text: "🤺 Challenge"}
	battle := tb.ReplyButton{Text: "🗺    Battle"}

	arenaCommands := [][]tb.ReplyButton{
		[]tb.ReplyButton{challenge},
		[]tb.ReplyButton{battle},
	}

	commands := [][]tb.ReplyButton{
		[]tb.ReplyButton{arena},
	}

	b.Handle(&challenge, func(m *tb.Message) {
		b.Send(m.Sender, "You challenges one fighter")
	})

	b.Handle(&battle, func(m *tb.Message) {
		b.Send(m.Sender, "You challenges all fighters")
	})

	b.Handle(&arena, func(m *tb.Message) {
		b.Send(m.Sender, "Choose challenge one fighter or battle all fighters", &tb.ReplyMarkup{
			ReplyKeyboard: arenaCommands,
		})
	})

	b.Handle("/start", func(m *tb.Message) {
		if !m.Private() {
			return
		}

		b.Send(m.Sender, "Welcome warrior!", &tb.ReplyMarkup{
			ReplyKeyboard: commands,
		})
	})

	return b
}
