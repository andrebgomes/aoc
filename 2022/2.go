package twok

import (
	"bufio"
	"strings"
)

type shape string

const (
	rock     shape = "Rock"
	paper    shape = "Paper"
	scissors shape = "Scissors"

	win  = 6
	draw = 3
	loss = 0
)

var opponentMap = map[string]shape{
	"A": rock,
	"B": paper,
	"C": scissors,
}

var meMap = map[string]shape{
	"X": rock,
	"Y": paper,
	"Z": scissors,
}

var rules = map[shape]map[shape]int{
	rock: {
		rock:     draw,
		paper:    win,
		scissors: loss,
	},
	paper: {
		rock:     loss,
		paper:    draw,
		scissors: win,
	},
	scissors: {
		rock:     win,
		paper:    loss,
		scissors: draw,
	},
}

var realMeMap = map[string]int{
	"X": loss,
	"Y": draw,
	"Z": win,
}

var realRules = map[shape]map[int]shape{
	rock: {
		win:  paper,
		draw: rock,
		loss: scissors,
	},
	paper: {
		win:  scissors,
		draw: paper,
		loss: rock,
	},
	scissors: {
		win:  rock,
		draw: scissors,
		loss: paper,
	},
}

var shapePoints = map[shape]int{
	rock:     1,
	paper:    2,
	scissors: 3,
}

func Two(input string) ([2]interface{}, error) {
	totalScore := 0
	totalScoreRealRules := 0

	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		line := scanner.Text()
		op := string(line[0])
		me := string(line[2])

		totalScore += rules[opponentMap[op]][meMap[me]] + shapePoints[meMap[me]]

		myShape := realRules[opponentMap[op]][realMeMap[me]]
		totalScoreRealRules += realMeMap[me] + shapePoints[myShape]
	}

	return [2]interface{}{totalScore, totalScoreRealRules}, nil
}
