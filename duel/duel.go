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
	SetChallenged(fighter fighter.Fighter)
	GetChallenger() fighter.Fighter
	GetChallenged() fighter.Fighter
	Ready() bool
	Finished() bool
	Winner() fighter.Fighter
}

// Duel implements Dueler interface
type Duel struct {
	ID         uuid.UUID
	Challenger fighter.Fighter
	Challenged fighter.Fighter
}

// Result of the duel
type Result struct {
	Challenger *fighter.FightResult
	Challenged *fighter.FightResult
}

// Duel start the combat between two duelists
func (d *Duel) Duel(duelists ...fighter.Fighter) *Result {
	cr := d.Challenger.Fight(d.Challenged)
	cd := d.Challenged.Fight(d.Challenger)
	return &Result{Challenger: cr, Challenged: cd}
}

// Finished return if duel has ended
func (d *Duel) Finished() bool {
	if d.Challenged.Dead() && d.Challenger.Dead() {
		return true
	}
	return false
}

// Winner return the winner
func (d *Duel) Winner() fighter.Fighter {
	var f fighter.Fighter
	if d.Challenged.Dead() {
		f = d.Challenged
	} else if d.Challenger.Dead() {
		f = d.Challenger
	}
	return f
}

// SetChallenged a fighter that accepted the challenge
func (d *Duel) SetChallenged(fighter fighter.Fighter) {
	d.Challenged = fighter
}

// GetChallenged return duel opponent
func (d *Duel) GetChallenged() fighter.Fighter {
	return d.Challenged
}

// GetChallenger return duel creator
func (d *Duel) GetChallenger() fighter.Fighter {
	return d.Challenger
}

// Ready return true if duel is ready to begin
func (d *Duel) Ready() bool {
	if d.Challenged.Ready() && d.Challenger.Ready() {
		return true
	}
	return false
}

// NewDuel create duel
func NewDuel(fighter fighter.Fighter) *Duel {
	d := &Duel{ID: uuid.NewV4(), Challenger: fighter}
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

// GetDuelByID returns a duel by its ID
func GetDuelByID(id uuid.UUID) Dueler {
	return duels[id]
}

// GetDuelByFighter returns a duel from the given fighter
func GetDuelByFighter(f fighter.Fighter) (Dueler, error) {
	for _, d := range duels {
		if d.GetChallenger().Identification() == f.Identification() {
			return d, nil
		}
		if d.GetChallenged().Identification() == f.Identification() {
			return d, nil
		}
	}
	return &Duel{}, errors.New("Duel not found")
}
