package telegram

import (
	"log"
	"time"

	"github.com/wiliamsouza/warbot/duel"
	"github.com/wiliamsouza/warbot/fighter"

	uuid "github.com/satori/go.uuid"
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

	arena := tb.ReplyButton{Text: "🗡 Arena"}
	challenge := tb.ReplyButton{Text: "🤺 Challenge"}
	battle := tb.ReplyButton{Text: "🗺 Battle"}

	headAttack := tb.ReplyButton{Text: "🗡 Head"}
	bodyAttack := tb.ReplyButton{Text: "🗡 Body"}
	feetAttack := tb.ReplyButton{Text: "🗡 Feet"}

	attackCommands := [][]tb.ReplyButton{
		[]tb.ReplyButton{headAttack},
		[]tb.ReplyButton{bodyAttack},
		[]tb.ReplyButton{feetAttack},
	}

	headDefense := tb.ReplyButton{Text: "🛡 Head"}
	bodyDefense := tb.ReplyButton{Text: "🛡 Body"}
	feetDefense := tb.ReplyButton{Text: "🛡 Feet"}

	defenseCommands := [][]tb.ReplyButton{
		[]tb.ReplyButton{headDefense},
		[]tb.ReplyButton{bodyDefense},
		[]tb.ReplyButton{feetDefense},
	}

	attackDefenseCommands := [][]tb.ReplyButton{
		[]tb.ReplyButton{headAttack},
		[]tb.ReplyButton{bodyAttack},
		[]tb.ReplyButton{feetAttack},
		[]tb.ReplyButton{headDefense},
		[]tb.ReplyButton{bodyDefense},
		[]tb.ReplyButton{feetDefense},
	}

	arenaCommands := [][]tb.ReplyButton{
		[]tb.ReplyButton{challenge},
		[]tb.ReplyButton{battle},
	}

	commands := [][]tb.ReplyButton{
		[]tb.ReplyButton{arena},
	}

	b.Handle(&challenge, func(m *tb.Message) {
		f := fighter.NewFighterFormTelegramUser(m.Sender)
		f.SetType("challenger")
		d := duel.NewDuel(f)

		b.Send(m.Sender, "Your opponent has to forward this message to the @xwarsbot:")
		b.Send(m.Sender, f.GetUsername()+" challenges you to a duel!")
		b.Send(m.Sender, "/duel "+d.ID.String())
	})

	b.Handle(&battle, func(m *tb.Message) {
		b.Send(m.Sender, "You challenges all fighters")
	})

	b.Handle(&arena, func(m *tb.Message) {
		b.Send(m.Sender, "Choose challenge one fighter or battle all fighters", &tb.ReplyMarkup{
			ReplyKeyboard: arenaCommands,
		})
	})

	b.Handle("/duel", func(m *tb.Message) {
		f := fighter.NewFighterFormTelegramUser(m.Sender)
		f.SetType("challenged")
		u, err := uuid.FromString(m.Payload)
		if err != nil {
			b.Send(m.Sender, "Invalid duel identification!")
		}
		d := duel.GetDuelID(u)
		d.SetOpponent(f)
		b.Send(m.Sender, "🤺 The duel has started!")
		b.Send(m.Sender, "Choose point of attack and point of block.", &tb.ReplyMarkup{
			ReplyKeyboard: attackDefenseCommands,
		})

		b.Send(d.GetFighter(), f.GetUsername()+" accepted your challenge!")
		b.Send(d.GetFighter(), "🤺 The duel has started!")
		b.Send(d.GetFighter(), "Choose point of attack and point of block.", &tb.ReplyMarkup{
			ReplyKeyboard: attackDefenseCommands,
		})
	})

	b.Handle(&headAttack, func(m *tb.Message) {
		f := fighter.NewFighterFormTelegramUser(m.Sender)
		a := fighter.NewAttack("head")
		f.SetAttack(a)
		b.Send(m.Sender, "Great! Now select your defence!", &tb.ReplyMarkup{
			ReplyKeyboard: defenseCommands,
		})
	})

	b.Handle(&bodyAttack, func(m *tb.Message) {
		f := fighter.NewFighterFormTelegramUser(m.Sender)
		a := fighter.NewAttack("body")
		f.SetAttack(a)
		b.Send(m.Sender, "Great! Now select your defence!", &tb.ReplyMarkup{
			ReplyKeyboard: defenseCommands,
		})
	})

	b.Handle(&feetAttack, func(m *tb.Message) {
		f := fighter.NewFighterFormTelegramUser(m.Sender)
		a := fighter.NewAttack("feet")
		f.SetAttack(a)
		b.Send(m.Sender, "Great! Now select your defence!", &tb.ReplyMarkup{
			ReplyKeyboard: defenseCommands,
		})
	})

	b.Handle(&headDefense, func(m *tb.Message) {
		f := fighter.NewFighterFormTelegramUser(m.Sender)
		a := fighter.NewDefense("head")
		f.SetDefense(a)
		b.Send(m.Sender, "Great! Now select your attack!", &tb.ReplyMarkup{
			ReplyKeyboard: attackCommands,
		})
	})

	b.Handle(&bodyDefense, func(m *tb.Message) {
		f := fighter.NewFighterFormTelegramUser(m.Sender)
		a := fighter.NewDefense("body")
		f.SetDefense(a)
		b.Send(m.Sender, "Great! Now select your attack!", &tb.ReplyMarkup{
			ReplyKeyboard: attackCommands,
		})
	})

	b.Handle(&feetDefense, func(m *tb.Message) {
		f := fighter.NewFighterFormTelegramUser(m.Sender)
		a := fighter.NewDefense("feet")
		f.SetDefense(a)
		b.Send(m.Sender, "Great! Now select your attack!", &tb.ReplyMarkup{
			ReplyKeyboard: attackCommands,
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
