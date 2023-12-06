package main

import (
	"math"
	"regexp"
	"slices"
	"strconv"
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

type Tables struct {
	seedToSoil, soilToFertilizer, fertilizerToWater, waterToLight, ligthToTemp, tempToHumid, humidToLocation map[uint64]uint64
}

var numRegex = regexp.MustCompile("\\d+\\s?")

func createLookupTable(inputs []string) map[uint64]uint64 {
	lut := make(map[uint64]uint64)
	var parsed []string
	var source, dest uint64
	for _, line := range inputs {
		parsed = strings.Split(line, " ")
		source, _ = strconv.ParseUint(parsed[0], 10, 64)
		dest, _ = strconv.ParseUint(parsed[1], 10, 64)
		for i := 0; i < pkg.MustAtoi(parsed[2]); i++ {
			lut[dest+uint64(i)] = source + uint64(i)
		}
	}

	return lut
}

func extractTable(input string, header LUT) map[uint64]uint64 {
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
	return createLookupTable(data)
}

func extractSeeds(input string) []uint64 {
	seedLine := strings.Split(input, "\n")[0]
	numbers := strings.Split(seedLine, " ")[1:]
	seeds := make([]uint64, len(numbers))
	for index, num := range numbers {
		seeds[index], _ = strconv.ParseUint(num, 10, 64)
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

func convertSeedToLocation(seed uint64, tables Tables) uint64 {
	finalNumber := seed
	if soil, prs := tables.seedToSoil[seed]; prs {
		finalNumber = soil
	}
	if fertilizer, prs := tables.soilToFertilizer[finalNumber]; prs {
		finalNumber = fertilizer
	}
	if water, prs := tables.fertilizerToWater[finalNumber]; prs {
		finalNumber = water
	}
	if light, prs := tables.waterToLight[finalNumber]; prs {
		finalNumber = light
	}
	if temp, prs := tables.ligthToTemp[finalNumber]; prs {
		finalNumber = temp
	}
	if humid, prs := tables.tempToHumid[finalNumber]; prs {
		finalNumber = humid
	}
	if location, prs := tables.humidToLocation[finalNumber]; prs {
		finalNumber = location
	}
	return finalNumber
}

func run(input string) (interface{}, interface{}) {
	seeds := extractSeeds(input)
	tables := generateTables(input)
	var receivedLocation uint64
	var min uint64 = math.MaxUint64
	for _, seed := range seeds {
		receivedLocation = convertSeedToLocation(seed, tables)
		if receivedLocation < min {
			min = receivedLocation
		}
	}

	return min, 0
}

func main() {
	execute.Run(run, Tests, Puzzle, true)
}
