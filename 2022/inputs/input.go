package inputs

import (
	"fmt"
	"os"
)

// ReadInput reads the entire file into memory and returns a string of it.
func ReadInput(day int) (string, error) {
	input, err := os.ReadFile(fmt.Sprintf("2022/inputs/%d", day))
	return string(input), err
}
