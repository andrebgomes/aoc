package twok

import (
	"bufio"
	"regexp"
	"strconv"
	"strings"
)

func Four(input string) ([2]interface{}, error) {
	pairsContains := 0
	pairsOverlap := 0

	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		line := scanner.Text()
		exp := `([0-9]*)-([0-9]*),([0-9]*)-([0-9]*)`
		re := regexp.MustCompile(exp)
		matches := re.FindStringSubmatch(line)

		a1, _ := strconv.Atoi(matches[1])
		a2, _ := strconv.Atoi(matches[2])
		b1, _ := strconv.Atoi(matches[3])
		b2, _ := strconv.Atoi(matches[4])

		if contains(a1, a2, b1, b2) {
			pairsContains++
		}

		if overlap(a1, a2, b1, b2) {
			pairsOverlap++
		}
	}

	return [2]interface{}{pairsContains, pairsOverlap}, nil
}

func contains(a1, a2, b1, b2 int) bool {
	aContainsB := a1 <= b1 && a2 >= b2
	bContainsA := b1 <= a1 && b2 >= a2
	return aContainsB || bContainsA
}

func overlap(a1, a2, b1, b2 int) bool {
	aOverlapB := (a1 >= b1 && a1 <= b2) || (a2 >= b1 && a2 <= b2)
	bOverlapA := (b1 >= a1 && b1 <= a2) || (b2 >= a1 && b2 <= a2)
	return aOverlapB || bOverlapA
}
