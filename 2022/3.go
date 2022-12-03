package twok

import (
	"bufio"
	"strings"
	"unicode"
)

const (
	lower = "abcdefghijklmnopqrstuvwxyz"
	upper = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

func Three(input string) ([2]interface{}, error) {
	prioritySum := 0
	priorityGroupSum := 0

	scanner := bufio.NewScanner(strings.NewReader(input))
	group := [2]string{"", ""}
	i := 0
	for scanner.Scan() {
		line := scanner.Text()
		lineLen := len(line)

		comp1 := line[:lineLen/2]
		comp2 := line[lineLen/2:]

		prioritySum += getPriority(findDuplicated(comp1, comp2))

		if i == 2 {
			priorityGroupSum += getPriority(findTriplicated(group[0], group[1], line))
			group = [2]string{"", ""}
			i = 0
			continue
		}
		if i == 1 {
			group[i] = line
			i++
			continue
		}
		group[i] = line
		i++
	}

	return [2]interface{}{prioritySum, priorityGroupSum}, nil
}

func findDuplicated(first, second string) rune {
	firstMap := make(map[rune]bool)
	for _, letter := range first {
		firstMap[letter] = true
	}

	for _, letter := range second {
		if firstMap[letter] {
			return letter
		}
	}

	return -1
}

func findTriplicated(first, second, third string) rune {
	firstMap := make(map[rune]bool)
	for _, letter := range first {
		firstMap[letter] = true
	}

	secondMap := make(map[rune]bool)
	for _, letter := range second {
		secondMap[letter] = true
	}

	for _, letter := range third {
		if firstMap[letter] && secondMap[letter] {
			return letter
		}
	}

	return -1
}

func getPriority(letter rune) int {
	if unicode.IsLower(letter) {
		for k, l := range lower {
			if l == letter {
				return k + 1
			}
		}
	}

	for k, l := range upper {
		if l == letter {
			return k + 27
		}
	}
	return 0
}
