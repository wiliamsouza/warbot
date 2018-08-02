package duel

import "github.com/wiliamsouza/warbot/fighter"

// Dueler interface
type Dueler interface {
	Duel(duelists ...fighter.Fighter) *DuelResult
}

// Duel implements Dueler interface
type Duel struct{}

// DuelResult of the duel
type DuelResult struct{}

// Duel start the combat between two duelists
func (d *Duel) Duel(duelists ...fighter.Fighter) *DuelResult {
	return &DuelResult{}
}

// NewDuel create duel
func NewDuel() *Duel {
	return &Duel{}
}
