package twok

import (
	"bufio"
	"regexp"
	"strconv"
	"strings"
)

type stacks map[int][]string

func (s stacks) ToString() string {
	result := []string{}
	for i := 1; i <= len(s); i++ {
		result = append(result, s[i][len(s[i])-1])
	}

	return strings.Join(result, "")
}

var cratesLines = []string{}
var stacks1 stacks = map[int][]string{}
var stacks2 stacks = map[int][]string{}

func Five(input string) ([2]interface{}, error) {

	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		line := scanner.Text()

		exp := `move ([0-9]*) from ([0-9]*) to ([0-9]*)`
		re := regexp.MustCompile(exp)
		if re.MatchString(line) {
			matches := re.FindStringSubmatch(line)

			from, _ := strconv.Atoi(matches[2])
			to, _ := strconv.Atoi(matches[3])
			nrCrates, _ := strconv.Atoi(matches[1])
			rearrangeInvertedOrder(stacks1, from, to, nrCrates)
			rearrangeSameOrder(stacks2, from, to, nrCrates)
			continue
		}

		if line == "" {
			fillStacks(stacks1)
			fillStacks(stacks2)
			continue
		}
		cratesLines = append(cratesLines, line)
	}

	return [2]interface{}{stacks1.ToString(), stacks2.ToString()}, nil
}

func fillStacks(s stacks) {
	for i := len(cratesLines) - 2; i >= 0; i-- {
		number := 0
		for j := 1; j < len(cratesLines[i]); j += 4 {
			number++
			crate := rune(cratesLines[i][j])
			if crate == 32 {
				continue
			}
			s[number] = append(s[number], string(rune(cratesLines[i][j])))
		}
	}

}

func rearrangeInvertedOrder(s stacks, stackFrom, stackTo, nrCrates int) {
	crates := []string{}

	fromLen := len(s[stackFrom])
	for i := fromLen - 1; i > fromLen-nrCrates-1; i-- {
		crates = append(crates, s[stackFrom][i])
		s[stackFrom] = append(s[stackFrom][:i], s[stackFrom][i+1:]...)
	}

	s[stackTo] = append(s[stackTo], crates...)
}

func rearrangeSameOrder(s stacks, stackFrom, stackTo, nrCrates int) {
	crates := []string{}

	fromLen := len(s[stackFrom])
	for i := fromLen - 1; i > fromLen-nrCrates-1; i-- {
		crates = append(crates, s[stackFrom][i])
		s[stackFrom] = append(s[stackFrom][:i], s[stackFrom][i+1:]...)
	}

	for i := len(crates) - 1; i >= 0; i-- {
		s[stackTo] = append(s[stackTo], crates[i])
	}
}
