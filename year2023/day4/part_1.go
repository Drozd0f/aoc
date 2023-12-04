package day4

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"slices"
	"strings"
)

func ReadAllFile(fileName string) []string {
	file, err := os.Open(fmt.Sprintf("input/%s", fileName))
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err = scanner.Err(); err != nil {
		panic(err)
	}

	return lines
}

func CalculatePart1(fileName string) int64 {
	lines := ReadAllFile(fileName)

	var result int64
	for _, line := range lines {
		result += calculatePoints(extractNumbers(line))
	}

	return result
}

var extractNumbersRegex = regexp.MustCompile(`\d+`)

func extractNumbers(line string) ([]string, []string) {
	splitLine := strings.Split(strings.Split(line, ": ")[1], " | ")

	return extractNumbersRegex.FindAllString(splitLine[0], -1), extractNumbersRegex.FindAllString(splitLine[1], -1)
}

func calculatePoints(winNumbers, numbers []string) int64 {
	var points int64
	for _, winNumber := range winNumbers {
		if slices.Contains(numbers, winNumber) {
			if points == 0 {
				points++

				continue
			}

			points *= 2
		}
	}

	return points
}
