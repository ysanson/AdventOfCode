package main

import (
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"sync/atomic"
	"testing"
	"time"
)

func setup(t *testing.T) *atomic.Int32 {
	t.Helper()
	t.Chdir(t.TempDir())
	t.Setenv("AOC_SESSION", "secret")
	var requests atomic.Int32
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requests.Add(1)
		if c, err := r.Cookie("session"); err != nil || c.Value != "secret" {
			t.Errorf("bad session cookie: %v", err)
		}
		if ua := r.UserAgent(); ua != userAgent {
			t.Errorf("User-Agent = %q", ua)
		}
		if strings.HasSuffix(r.URL.Path, "/input") {
			w.Write([]byte("1,2\n3,4\n"))
			return
		}
		w.Write([]byte("<pre><code>a &lt; <em>b</em>\nc\n</code></pre><pre><code>no</code></pre>"))
	}))
	t.Cleanup(server.Close)
	oldURL, oldNow := baseURL, now
	baseURL = server.URL
	now = func() time.Time { return time.Date(2025, time.December, 5, 12, 0, 0, 0, aocZone) }
	t.Cleanup(func() { baseURL, now = oldURL, oldNow })
	return &requests
}

func TestNewScaffoldsDay(t *testing.T) {
	setup(t)
	if err := run([]string{"new", "-y", "2024", "-d", "25"}); err != nil {
		t.Fatal(err)
	}
	dir := filepath.Join("2024", "Day25")
	puzzle, err := os.ReadFile(filepath.Join(dir, "puzzle.go"))
	if err != nil {
		t.Fatal(err)
	}
	for _, want := range []string{"const Puzzle = `1,2\n3,4`", "const Test1 = `a < b\nc`"} {
		if !strings.Contains(string(puzzle), want) {
			t.Errorf("puzzle.go missing %q:\n%s", want, puzzle)
		}
	}
	for _, f := range []string{"main.go", "main_test.go"} {
		if _, err := os.Stat(filepath.Join(dir, f)); err != nil {
			t.Error(err)
		}
	}
}

func TestNewDefaultsToToday(t *testing.T) {
	setup(t)
	if err := run([]string{"new"}); err != nil {
		t.Fatal(err)
	}
	if _, err := os.Stat(filepath.Join("2025", "Day05", "puzzle.go")); err != nil {
		t.Error(err)
	}
}

func TestNewRefusesToOverwrite(t *testing.T) {
	requests := setup(t)
	dir := filepath.Join("2024", "Day25")
	if err := os.MkdirAll(dir, 0o755); err != nil {
		t.Fatal(err)
	}
	path := filepath.Join(dir, "main.go")
	if err := os.WriteFile(path, []byte("mine"), 0o644); err != nil {
		t.Fatal(err)
	}
	if err := run([]string{"new", "-y", "2024", "-d", "25"}); err == nil {
		t.Fatal("expected an error")
	}
	if got, _ := os.ReadFile(path); string(got) != "mine" {
		t.Errorf("file was overwritten: %q", got)
	}
	if requests.Load() != 0 {
		t.Errorf("made %d requests", requests.Load())
	}
}

func TestNewRefusesLockedPuzzle(t *testing.T) {
	requests := setup(t)
	if err := run([]string{"new", "-y", "2025", "-d", "6"}); err == nil || !strings.Contains(err.Error(), "unlocked") {
		t.Fatalf("err = %v", err)
	}
	if requests.Load() != 0 {
		t.Errorf("made %d requests", requests.Load())
	}
}

func TestFetchOnlyFillsMissing(t *testing.T) {
	setup(t)
	missing, present := filepath.Join("2024", "Day11"), filepath.Join("2024", "Day12")
	for _, d := range []string{missing, present} {
		if err := os.MkdirAll(d, 0o755); err != nil {
			t.Fatal(err)
		}
	}
	existing := filepath.Join(present, "puzzle.go")
	if err := os.WriteFile(existing, []byte("mine"), 0o644); err != nil {
		t.Fatal(err)
	}
	if err := run([]string{"fetch"}); err != nil {
		t.Fatal(err)
	}
	if _, err := os.Stat(filepath.Join(missing, "puzzle.go")); err != nil {
		t.Error(err)
	}
	if got, _ := os.ReadFile(existing); string(got) != "mine" {
		t.Errorf("file was overwritten: %q", got)
	}
}

func TestGoLiteralHandlesBackticks(t *testing.T) {
	if got := goLiteral("a`b"); got != `"a`+"`"+`b"` {
		t.Errorf("got %s", got)
	}
}
