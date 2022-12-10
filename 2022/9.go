package twok

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

type direction string

type knots [][2]int

const (
	up    direction = "U"
	down  direction = "D"
	right direction = "R"
	left  direction = "L"
)

var knotsPositions1 = knots{
	{0, 0},
	{0, 0},
}

var knotsPositions2 = knots{
	{0, 0},
	{0, 0},
	{0, 0},
	{0, 0},
	{0, 0},
	{0, 0},
	{0, 0},
	{0, 0},
	{0, 0},
	{0, 0},
}

var tailVisited1 = map[string]bool{
	"0-0": true,
}
var tailVisited2 = map[string]bool{
	"0-0": true,
}

func Nine(input string) ([2]interface{}, error) {
	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		lineArray := strings.Split(scanner.Text(), " ")
		d := direction(lineArray[0])
		positions, _ := strconv.Atoi(lineArray[1])

		iterate(knotsPositions1, tailVisited1, d, positions)
		iterate(knotsPositions2, tailVisited2, d, positions)
	}

	return [2]interface{}{len(tailVisited1), len(tailVisited2)}, nil
}

func iterate(knots knots, tailVisited map[string]bool, d direction, positions int) {
	for i := 1; i <= positions; i++ {
		moveHeadOnePosition(knots, d)

		for k := 1; k < len(knots); k++ {
			moveKnot(&knots[k-1], &knots[k], tailVisited, k == len(knots)-1)
		}

	}
}

func moveHeadOnePosition(knots knots, d direction) {
	switch d {
	case up:
		knots[0][0] = knots[0][0] - 1
	case down:
		knots[0][0] = knots[0][0] + 1
	case right:
		knots[0][1] = knots[0][1] + 1
	case left:
		knots[0][1] = knots[0][1] - 1
	}
}

func moveKnot(prev, curr *[2]int, tailVisited map[string]bool, isTail bool) {
	diffX := prev[0] - curr[0]
	diffY := prev[1] - curr[1]
	if diffX == 2 || diffX == -2 || diffY == 2 || diffY == -2 {
		if diffX > 0 {
			curr[0]++
		}
		if diffX < 0 {
			curr[0]--
		}
		if diffY > 0 {
			curr[1]++
		}
		if diffY < 0 {
			curr[1]--
		}
	}

	if isTail {
		tailVisited[fmt.Sprintf("%d-%d", curr[0], curr[1])] = true
	}
}
