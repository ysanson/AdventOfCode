package main

import (
	"math"
	"regexp"
	"slices"
	"strings"
	"sync"

	"github.com/ysanson/AdventOfCode/pkg"
	"github.com/ysanson/AdventOfCode/pkg/execute"
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

func extractSeedRanges(input string) Ranges {
	seedLine := strings.Split(input, "\n")[0]
	numbers := strings.Split(seedLine, " ")[1:]
	ranges := make(Ranges, 0, len(numbers)/2)
	var from, length int
	for i := 0; i < len(numbers); i += 2 {
		from = pkg.MustAtoi(numbers[i])
		length = pkg.MustAtoi(numbers[i+1])
		ranges = append(ranges, Range{
			from:       from,
			to:         from + length,
			difference: 0,
		})
	}
	return ranges
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

func analyzeSeedRange(seedRange Range, tables Tables, c chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	var location int
	minLocation := math.MaxInt
	for i := seedRange.from; i < seedRange.to; i++ {
		location = convertSeedToLocation(i, tables)
		if location < minLocation {
			minLocation = location
		}
	}
	c <- minLocation
}

func run(input string) (interface{}, interface{}) {
	P1seeds := extractSeeds(input)
	tables := generateTables(input)
	P1locations := make([]int, len(P1seeds))
	for index, seed := range P1seeds {
		P1locations[index] = convertSeedToLocation(seed, tables)
	}

	seedPairs := extractSeedRanges(input)
	var wg sync.WaitGroup
	ch := make(chan int)
	p2Locations := make([]int, 0, len(seedPairs))

	for _, seedRange := range seedPairs {
		wg.Add(1)
		go analyzeSeedRange(seedRange, tables, ch, &wg)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	for i := range ch {
		p2Locations = append(p2Locations, i)
	}

	return pkg.Min(P1locations...), pkg.Min(p2Locations...)
}

func main() {
	execute.Run(run, Tests, Puzzle, true)
}
