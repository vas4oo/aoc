package helpers

import (
	"os"
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
