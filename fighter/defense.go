package fighter

// Defender interface
type Defender interface {
	Defend(Fighter) *DefenseResult
}

// DefenseResult of a defence
type DefenseResult struct{}

// Defense implements Defender interface
type Defense struct {
	Type string
}

// Defend an attack with the given defence
func (d *Defense) Defend(opponent Fighter) *DefenseResult {
	return &DefenseResult{}
}

// NewDefense create a defense
func NewDefense(t string) *Defense {
	return &Defense{Type: t}
}
