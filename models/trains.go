package models

type Direction int

const (
	North Direction = iota
	South
)

func (d Direction) String() string {
	return [...]string{"North", "South"}[d]
}

type Train struct {
	RouteID     string
	Direction   Direction
	ArrivalTime int64
}
