package day7

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

var camelCardsMap = map[string]int64{
	"A": 14,
	"K": 13,
	"Q": 12,
	"J": 11,
	"T": 10,
	"9": 9,
	"8": 8,
	"7": 7,
	"6": 6,
	"5": 5,
	"4": 4,
	"3": 3,
	"2": 2,
}

var pairMap = map[string]int64{
	"0":  1,
	"1":  2,
	"2":  3,
	"22": 4,
	"3":  5,
	"23": 6,
	"32": 6,
	"4":  7,
	"5":  8,
}

type player struct {
	hand string
	bid  int64
}

func (p player) calculateHand() int64 {
	var pairs string
	duplCard := make(map[string]bool, 2)
	for _, card := range p.hand {
		if cntPairs := int64(strings.Count(p.hand, string(card))); !duplCard[string(card)] && cntPairs > 1 {
			pairs += fmt.Sprintf("%d", cntPairs)

			duplCard[string(card)] = true
		}
	}

	if pairs == "" {
		return pairMap["0"]
	}

	return pairMap[pairs]
}

func newPlayer(line string) player {
	rawPlayer := strings.Split(line, " ")
	bid, err := strconv.Atoi(rawPlayer[1])
	if err != nil {
		panic(err)
	}

	return player{
		hand: rawPlayer[0],
		bid:  int64(bid),
	}
}

func newPlayers(lines []string) []player {
	players := make([]player, len(lines))
	for idx, line := range lines {
		players[idx] = newPlayer(line)
	}

	return players
}

func CalculatePartOne(lines []string) int64 {
	return calculateWinners(newPlayers(lines))
}

func calculateWinners(players []player) int64 {
	playersHands := make(map[int64][]player)
	for _, p := range players {
		var isInsert bool
		hand := p.calculateHand()

	playersHandsLoop:
		for playerHandIdx, playerHand := range playersHands[hand] {
			for idx, card := range p.hand {
				if camelCardsMap[string(card)] < camelCardsMap[string(playerHand.hand[idx])] {
					playersHands[hand] = slices.Insert(playersHands[hand], playerHandIdx, p)
					isInsert = true
					break playersHandsLoop
				} else if camelCardsMap[string(card)] > camelCardsMap[string(playerHand.hand[idx])] {
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
	for pair := int64(0); pair <= int64(len(pairMap)); pair++ {
		for _, playerHand := range playersHands[pair] {
			res += rank * playerHand.bid
			rank++
		}
	}

	return res
}
