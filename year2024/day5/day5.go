package day5

import (
	"log"
	"strings"
)

type Rules map[string]map[int][]string

const (
	up   = 0
	down = 1
)

func parseInput(lines []string) (Rules, [][]string) {
	var stopLineIdx int

	rules := make(Rules, len(lines))
	for idx, line := range lines {
		if line == "" {
			stopLineIdx = idx
			break
		}

		split := strings.Split(line, "|")
		if len(split) != 2 {
			log.Fatalf("[%d] invalid line in day5: %s", idx+1, line)
		}

		rootNum, subNum := split[0], split[1]
		if rules[rootNum] == nil {
			rules[rootNum] = make(map[int][]string, 2)
			rules[rootNum][up] = make([]string, 0, len(lines))
			rules[rootNum][down] = make([]string, 0, len(lines))
		}

		if rules[subNum] == nil {
			rules[subNum] = make(map[int][]string, 2)
			rules[subNum][up] = make([]string, 0, len(lines))
			rules[subNum][down] = make([]string, 0, len(lines))
		}

		rules[rootNum][down] = append(rules[rootNum][down], subNum)
		rules[subNum][up] = append(rules[subNum][up], rootNum)
	}

	pages := make([][]string, len(lines[stopLineIdx+1:]))
	for idx, line := range lines[stopLineIdx+1:] {
		pages[idx] = strings.Split(line, ",")
	}

	return rules, pages
}
