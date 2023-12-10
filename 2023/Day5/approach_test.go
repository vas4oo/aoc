package day5

import (
	"log"
	"os"
	"strings"
	"testing"
	"time"
)

func TestProcess(t *testing.T) {
	currentTime := time.Now()
	value := process(GetFile(`D:\Projects\aoc\2023\inputs\day5.txt`))
	completedIn := time.Since(currentTime)
	log.Println("The answer is", value, "found in", completedIn)

}

func GetFile(filename string) []string {
	dat, err := os.ReadFile(filename)
	if err != nil {
		panic("Unable to load file: " + err.Error())
	}
	return strings.Split(string(dat), "\n")
}
