package twok

import (
	"bufio"
	"strconv"
	"strings"
)

func One(input string) ([]interface{}, error) {
	mostCals := 0
	mostCalsArray := [3]int{0, 0, 0}

	currCals := 0

	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		cals, _ := strconv.Atoi(scanner.Text())

		if cals != 0 {
			currCals += cals
			continue
		}

		if currCals > mostCals {
			mostCals = currCals
		}
		handleTopThreeCals(&mostCalsArray, currCals)
		currCals = 0
	}

	topThreeCals := mostCalsArray[0] + mostCalsArray[1] + mostCalsArray[2]

	return []interface{}{mostCals, topThreeCals}, nil
}

func handleTopThreeCals(array *[3]int, cals int) {
	index := 0
	min := array[0]

	// Find min
	for i := 1; i < 3; i++ {
		if array[i] < min {
			min = array[i]
			index = i
		}
	}

	if cals > min {
		array[index] = cals
	}
}
