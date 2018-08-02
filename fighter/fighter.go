package fighter

// Fighter interface
type Fighter interface {
	Fight(opponent Fighter) *FightResult
	SetAttack(atack Attacker)
	SetDefense(defense Defender)
}

// FightResult of the fight
type FightResult struct{}

// F implements Fighter interface
type F struct {
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

// NewFighter create a fighter
func NewFighter(health int) *F {
	return &F{Health: health}
}
