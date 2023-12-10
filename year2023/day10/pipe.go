package day10

const (
	verticalPipe = "|"
	ewPipe       = "-"
	nePipe       = "L"
	nwPipe       = "J"
	swPipe       = "7"
	sePipe       = "F"
	startingPipe = "S"
	ground       = "."
)

type move string

const (
	upMove    move = "up"
	downMove  move = "down"
	leftMove  move = "left"
	rightMove move = "right"
)

var moveMap = map[move][2]int64{
	upMove:    {1, 0},
	downMove:  {-1, 0},
	leftMove:  {0, -1},
	rightMove: {0, 1},
}

var validMovePipes = map[move][]string{
	upMove:    {verticalPipe, swPipe, sePipe},
	downMove:  {verticalPipe, nePipe, nwPipe},
	leftMove:  {ewPipe, sePipe, nePipe},
	rightMove: {ewPipe, swPipe, nwPipe},
}
