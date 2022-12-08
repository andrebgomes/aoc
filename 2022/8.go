package twok

import (
	"bufio"
	"bytes"
	"strconv"
	"strings"
)

func Eight(input string) ([2]interface{}, error) {
	grid := make([][]int, 0)
	visible := 0

	scanner := bufio.NewScanner(strings.NewReader(input))
	row := 0
	for scanner.Scan() {
		grid = append(grid, []int{})
		line := scanner.Text()
		digitsArray := strings.Split(line, "")
		for col := 0; col < len(line); col++ {
			height, _ := strconv.Atoi(digitsArray[col])
			grid[row] = append(grid[row], height)
		}
		row++
	}

	bestScore := 0
	for row := 0; row < len(grid); row++ {
		for col := 0; col < len(grid[row]); col++ {
			if isVisible(grid, row, col) {
				visible++
			}
			score := getScenicScore(grid, row, col)
			if score > bestScore {
				bestScore = score
			}
		}
	}

	return [2]interface{}{visible, bestScore}, nil
}

func isVisible(grid [][]int, row, col int) bool {
	if row == len(grid)-1 || col == len(grid[row])-1 || row == 0 || col == 0 {
		return true
	}
	c := make(chan bool)
	height := grid[row][col]
	// up
	go func() {
		for i := row - 1; i >= 0; i-- {
			if grid[i][col] >= height {
				c <- false
				return
			}
		}
		c <- true
	}()

	// down
	go func() {
		for i := row + 1; i < len(grid); i++ {
			if grid[i][col] >= height {
				c <- false
				return
			}
		}
		c <- true
	}()

	// left
	go func() {
		for i := col - 1; i >= 0; i-- {
			if grid[row][i] >= height {
				c <- false
				return
			}
		}
		c <- true
	}()

	// right
	go func() {
		for i := col + 1; i < len(grid[row]); i++ {
			if grid[row][i] >= height {
				c <- false
				return
			}
		}
		c <- true
	}()

	a, b, s, n := <-c, <-c, <-c, <-c
	return a || b || s || n
}

func getScenicScore(grid [][]int, row, col int) int {
	if row == len(grid)-1 || col == len(grid[row])-1 || row == 0 || col == 0 {
		return 0
	}
	c := make(chan int)
	height := grid[row][col]
	score := 1
	// up
	go func(score int) {
		for i := row - 1; i > 0; i-- {
			if grid[i][col] >= height {
				break
			}
			score++
		}
		c <- score
	}(score)

	// down
	go func(score int) {
		for i := row + 1; i < len(grid)-1; i++ {
			if grid[i][col] >= height {
				break
			}
			score++
		}
		c <- score
	}(score)

	// left
	go func(score int) {
		for i := col - 1; i > 0; i-- {
			if grid[row][i] >= height {
				break
			}
			score++
		}
		c <- score
	}(score)

	// right
	go func(score int) {
		for i := col + 1; i < len(grid[row])-1; i++ {
			if grid[row][i] >= height {
				break
			}
			score++
		}
		c <- score
	}(score)

	a, b, s, n := <-c, <-c, <-c, <-c
	return a * b * s * n
}

func toString(grid [][]int) string {
	var b bytes.Buffer
	for row := 0; row < len(grid); row++ {
		for col := 0; col < len(grid[row]); col++ {
			b.WriteString(strconv.Itoa(grid[row][col]))
		}
		b.WriteString("\n")
	}
	return b.String()
}
