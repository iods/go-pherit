package keyboard

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

/*
GetFloat Captures and trims the user's keyboard input returning
  it as a float64 Data Type. */
func GetFloat() (float64, error) {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		return 0, err
	}

	input = strings.TrimSpace(input)
	number, err := strconv.ParseFloat(input, 64)
	if err != nil {
		return 0, err
	}
	return number, nil
}