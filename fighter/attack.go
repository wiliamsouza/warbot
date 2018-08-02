package fighter

// Attacker interface
type Attacker interface {
	Attack() *AttackResult
}

// AttackResult of an attack
type AttackResult struct{}

// Attack implements Attacker interface
type Attack struct{}

// Attack an opponent with the given attack
func (a *Attack) Attack(opponent Fighter) *AttackResult {
	return &AttackResult{}
}

// NewAttack create an attack
func NewAttack() *Attack {
	return &Attack{}
}
