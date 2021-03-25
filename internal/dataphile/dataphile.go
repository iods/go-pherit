package dataphile

import (
	"bufio"
	"os"

	"github.com/iods/go-pherit/internal/helpers/error"
)

func GetStrings(fileName string) []string {
	var lines []string
	file, err := os.Open(fileName)
	if os.IsNotExist(err) {
		return nil
	}
	error.HandleError(err)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	error.HandleError(scanner.Err())
	return lines
}

