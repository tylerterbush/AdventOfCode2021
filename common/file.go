package common

import (
	"bufio"
	"log"
	"os"
)

func GetLinesFromFile(filepath string) ([]string, error) {
	var (
		err   error
		lines []string
	)

	file, err := os.Open(filepath)
	if err != nil {
		return lines, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return lines, err
	}

	return lines, err
}

func FatalIf(err error) {
	if err != nil {
		log.Fatalf(err.Error())
	}
}
