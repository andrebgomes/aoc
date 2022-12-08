package main

import (
	"flag"
	"fmt"
	"log"

	twok "aoc/2022"
	"aoc/2022/inputs"
)

// dayFunc represents the type that day funtions must be of.
// they should be given a string representing the input and return an array of
// two indexes, one for each result of the two puzzles.
type dayFunc func(string) ([2]interface{}, error)

// functions include all implemented day funcs
var functions = map[int]map[int]dayFunc{
	2022: {
		1: twok.One,
		2: twok.Two,
		3: twok.Three,
		4: twok.Four,
		5: twok.Five,
		6: twok.Six,
		7: twok.Seven,
		8: twok.Eight,
	},
}

func main() {
	// Define and validate year and day flags
	year := flag.Int("y", 0, "year")
	day := flag.Int("d", 0, "day")
	flag.Parse()

	if year == nil || *year != 2022 {
		log.Fatal("available years are: 2022")
	}

	if day == nil || *day < 1 || *day > 25 {
		log.Fatal("choose a day between 1 and 25")
	}

	// Read input
	input, err := inputs.ReadInput(*year, *day)
	if err != nil {
		log.Fatalf("error reading input for year/day %d/%d: %v", *year, *day, err)
	}

	// Execute
	if f, ok := functions[*year][*day]; ok {
		res, err := f(input)
		if err != nil {
			log.Fatalf("error executing the function for year/day %d/%d: %v", *year, *day, err)
		}

		// Print the solution
		fmt.Println(res)
		return
	}

	log.Fatalf("function for year/day %d/%d is not yet available", *year, *day)
}
