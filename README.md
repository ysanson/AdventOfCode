# AdventOfCode
My contributions to the advent of code puzzles.

## Usage

Put your adventofcode.com session cookie in `.env` as `AOC_SESSION=<cookie>`, or export it. Then, from the repository root:

```sh
go run ./cmd/aoc new                # scaffold today's puzzle and download its input
go run ./cmd/aoc new -y 2025 -d 10  # scaffold a specific day
go run ./cmd/aoc fetch              # regenerate any missing puzzle.go files
go run ./2025/Day10                 # run a solution
go test ./2025/Day10                # check its answers
go test -bench=. ./2025/Day10       # benchmark it
```

`new` never overwrites existing files. The first code block on the puzzle page becomes `Test1`, so check that it is the right example. Fill in the expected answers in `puzzle.go`; a value left as `""` is not checked.

Puzzle inputs live in `puzzle.go` and are gitignored.
