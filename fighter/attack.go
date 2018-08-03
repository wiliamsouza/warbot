package fighter

// Attacker interface
type Attacker interface {
	Attack(Fighter) *AttackResult
}

// AttackResult of an attack
type AttackResult struct{}

// Attack implements Attacker interface
type Attack struct {
	Type string
}

// Attack an opponent with the given attack
func (a *Attack) Attack(opponent Fighter) *AttackResult {
	return &AttackResult{}
}

// NewAttack create an attack
func NewAttack(t string) *Attack {
	return &Attack{Type: t}
}
