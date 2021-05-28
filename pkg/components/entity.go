package components

type COMPONENTS byte

const (
	POSITION COMPONENTS = 1 << iota
	MOMENTUM
)

type Entity struct {
	ID       uint32
	compMask COMPONENTS
}
