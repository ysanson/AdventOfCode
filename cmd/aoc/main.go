// Command aoc scaffolds Advent of Code days and downloads their inputs.
//
//	go run ./cmd/aoc new [-y YEAR] [-d DAY]
//	go run ./cmd/aoc fetch [-y YEAR]
//
// Run it from the repository root. The session cookie is read from the
// AOC_SESSION environment variable, or from a .env file.
package main

import (
	"bufio"
	"bytes"
	"embed"
	"errors"
	"flag"
	"fmt"
	"html"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"text/template"
	"time"
)

const userAgent = "github.com/ysanson/AdventOfCode cmd/aoc"

//go:embed templates/*.tmpl
var templates embed.FS

var (
	baseURL = "https://adventofcode.com"
	now     = time.Now
	// AoC puzzles unlock at midnight UTC-5.
	aocZone = time.FixedZone("UTC-5", -5*60*60)

	exampleRe = regexp.MustCompile(`(?s)<pre><code>(.*?)</code></pre>`)
	tagRe     = regexp.MustCompile(`<[^>]+>`)
	dayDirRe  = regexp.MustCompile(`^Day(\d{2})$`)
	yearDirRe = regexp.MustCompile(`^\d{4}$`)
)

func main() {
	if err := run(os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, "aoc:", err)
		os.Exit(1)
	}
}

func run(args []string) error {
	if len(args) == 0 {
		return errors.New("usage: aoc new [-y YEAR] [-d DAY] | aoc fetch [-y YEAR]")
	}
	flags := flag.NewFlagSet("aoc "+args[0], flag.ContinueOnError)
	year := flags.Int("y", 0, "year (defaults to the current year)")
	switch args[0] {
	case "new":
		day := flags.Int("d", 0, "day (defaults to today during December)")
		if err := flags.Parse(args[1:]); err != nil {
			return err
		}
		y, d, err := resolveDate(*year, *day)
		if err != nil {
			return err
		}
		return newDay(y, d)
	case "fetch":
		if err := flags.Parse(args[1:]); err != nil {
			return err
		}
		return fetchMissing(*year)
	}
	return fmt.Errorf("unknown command %q", args[0])
}

// resolveDate fills in defaults and rejects puzzles that have not unlocked.
func resolveDate(year, day int) (int, int, error) {
	current := now().In(aocZone)
	if year == 0 {
		year = current.Year()
	}
	if day == 0 {
		if current.Month() != time.December || current.Year() != year {
			return 0, 0, errors.New("no puzzle today: pass -d DAY")
		}
		day = current.Day()
	}
	if day < 1 || day > 25 {
		return 0, 0, fmt.Errorf("day %d is out of range", day)
	}
	return year, day, checkUnlocked(year, day)
}

func checkUnlocked(year, day int) error {
	unlock := time.Date(year, time.December, day, 0, 0, 0, 0, aocZone)
	if now().Before(unlock) {
		return fmt.Errorf("%d day %d has not unlocked yet (unlocks %s)", year, day, unlock.Format(time.RFC1123Z))
	}
	return nil
}

func dayDir(year, day int) string {
	return filepath.Join(strconv.Itoa(year), fmt.Sprintf("Day%02d", day))
}

func newDay(year, day int) error {
	dir := dayDir(year, day)
	files := []string{"main.go", "main_test.go", "puzzle.go"}
	for _, f := range files {
		path := filepath.Join(dir, f)
		if _, err := os.Stat(path); err == nil {
			return fmt.Errorf("%s already exists, not overwriting (use fetch to restore a missing puzzle.go)", path)
		}
	}
	puzzle, err := renderPuzzle(year, day)
	if err != nil {
		return err
	}
	if err := os.MkdirAll(dir, 0o755); err != nil {
		return err
	}
	for _, f := range files[:2] {
		content, err := templates.ReadFile("templates/" + f + ".tmpl")
		if err != nil {
			return err
		}
		if err := writeNew(filepath.Join(dir, f), content); err != nil {
			return err
		}
	}
	if err := writeNew(filepath.Join(dir, "puzzle.go"), puzzle); err != nil {
		return err
	}
	fmt.Println("created", dir)
	return nil
}

// fetchMissing regenerates puzzle.go for existing day directories lacking one.
func fetchMissing(onlyYear int) error {
	years, err := os.ReadDir(".")
	if err != nil {
		return err
	}
	missing := 0
	for _, y := range years {
		if !y.IsDir() || !yearDirRe.MatchString(y.Name()) {
			continue
		}
		year, _ := strconv.Atoi(y.Name())
		if onlyYear != 0 && year != onlyYear {
			continue
		}
		days, err := os.ReadDir(y.Name())
		if err != nil {
			return err
		}
		sort.Slice(days, func(i, j int) bool { return days[i].Name() < days[j].Name() })
		for _, d := range days {
			match := dayDirRe.FindStringSubmatch(d.Name())
			if !d.IsDir() || match == nil {
				continue
			}
			path := filepath.Join(y.Name(), d.Name(), "puzzle.go")
			if _, err := os.Stat(path); err == nil {
				continue
			}
			day, _ := strconv.Atoi(match[1])
			puzzle, err := renderPuzzle(year, day)
			if err != nil {
				return fmt.Errorf("%s: %w", path, err)
			}
			if err := writeNew(path, puzzle); err != nil {
				return err
			}
			fmt.Println("created", path)
			missing++
		}
	}
	if missing == 0 {
		fmt.Println("nothing to fetch")
	}
	return nil
}

func writeNew(path string, content []byte) error {
	f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_EXCL, 0o644)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = f.Write(content)
	return err
}

func renderPuzzle(year, day int) ([]byte, error) {
	session, err := loadSession()
	if err != nil {
		return nil, err
	}
	input, err := get(fmt.Sprintf("%s/%d/day/%d/input", baseURL, year, day), session)
	if err != nil {
		return nil, err
	}
	example := ""
	page, err := get(fmt.Sprintf("%s/%d/day/%d", baseURL, year, day), session)
	if err != nil {
		fmt.Fprintln(os.Stderr, "warning: no example extracted:", err)
	} else {
		example = extractExample(page)
	}

	tmpl, err := template.ParseFS(templates, "templates/puzzle.go.tmpl")
	if err != nil {
		return nil, err
	}
	var out bytes.Buffer
	err = tmpl.Execute(&out, map[string]string{
		"Example": goLiteral(example),
		"Puzzle":  goLiteral(strings.TrimRight(input, "\r\n")),
	})
	return out.Bytes(), err
}

func extractExample(page string) string {
	match := exampleRe.FindStringSubmatch(page)
	if match == nil {
		return ""
	}
	return strings.TrimRight(html.UnescapeString(tagRe.ReplaceAllString(match[1], "")), "\r\n")
}

// goLiteral prefers a readable raw string, which cannot contain a backtick.
func goLiteral(s string) string {
	if strings.ContainsAny(s, "`\r") {
		return strconv.Quote(s)
	}
	return "`" + s + "`"
}

func loadSession() (string, error) {
	if session := strings.TrimSpace(os.Getenv("AOC_SESSION")); session != "" {
		return session, nil
	}
	missing := errors.New("AOC_SESSION is not set: export it or add AOC_SESSION=<cookie> to .env")
	f, err := os.Open(".env")
	if errors.Is(err, os.ErrNotExist) {
		return "", missing
	}
	if err != nil {
		return "", err
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		key, value, ok := strings.Cut(strings.TrimSpace(scanner.Text()), "=")
		if ok && strings.TrimSpace(strings.TrimPrefix(key, "export ")) == "AOC_SESSION" {
			if value = strings.Trim(strings.TrimSpace(value), `"'`); value != "" {
				return value, nil
			}
		}
	}
	if err := scanner.Err(); err != nil {
		return "", err
	}
	return "", missing
}

func get(url, session string) (string, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return "", err
	}
	req.Header.Set("User-Agent", userAgent)
	req.AddCookie(&http.Cookie{Name: "session", Value: session})
	client := http.Client{Timeout: 30 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("GET %s: %s (is the session cookie still valid?)", url, resp.Status)
	}
	return string(body), nil
}
