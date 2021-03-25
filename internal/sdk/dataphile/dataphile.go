package dataphile

import (
	"bufio"
	"github.com/iods/go-pherit/internal/common"
	"os"
)

func GetStrings(fileName string) []string {
	var lines []string
	file, err := os.Open(fileName)
	if os.IsNotExist(err) {
		return nil
	}
	errors.ErrorCheck(err)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	errors.ErrorCheck(scanner.Err())
	return lines
}

