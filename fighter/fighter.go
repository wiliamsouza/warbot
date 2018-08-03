package fighter

import (
	"strconv"
	"sync"

	tb "gopkg.in/tucnak/telebot.v2"
)

var (
	fightersMu sync.RWMutex
	fighters   = make(map[int]Fighter)
)

// Fighter interface
type Fighter interface {
	Fight(Fighter) *FightResult
	SetAttack(Attacker)
	SetDefense(Defender)
	Identification() int
	SetType(t string)
	GetUsername() string
	Recipient() string
}

// FightResult of the fight
type FightResult struct{}

// F implements Fighter interface
type F struct {
	ID int

	FirstName string
	LastName  string
	Username  string

	Type string

	Health  int
	Attack  Attacker
	Defense Defender
}

// Fight an opponent
func (f *F) Fight(opponent Fighter) *FightResult {
	return &FightResult{}
}

// SetAttack set the given attack
func (f *F) SetAttack(atack Attacker) {
	f.Attack = atack
}

// SetDefense set the given defense
func (f *F) SetDefense(defense Defender) {
	f.Defense = defense
}

// Identification return fighter identification
func (f *F) Identification() int {
	return f.ID
}

// SetType set duel type
func (f *F) SetType(t string) {
	f.Type = t
}

// GetUsername an opponent
func (f *F) GetUsername() string {
	return f.Username
}

// Recipient returns user ID (see telebot.v2.Recipient interface).
func (f *F) Recipient() string {
	return strconv.Itoa(f.ID)
}

// NewFighterFormTelegramUser create a fighter from telegram user
func NewFighterFormTelegramUser(u *tb.User) Fighter {
	if fighter, exist := fighters[u.ID]; exist {
		return fighter
	}
	f := &F{ID: u.ID, Username: u.Username, FirstName: u.FirstName, LastName: u.LastName, Health: 100}
	register(f.ID, f)
	return f
}

// register a fighter
func register(id int, fighter Fighter) {
	fightersMu.Lock()
	defer fightersMu.Unlock()
	if fighter == nil {
		// TODO: Change to return error
		panic("fighters: Register fighter is nil")
	}
	if _, dup := fighters[id]; dup {
		// TODO: Change to return error
		panic("fighter: Register called twice for fighter ")
	}
	fighters[id] = fighter
}
