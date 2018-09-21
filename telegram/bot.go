package telegram

import (
	"log"
	"strconv"
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

	home := tb.ReplyButton{Text: "🏠 Home"}

	arena := tb.ReplyButton{Text: "⚔ War"}
	challenge := tb.ReplyButton{Text: "🤺 Challenge"}
	battle := tb.ReplyButton{Text: "🗺 Battle"}

	headAttack := tb.ReplyButton{Text: "⚔ Head"}
	bodyAttack := tb.ReplyButton{Text: "⚔ Body"}
	feetAttack := tb.ReplyButton{Text: "⚔ Feet"}

	attackCommands := [][]tb.ReplyButton{
		[]tb.ReplyButton{headAttack},
		[]tb.ReplyButton{bodyAttack},
		[]tb.ReplyButton{feetAttack},
	}

	headDefense := tb.ReplyButton{Text: "⛨ Head"}
	bodyDefense := tb.ReplyButton{Text: "⛨ Body"}
	feetDefense := tb.ReplyButton{Text: "⛨ Feet"}

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
		[]tb.ReplyButton{home},
	}

	commands := [][]tb.ReplyButton{
		[]tb.ReplyButton{arena},
	}

	b.Handle(&challenge, func(m *tb.Message) {
		f := fighter.NewFighterFormTelegramUser(m.Sender)
		d := duel.NewDuel(f)

		b.Send(m.Sender, "Your opponent has to forward this message to the @twarsbot:")
		// TODO: username is not available all time add a fallback first name
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
		u, err := uuid.FromString(m.Payload)
		if err != nil {
			b.Send(m.Sender, "Invalid duel identification!")
		}
		// TODO: Add not found error return value here!
		d := duel.GetDuelByID(u)
		d.SetChallenged(f)
		b.Send(m.Sender, "🤺 The duel has started!")
		b.Send(m.Sender, "Choose point of attack and point of block.", &tb.ReplyMarkup{
			ReplyKeyboard: attackDefenseCommands,
		})

		// NOTE: Comunicate about challenger accepted to the challenger
		// TODO: username is not available all time add a fallback first name
		b.Send(d.GetChallenger(), f.GetUsername()+" accepted your challenge!")
		b.Send(d.GetChallenger(), "🤺 The duel has started!")
		b.Send(d.GetChallenger(), "Choose point of attack and point of block.", &tb.ReplyMarkup{
			ReplyKeyboard: attackDefenseCommands,
		})
	})

	b.Handle(&headAttack, func(m *tb.Message) {
		f := fighter.NewFighterFormTelegramUser(m.Sender)
		f.SetAttack("head")
		if !f.HasDefense() {
			b.Send(m.Sender, "Great! Now select your defence!", &tb.ReplyMarkup{
				ReplyKeyboard: defenseCommands,
			})
		}
		if f.Ready() {
			b.Send(m.Sender, "Good plan! let's see what happens next.")
			// TODO: Do not ignore this error!
			d, _ := duel.GetDuelByFighter(f)
			if d.Ready() {
				result := d.Duel()
				if d.Finished() {
					w := d.Winner()
					b.Send(d.GetChallenger(), w.GetUsername()+"(❤ "+strconv.Itoa(w.Health())+") won the duel!", &tb.ReplyMarkup{
						ReplyKeyboard: commands,
					})
					b.Send(d.GetChallenged(), w.GetUsername()+"(❤ "+strconv.Itoa(w.Health())+") won the duel!", &tb.ReplyMarkup{
						ReplyKeyboard: commands,
					})
				}
				b.Send(d.GetChallenger(), result.Challenged.Result+" "+result.Challenger.Result, &tb.ReplyMarkup{
					ReplyKeyboard: attackDefenseCommands,
				})
				b.Send(d.GetChallenged(), result.Challenged.Result+" "+result.Challenger.Result, &tb.ReplyMarkup{
					ReplyKeyboard: attackDefenseCommands,
				})
			}
		}
	})

	b.Handle(&bodyAttack, func(m *tb.Message) {
		f := fighter.NewFighterFormTelegramUser(m.Sender)
		f.SetAttack("body")

		d, _ := duel.GetDuelByFighter(f)

		if !f.HasDefense() {
			b.Send(m.Sender, "Great! Now select your defence!", &tb.ReplyMarkup{
				ReplyKeyboard: defenseCommands,
			})
		}
		if d.Finished() {
			w := d.Winner()
			b.Send(d.GetChallenger(), w.GetUsername()+"(❤ "+strconv.Itoa(w.Health())+") won the duel!", &tb.ReplyMarkup{
				ReplyKeyboard: commands,
			})
			b.Send(d.GetChallenged(), w.GetUsername()+"(❤ "+strconv.Itoa(w.Health())+") won the duel!", &tb.ReplyMarkup{
				ReplyKeyboard: commands,
			})
		}
		if f.Ready() {
			b.Send(m.Sender, "Good plan! let's see what happens next.")
			// TODO: Do not ignore this error!

			if d.Ready() {
				result := d.Duel()

				b.Send(d.GetChallenger(), result.Challenged.Result+" "+result.Challenger.Result, &tb.ReplyMarkup{
					ReplyKeyboard: attackDefenseCommands,
				})
				b.Send(d.GetChallenged(), result.Challenged.Result+" "+result.Challenger.Result, &tb.ReplyMarkup{
					ReplyKeyboard: attackDefenseCommands,
				})
			}
		}
	})

	b.Handle(&feetAttack, func(m *tb.Message) {
		f := fighter.NewFighterFormTelegramUser(m.Sender)
		f.SetAttack("feet")
		if !f.HasDefense() {
			b.Send(m.Sender, "Great! Now select your defence!", &tb.ReplyMarkup{
				ReplyKeyboard: defenseCommands,
			})
		}
		if f.Ready() {
			b.Send(m.Sender, "Good plan! let's see what happens next.")
			// TODO: Do not ignore this error!
			d, _ := duel.GetDuelByFighter(f)
			if d.Ready() {
				result := d.Duel()
				if d.Finished() {
					w := d.Winner()
					b.Send(d.GetChallenger(), w.GetUsername()+"(❤ "+strconv.Itoa(w.Health())+") won the duel!", &tb.ReplyMarkup{
						ReplyKeyboard: commands,
					})
					b.Send(d.GetChallenged(), w.GetUsername()+"(❤ "+strconv.Itoa(w.Health())+") won the duel!", &tb.ReplyMarkup{
						ReplyKeyboard: commands,
					})
				}
				b.Send(d.GetChallenger(), result.Challenged.Result+" "+result.Challenger.Result, &tb.ReplyMarkup{
					ReplyKeyboard: attackDefenseCommands,
				})
				b.Send(d.GetChallenged(), result.Challenged.Result+" "+result.Challenger.Result, &tb.ReplyMarkup{
					ReplyKeyboard: attackDefenseCommands,
				})
			}
		}
	})

	b.Handle(&headDefense, func(m *tb.Message) {
		f := fighter.NewFighterFormTelegramUser(m.Sender)
		f.SetDefense("head")
		if !f.HasAttack() {
			b.Send(m.Sender, "Great! Now select your attack!", &tb.ReplyMarkup{
				ReplyKeyboard: attackCommands,
			})
		}
		if f.Ready() {
			b.Send(m.Sender, "Good plan! let's see what happens next.")
			// TODO: Do not ignore this error!
			d, _ := duel.GetDuelByFighter(f)
			if d.Ready() {
				result := d.Duel()
				if d.Finished() {
					w := d.Winner()
					b.Send(d.GetChallenger(), w.GetUsername()+"(❤ "+strconv.Itoa(w.Health())+") won the duel!", &tb.ReplyMarkup{
						ReplyKeyboard: commands,
					})
					b.Send(d.GetChallenged(), w.GetUsername()+"(❤ "+strconv.Itoa(w.Health())+") won the duel!", &tb.ReplyMarkup{
						ReplyKeyboard: commands,
					})
				}
				b.Send(d.GetChallenger(), result.Challenged.Result+" "+result.Challenger.Result, &tb.ReplyMarkup{
					ReplyKeyboard: attackDefenseCommands,
				})
				b.Send(d.GetChallenged(), result.Challenged.Result+" "+result.Challenger.Result, &tb.ReplyMarkup{
					ReplyKeyboard: attackDefenseCommands,
				})
			}
		}
	})

	b.Handle(&bodyDefense, func(m *tb.Message) {
		f := fighter.NewFighterFormTelegramUser(m.Sender)
		f.SetDefense("body")
		if !f.HasAttack() {
			b.Send(m.Sender, "Great! Now select your attack!", &tb.ReplyMarkup{
				ReplyKeyboard: attackCommands,
			})
		}
		if f.Ready() {
			b.Send(m.Sender, "Good plan! let's see what happens next.")
			// TODO: Do not ignore this error!
			d, _ := duel.GetDuelByFighter(f)
			if d.Ready() {
				result := d.Duel()
				if d.Finished() {
					w := d.Winner()
					b.Send(d.GetChallenger(), w.GetUsername()+"(❤ "+strconv.Itoa(w.Health())+") won the duel!", &tb.ReplyMarkup{
						ReplyKeyboard: commands,
					})
					b.Send(d.GetChallenged(), w.GetUsername()+"(❤ "+strconv.Itoa(w.Health())+") won the duel!", &tb.ReplyMarkup{
						ReplyKeyboard: commands,
					})
				}
				b.Send(d.GetChallenger(), result.Challenged.Result+" "+result.Challenger.Result, &tb.ReplyMarkup{
					ReplyKeyboard: attackDefenseCommands,
				})
				b.Send(d.GetChallenged(), result.Challenged.Result+" "+result.Challenger.Result, &tb.ReplyMarkup{
					ReplyKeyboard: attackDefenseCommands,
				})
			}
		}
	})

	b.Handle(&feetDefense, func(m *tb.Message) {
		f := fighter.NewFighterFormTelegramUser(m.Sender)
		f.SetDefense("feet")
		if !f.HasAttack() {
			b.Send(m.Sender, "Great! Now select your attack!", &tb.ReplyMarkup{
				ReplyKeyboard: attackCommands,
			})
		}
		if f.Ready() {
			b.Send(m.Sender, "Good plan! let's see what happens next.")
			// TODO: Do not ignore this error!
			d, _ := duel.GetDuelByFighter(f)
			if d.Ready() {
				result := d.Duel()
				if d.Finished() {
					w := d.Winner()
					b.Send(d.GetChallenger(), w.GetUsername()+"(❤ "+strconv.Itoa(w.Health())+") won the duel!", &tb.ReplyMarkup{
						ReplyKeyboard: commands,
					})
					b.Send(d.GetChallenged(), w.GetUsername()+"(❤ "+strconv.Itoa(w.Health())+") won the duel!", &tb.ReplyMarkup{
						ReplyKeyboard: commands,
					})
				}
				b.Send(d.GetChallenger(), result.Challenged.Result+" "+result.Challenger.Result, &tb.ReplyMarkup{
					ReplyKeyboard: attackDefenseCommands,
				})
				b.Send(d.GetChallenged(), result.Challenged.Result+" "+result.Challenger.Result, &tb.ReplyMarkup{
					ReplyKeyboard: attackDefenseCommands,
				})
			}
		}
	})

	b.Handle(&home, func(m *tb.Message) {
		b.Send(m.Sender, "🔙", &tb.ReplyMarkup{
			ReplyKeyboard: commands,
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
