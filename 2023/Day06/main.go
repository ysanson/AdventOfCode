package main

import (
	"strings"

	"github.com/ysanson/AdventOfCode/pkg"
	"github.com/ysanson/AdventOfCode/pkg/execute"
)

type Race struct {
  distance, time int
}

func getDistanceTraveled(totalDuration int, buttonPressTime int) int {
	remainingTime := totalDuration - buttonPressTime
	return remainingTime * buttonPressTime
}

func extractRaces(input string) []Race {
  split := strings.Split(input, "\n")
  time := pkg.StandardizeSpaces(split[0])
  distance := pkg.StandardizeSpaces(split[1])
  timeNumbers := strings.Split(time, " ")[1:]
  distanceNumbers := strings.Split(distance, " ")[1:]
  races := make([]Race, len(timeNumbers))
  for i := 0; i<len(timeNumbers); i++ {
    races[i] = Race{
      time: pkg.MustAtoi(timeNumbers[i]),
      distance: pkg.MustAtoi(distanceNumbers[i]),
    }
  }
  return races
}

func extractLongerRace(input string) Race {
  split := strings.Split(input, "\n")
  time := strings.Fields(split[0])[1:] 
  distance := strings.Fields(split[1])[1:] 
  return Race{
    time: pkg.MustAtoi(strings.Join(time, "")),
    distance: pkg.MustAtoi(strings.Join(distance, "")),
  }
}

func simulateRace(race Race) int {
  greaterDistance := 0
  for btnPressTime := 1; btnPressTime < race.time; btnPressTime++ {
    if getDistanceTraveled(race.time, btnPressTime) > race.distance {
      greaterDistance++
    } 
  }
  return greaterDistance
}

func run(input string) (interface{}, interface{}) {
  races := extractRaces(input)
  part1 := 1
  possibilities := 0
  for _,race := range races {
    possibilities = simulateRace(race)
    if possibilities > 0 {
      part1 *= possibilities
    }
  }

  longerRace := extractLongerRace(input)
  possibilities = simulateRace(longerRace)

	return part1, possibilities
}

func main() {
	execute.Run(run, Tests, Puzzle, true)
}
