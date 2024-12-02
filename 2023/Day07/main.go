package main

import (
	"slices"
	"strings"

	"github.com/ysanson/AdventOfCode/pkg"
	"github.com/ysanson/AdventOfCode/pkg/execute"
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

var cardsPart1 = []rune{'2', '3', '4', '5', '6', '7', '8', '9', 'T', 'J', 'Q', 'K', 'A'}
var cardsPart2 = []rune{'J', '2', '3', '4', '5', '6', '7', '8', '9', 'T', 'Q', 'K', 'A'}

func getHandStrengh(hand string, includeJoker bool) int {
	cardOccurences := make(map[rune]int, len(cardsPart1))
	nbJokers := strings.Count(hand, "J")
	for _, card := range cardsPart1 {
		cardOccurences[card] = strings.Count(hand, string(card))
	}
	if pkg.MapContains(cardOccurences, 5) {
		return FiveKind
	} else if pkg.MapContains(cardOccurences, 4) {
		if includeJoker && nbJokers > 0 {
			return FiveKind
		}
		return FourKind
	} else if pkg.MapContains(cardOccurences, 3) && pkg.MapContains(cardOccurences, 2) {
		if includeJoker && nbJokers > 0 {
			return FiveKind
		}
		return FullHouse
	} else if pkg.MapContains(cardOccurences, 3) && !pkg.MapContains(cardOccurences, 2) {
		if includeJoker && nbJokers == 1 {
			return FourKind
		}
		return ThreeKind
	} else if pkg.MapCountOccurrences(cardOccurences, 2) == 2 {
		if includeJoker {
			switch nbJokers {
			case 1:
				return FullHouse
			case 2:
				return FourKind
			default:
				return TwoPair
			}
		}
		return TwoPair
	} else if pkg.MapCountOccurrences(cardOccurences, 2) == 1 {
		if includeJoker && nbJokers > 0 {
			return ThreeKind
		}
		return OnePair
	} else {
		if includeJoker && nbJokers > 0 {
			return OnePair
		}
		return HighCard
	}
}

func getCommonPrefixLength(hand1, hand2 string) int {
	for i := 0; i < len(hand1); i++ {
		if hand1[i] != hand2[i] {
			return i
		}
	}
	return len(hand1)
}

func run(input string) (interface{}, interface{}) {
	lines := strings.Split(input, "\n")
	handsPart1 := make([]Hand, len(lines))
	var split []string
	for idx, line := range lines {
		split = strings.Split(line, " ")
		handsPart1[idx] = Hand{
			hand:     split[0],
			bid:      pkg.MustAtoi(split[1]),
			strength: getHandStrengh(split[0], false),
		}
	}

	slices.SortFunc(handsPart1, func(a, b Hand) int {
		if a.strength != b.strength {
			return a.strength - b.strength
		} else {
			commonPrefix := getCommonPrefixLength(a.hand, b.hand)
			return slices.Index(cardsPart1, pkg.GetChar(a.hand, commonPrefix)) - slices.Index(cardsPart1, pkg.GetChar(b.hand, commonPrefix))
		}
	})

	part1 := 0

	for idx, hand := range handsPart1 {
		part1 += hand.bid * (idx + 1)
	}

	handsPart2 := make([]Hand, len(lines))
	for idx, line := range lines {
		split = strings.Split(line, " ")
		handsPart2[idx] = Hand{
			hand:     split[0],
			bid:      pkg.MustAtoi(split[1]),
			strength: getHandStrengh(split[0], true),
		}
	}

	slices.SortFunc(handsPart2, func(a, b Hand) int {
		if a.strength != b.strength {
			return a.strength - b.strength
		} else {
			commonPrefix := getCommonPrefixLength(a.hand, b.hand)
			return slices.Index(cardsPart2, pkg.GetChar(a.hand, commonPrefix)) - slices.Index(cardsPart2, pkg.GetChar(b.hand, commonPrefix))
		}
	})
	part2 := 0
	for idx, hand := range handsPart2 {
		part2 += hand.bid * (idx + 1)
	}

	return part1, part2
}

func main() {
	execute.Run(run, Tests, Puzzle, true)
}
