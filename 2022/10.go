package twok

import (
	"bufio"
	"bytes"
	"fmt"
	"strconv"
	"strings"
)

type action string

const (
	addx action = "addx"
	noop action = "noop"

	pixelLit  pixel = "#"
	pixelDark pixel = "."
)

type instruction struct {
	action action
	q      int
}

var cycles = map[int]bool{
	20:  true,
	60:  true,
	100: true,
	140: true,
	180: true,
	220: true,
}

var instructionCycles = map[action]int{
	addx: 2,
	noop: 1,
}

var signalStrength = 0

type pixel string

type CRT [240]pixel

var crtLineBreaks = map[int]bool{
	40:  true,
	80:  true,
	120: true,
	160: true,
	200: true,
}

func (c CRT) String() string {
	var b bytes.Buffer
	for i := 1; i <= len(c); i++ {
		b.WriteString(string(c[i-1]))
		if crtLineBreaks[i] {
			b.WriteString("\n")
		}
	}
	return b.String()
}

func (c *CRT) drawPixel(cycle, x int) {
	var pos int
	if cycle >= 200 {
		pos = cycle - 200
	} else if cycle >= 160 {
		pos = cycle - 160
	} else if cycle >= 120 {
		pos = cycle - 120
	} else if cycle >= 80 {
		pos = cycle - 80
	} else if cycle >= 40 {
		pos = cycle - 40
	} else {
		pos = cycle
	}
	if pos == x || pos == x-1 || pos == x+1 {
		c[cycle] = pixelLit
		return
	}
	c[cycle] = pixelDark
}

var crt = CRT{}

func Ten(input string) ([2]interface{}, error) {
	scanner := bufio.NewScanner(strings.NewReader(input))
	cycle := 0
	x := 1
	for scanner.Scan() {
		lineArray := strings.Split(scanner.Text(), " ")
		instr := instruction{
			action: action(lineArray[0]),
		}
		if instr.action == addx {
			instr.q, _ = strconv.Atoi(lineArray[1])
		}

		for i := 1; i <= instructionCycles[instr.action]; i++ {
			crt.drawPixel(cycle, x)
			cycle++
			if cycles[cycle] {
				signalStrength += cycle * x
			}
			if instr.action == addx && i == instructionCycles[instr.action] {
				x += instr.q
			}
		}
	}

	fmt.Printf("%v\n", crt)

	return [2]interface{}{signalStrength, 0}, nil
}
