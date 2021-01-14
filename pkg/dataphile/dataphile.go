package dataphile

import (
	"bufio"
	"github.com/iods/go-golph/internal/common"
	"os"
)

func GetStrings(fileName string) []string {
	var lines []string
	file, err := os.Open(fileName)
	if os.IsNotExist(err) {
		return nil
	}
	common.ErrorCheck(err)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	common.ErrorCheck(scanner.Err())
	return lines
}

