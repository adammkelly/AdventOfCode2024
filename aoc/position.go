package aoc

type PositionPoint interface {
	Neighbors()
	Next_neighbor_by_direction()
}

type Point struct {
	X int
	Y int
}

var TOP_LEFT string = "TOP_LEFT"
var TOP string = "TOP"
var TOP_RIGHT string = "TOP_RIGHT"
var LEFT string = "LEFT"
var RIGHT string = "RIGHT"
var BOTTOM_LEFT string = "BOTTOM_LEFT"
var BOTTOM string = "BOTTOM"
var BOTTOM_RIGHT string = "BOTTOM_RIGHT"

var POSITIONS = map[string]Point{
	TOP_LEFT:     {-1, -1},
	TOP:          {0, -1},
	TOP_RIGHT:    {1, -1},
	LEFT:         {-1, 0},
	RIGHT:        {1, 0},
	BOTTOM_LEFT:  {-1, 1},
	BOTTOM:       {0, 1},
	BOTTOM_RIGHT: {1, 1},
}

func (p Point) get_next_position(next Point) Point {
	var cur_x int = p.X + next.X
	var cur_y int = p.Y + next.Y
	return Point{cur_x, cur_y}
}

func (p Point) Neighbors() map[string]Point {
	neighbors := make(map[string]Point)

	for position_name, j := range POSITIONS {
		neighbors[position_name] = p.get_next_position(j)
	}
	return neighbors
}

func (p Point) Next_neighbor_by_direction(direction string) Point {
	return p.get_next_position(POSITIONS[direction])
}
