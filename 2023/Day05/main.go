package main

import (
	"regexp"
	"slices"
	"strings"

	"github.com/ysanson/AdventOfCode/2023/pkg"
	"github.com/ysanson/AdventOfCode/2023/pkg/execute"
)

type LUT string

const (
	SeedToSoil        LUT = "seed-to-soil map:"
	SoilToFertilizer  LUT = "soil-to-fertilizer map:"
	FertilizerToWater LUT = "fertilizer-to-water map:"
	WaterToLight      LUT = "water-to-light map:"
	LightToTemp       LUT = "light-to-temperature map:"
	TempToHumid       LUT = "temperature-to-humidity map:"
	HumidToLocation   LUT = "humidity-to-location map:"
)

type Range struct {
	from, to, difference int
}

type Ranges []Range

type Tables struct {
	seedToSoil, soilToFertilizer, fertilizerToWater, waterToLight, ligthToTemp, tempToHumid, humidToLocation Ranges
}

var numRegex = regexp.MustCompile(`\d+\s?`)

func createRange(inputs []string) Ranges {
	var parsed []string
	var source, dest, length int
	ranges := make([]Range, len(inputs))
	for index, line := range inputs {
		parsed = strings.Split(line, " ")
		source = pkg.MustAtoi(parsed[1])
		dest = pkg.MustAtoi(parsed[0])
		length = pkg.MustAtoi(parsed[2])
		ranges[index] = Range{
			from:       source,
			to:         source + length - 1,
			difference: dest - source,
		}
	}

	return ranges
}

func extractTable(input string, header LUT) Ranges {
	start := strings.Index(input, string(header))
	end := strings.Index(input[start:], "\n\n")
	var data []string
	if end == -1 {
		data = strings.Split(input[start:], "\n")
	} else {
		data = strings.Split(input[start:start+end], "\n")
	}

	data = slices.DeleteFunc(data, func(elt string) bool {
		return !numRegex.MatchString(elt)
	})
	return createRange(data)
}

func extractSeeds(input string) []int {
	seedLine := strings.Split(input, "\n")[0]
	numbers := strings.Split(seedLine, " ")[1:]
	seeds := make([]int, len(numbers))
	for index, num := range numbers {
		seeds[index] = pkg.MustAtoi(num)
	}
	return seeds
}

func generateTables(input string) Tables {
	var tables Tables
	tables.seedToSoil = extractTable(input, SeedToSoil)
	tables.soilToFertilizer = extractTable(input, SoilToFertilizer)
	tables.fertilizerToWater = extractTable(input, FertilizerToWater)
	tables.waterToLight = extractTable(input, WaterToLight)
	tables.ligthToTemp = extractTable(input, LightToTemp)
	tables.tempToHumid = extractTable(input, TempToHumid)
	tables.humidToLocation = extractTable(input, HumidToLocation)
	return tables
}

func getMatchingNumber(sourceNumber int, ranges Ranges) int {
	for _, r := range ranges {
		if r.from <= sourceNumber && sourceNumber <= r.to {
			return sourceNumber + r.difference
		}
	}
	return sourceNumber
}

func convertSeedToLocation(seed int, tables Tables) int {
	finalNumber := getMatchingNumber(seed, tables.seedToSoil)
	finalNumber = getMatchingNumber(finalNumber, tables.soilToFertilizer)
	finalNumber = getMatchingNumber(finalNumber, tables.fertilizerToWater)
	finalNumber = getMatchingNumber(finalNumber, tables.waterToLight)
	finalNumber = getMatchingNumber(finalNumber, tables.ligthToTemp)
	finalNumber = getMatchingNumber(finalNumber, tables.tempToHumid)
	return getMatchingNumber(finalNumber, tables.humidToLocation)
}

func run(input string) (interface{}, interface{}) {
	seeds := extractSeeds(input)
	tables := generateTables(input)
	locations := make([]int, len(seeds))
	for index, seed := range seeds {
		locations[index] = convertSeedToLocation(seed, tables)
	}

	return pkg.Min(locations...), 0
}

func main() {
	execute.Run(run, Tests, Puzzle, true)
}
