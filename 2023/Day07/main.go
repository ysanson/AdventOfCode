package main

import (
	"fmt"
	"slices"
	"strings"

	"github.com/ysanson/AdventOfCode/2023/pkg"
	"github.com/ysanson/AdventOfCode/2023/pkg/execute"
)

const (
	FiveKind  int = 7
	FourKind  int = 6
	FullHouse int = 5
	ThreeKind int = 4
	TwoPair   int = 3
	OnePair   int = 2
	HighCard  int = 1
)

type Hand struct {
	hand     string
	bid      int
	strength int
}

var cards = []rune{'2', '3', '4', '5', '6', '7', '8', '9', 'T', 'J', 'Q', 'K', 'A'}

func getHandStrengh(hand string) int {
	cardOccurences := make(map[rune]int, len(cards))
	for _, card := range cards {
		cardOccurences[card] = strings.Count(hand, string(card))
	}
	if pkg.MapContains(cardOccurences, 5) {
		return FiveKind
	} else if pkg.MapContains(cardOccurences, 4) {
		return FourKind
	} else if pkg.MapContains(cardOccurences, 3) && pkg.MapContains(cardOccurences, 2) {
		return ThreeKind
	} else if pkg.MapContains(cardOccurences, 3) && !pkg.MapContains(cardOccurences, 2) {
		return FullHouse
	} else if pkg.MapCountOccurrences(cardOccurences, 2) == 2 {
		return TwoPair
	} else if pkg.MapCountOccurrences(cardOccurences, 2) == 1 {
		return OnePair
	} else {
		return HighCard
	}
}

func convertHandToNumber(hand string) int {
	var num strings.Builder
	for _, card := range hand {
		num.WriteString(fmt.Sprint(slices.Index(cards, card)))
	}
	str := num.String()
	return pkg.MustAtoi(str)
}

func run(input string) (interface{}, interface{}) {
	lines := strings.Split(input, "\n")
	hands := make([]Hand, len(lines))
	var split []string
	var h string
	for idx, line := range lines {
		split = strings.Split(line, " ")
		h = split[0]
		hands[idx] = Hand{
			hand:     h,
			bid:      pkg.MustAtoi(split[1]),
			strength: getHandStrengh(h),
		}
	}

	slices.SortFunc(hands, func(a, b Hand) int {
		if a.strength != b.strength {
			return a.strength - b.strength
		} else {
			return convertHandToNumber(a.hand) - convertHandToNumber(b.hand)
		}
	})

	part1 := 0

	for idx, hand := range hands {
		part1 += hand.strength * (idx + 1)
	}

	return part1, 0
}

func main() {
	execute.Run(run, Tests, Puzzle, true)
}
