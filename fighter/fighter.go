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
	SetAttack(string)
	SetDefense(string)
	Defense() string
	Attack() string
	Identification() int
	GetUsername() string
	Recipient() string
	HasAttack() bool
	HasDefense() bool
	Ready() bool
	SetHealth(int)
	Health() int
	Reset()
	Dead() bool
}

// FightResult of the fight
type FightResult struct {
	Result string
}

// F implements Fighter interface
type F struct {
	ID int

	FirstName string
	LastName  string
	Username  string

	health  int
	attack  string
	defense string
}

// Fight an opponent
func (f *F) Fight(opponent Fighter) *FightResult {
	var result string

	if f.Attack() != opponent.Defense() {
		opponent.SetHealth(opponent.Health() - 20)
		result = f.GetUsername() + "(❤ " + strconv.Itoa(f.Health()) + ")" + " smashed his opponent and unexpectedly stabbed his opponent at the " + f.Attack() + "!"
	} else if f.Attack() == opponent.Defense() {
		result = f.GetUsername() + "(❤ " + strconv.Itoa(f.Health()) + ")" + " smashed his opponent at the " + f.Attack() + ", but " + opponent.GetUsername() + "(❤ " + strconv.Itoa(opponent.Health()) + ")" + " managed to block!"
	}

	f.Reset()
	return &FightResult{Result: result}
}

// Reset attack and defence
func (f *F) Reset() {
	f.SetAttack("")
	f.SetDefense("")
}

// SetAttack set the given attack
func (f *F) SetAttack(atack string) {
	f.attack = atack
}

// SetDefense set the given defense
func (f *F) SetDefense(defense string) {
	f.defense = defense
}

// Attack return attack
func (f *F) Attack() string {
	return f.attack
}

// Defense return defense
func (f *F) Defense() string {
	return f.defense
}

// HasAttack return true is attack is set
func (f *F) HasAttack() bool {
	return f.attack != ""
}

// HasDefense return true is attack is set
func (f *F) HasDefense() bool {
	return f.defense != ""
}

// Identification return fighter identification
func (f *F) Identification() int {
	return f.ID
}

// GetUsername an opponent
func (f *F) GetUsername() string {
	return f.Username
}

// Recipient returns user ID (see telebot.v2.Recipient interface).
func (f *F) Recipient() string {
	return strconv.Itoa(f.ID)
}

// Ready return true if fighter is ready to duel
func (f *F) Ready() bool {
	if f.HasAttack() && f.HasDefense() {
		return true
	}
	return false
}

// SetHealth set fighter health
func (f *F) SetHealth(h int) {
	f.health = h
}

// Health return fighter health
func (f *F) Health() int {
	return f.health
}

// Dead return true figther is dead
func (f *F) Dead() bool {
	return f.health == 0
}

// NewFighterFormTelegramUser create a fighter from telegram user
func NewFighterFormTelegramUser(u *tb.User) Fighter {
	if fighter, exist := fighters[u.ID]; exist {
		return fighter
	}
	f := &F{ID: u.ID, Username: u.Username, FirstName: u.FirstName, LastName: u.LastName, health: 100}
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
