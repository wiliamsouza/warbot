package duel

import (
	"errors"
	"sync"

	uuid "github.com/satori/go.uuid"
	"github.com/wiliamsouza/warbot/fighter"
)

var (
	duelsMu sync.RWMutex
	duels   = make(map[uuid.UUID]Dueler)
)

// Dueler interface
type Dueler interface {
	Duel(duelists ...fighter.Fighter) *Result
	SetOpponent(fighter fighter.Fighter)
	GetFighter() fighter.Fighter
	GetOpponent() fighter.Fighter
}

// Duel implements Dueler interface
type Duel struct {
	ID       uuid.UUID
	Fighter  fighter.Fighter
	Opponent fighter.Fighter
}

// Result of the duel
type Result struct{}

// Duel start the combat between two duelists
func (d *Duel) Duel(duelists ...fighter.Fighter) *Result {
	return &Result{}
}

// SetOpponent a fighter that accepted the challenge
func (d *Duel) SetOpponent(fighter fighter.Fighter) {
	d.Opponent = fighter
}

// GetFighter return duel creator
func (d *Duel) GetFighter() fighter.Fighter {
	return d.Fighter
}

// GetOpponent return duel opponent
func (d *Duel) GetOpponent() fighter.Fighter {
	return d.Opponent
}

// NewDuel create duel
func NewDuel(fighter fighter.Fighter) *Duel {
	d := &Duel{ID: uuid.NewV4(), Fighter: fighter}
	register(d.ID, d)
	return d
}

// register a duel
func register(id uuid.UUID, duel Dueler) {
	duelsMu.Lock()
	defer duelsMu.Unlock()
	if duel == nil {
		// TODO: Change to return error
		panic("duels: Register duel is nil")
	}
	if _, dup := duels[id]; dup {
		// TODO: Change to return error
		panic("duel: Register called twice for duel ")
	}
	duels[id] = duel
}

// GetDuelID returns a duel by its ID
func GetDuelID(id uuid.UUID) Dueler {
	return duels[id]
}

// GetDuelByFighter returns a duel from the given fighter
func GetDuelByFighter(f fighter.Fighter) (Dueler, error) {
	for _, d := range duels {
		if d.GetFighter().Identification() == f.Identification() {
			return d, nil
		}
		if d.GetOpponent().Identification() == f.Identification() {
			return d, nil
		}
	}
	return &Duel{}, errors.New("Duel not found")
}
