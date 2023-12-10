package helpers

import (
	"os"
	"strconv"
	"strings"
)

func ReadLines(filename string) []string {
	pwd, _ := os.Getwd()
	data, _ := os.ReadFile(pwd + "\\inputs\\" + filename)
	if data[len(data)-1] == '\n' {
		data = data[:len(data)-1]
	}
	return strings.Split(string(data), "\n")
}

func InBetween(i, min, max int) bool {
	if (i >= min) && (i <= max) {
		return true
	}
	return false
}

func GetNumber(stringNumber string) (k int) {
	if v, err := strconv.Atoi(stringNumber); err == nil {
		k = v
	}

	return
}
