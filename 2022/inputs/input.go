package inputs

import (
	"fmt"
	"os"
)

// ReadInput reads the entire file into memory and returns a string of it.
func ReadInput(year, day int) (string, error) {
	input, err := os.ReadFile(fmt.Sprintf("%d/inputs/%d", year, day))
	return string(input), err
}
