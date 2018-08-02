package fighter

// Defender interface
type Defender interface {
	Defend()
}

// DefenseResult of a defence
type DefenseResult struct{}

// Defense implements Defender interface
type Defense struct{}

// Defend an attack with the given defence
func (d *Defense) Defend(opponent Fighter) *DefenseResult {
	return &DefenseResult{}
}

// NewDefense create a defense
func NewDefense() *Defense {
	return &Defense{}
}
