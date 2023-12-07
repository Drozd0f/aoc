package day7

import (
	"fmt"
	"slices"
	"strings"
)

var pairMapPartTwo = map[string]int64{
	"1":  1,
	"2":  2,
	"22": 3,
	"3":  4,
	"23": 5,
	"32": 5,
	"4":  6,
	"5":  7,
}

var camelCardsMapParTwo = map[string]int64{
	"A": 14,
	"K": 13,
	"Q": 12,
	"T": 10,
	"9": 9,
	"8": 8,
	"7": 7,
	"6": 6,
	"5": 5,
	"4": 4,
	"3": 3,
	"2": 2,
	"J": 1, // joker
}

func (p player) calculateHandPartTwo() int64 {
	var jokerCnt int64
	rawPairs := make([]int64, 0, 2)
	duplCard := make(map[string]bool, 2)
	for _, card := range p.hand {
		if "J" == string(card) {
			jokerCnt++

			continue
		}

		if cntPairs := int64(strings.Count(p.hand, string(card))); !duplCard[string(card)] && cntPairs > 1 {
			rawPairs = append(rawPairs, cntPairs)

			duplCard[string(card)] = true
		}
	}

	if len(rawPairs) == 0 {
		if jokerCnt < 1 || jokerCnt == 5 {
			return pairMapPartTwo[fmt.Sprintf("%d", jokerCnt)]
		}

		return pairMapPartTwo[fmt.Sprintf("%d", jokerCnt+1)]
	}

	var pairs string
	maxPair := slices.Max(rawPairs)
	for _, rawPair := range rawPairs {
		if rawPair == maxPair && jokerCnt > 0 {
			pairs += fmt.Sprintf("%d", rawPair+jokerCnt)
			jokerCnt = 0

			continue
		}

		pairs += fmt.Sprintf("%d", rawPair)
	}

	return pairMapPartTwo[pairs]
}

func CalculatePartTwo(lines []string) int64 {
	return calculateWinnersPartTwo(newPlayers(lines))
}

func calculateWinnersPartTwo(players []player) int64 {
	playersHands := make(map[int64][]player)
	for _, p := range players {
		var isInsert bool
		hand := p.calculateHandPartTwo()

	playersHandsLoop:
		for playerHandIdx, playerHand := range playersHands[hand] {
			for idx, card := range p.hand {
				if camelCardsMapParTwo[string(card)] < camelCardsMapParTwo[string(playerHand.hand[idx])] {
					playersHands[hand] = slices.Insert(playersHands[hand], playerHandIdx, p)
					isInsert = true
					break playersHandsLoop
				} else if camelCardsMapParTwo[string(card)] > camelCardsMapParTwo[string(playerHand.hand[idx])] {
					break
				}
			}
		}

		if !isInsert {
			playersHands[hand] = append(playersHands[hand], p)
		}
	}

	var res int64
	rank := int64(1)
	for pair := int64(0); pair <= int64(len(pairMapPartTwo)); pair++ {
		for _, playerHand := range playersHands[pair] {
			res += rank * playerHand.bid
			rank++
		}
	}

	return res
}
