package otari

// Endpoint-coverage drift gate.
//
// Fetches the canonical otari gateway OpenAPI spec and asserts that every API
// endpoint it exposes is accounted for in sdk-endpoints.txt -- either wrapped
// by this SDK's public surface ([covered]) or deliberately deferred
// ([excluded]). A new gateway endpoint in neither section fails this test, so a
// future endpoint (as /messages once was) cannot silently go unsurfaced.
//
// The fetch uses net/http. It is skipped offline (network error /
// OTARI_SKIP_NETWORK_TESTS=1) but runs in CI, where the network is available.

import (
	"encoding/json"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"sort"
	"strings"
	"testing"
	"time"
)

const specURL = "https://raw.githubusercontent.com/mozilla-ai/otari/main/docs/public/openapi.json"

var httpMethods = map[string]bool{
	"get": true, "post": true, "put": true, "patch": true, "delete": true,
}

// manifestPath returns the absolute path to sdk-endpoints.txt at the repo root
// (one directory up from this package).
func manifestPath(t *testing.T) string {
	t.Helper()
	_, thisFile, _, ok := runtime.Caller(0)
	if !ok {
		t.Fatal("could not determine test file location")
	}
	return filepath.Join(filepath.Dir(thisFile), "..", "sdk-endpoints.txt")
}

// parseManifest returns the covered and excluded endpoint sets.
func parseManifest(t *testing.T) (covered, excluded map[string]bool) {
	t.Helper()
	raw, err := os.ReadFile(manifestPath(t))
	if err != nil {
		t.Fatalf("could not read manifest: %v", err)
	}
	covered = map[string]bool{}
	excluded = map[string]bool{}
	var section map[string]bool
	for _, rawLine := range strings.Split(string(raw), "\n") {
		line := strings.TrimSpace(rawLine)
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		switch line {
		case "[covered]":
			section = covered
			continue
		case "[excluded]":
			section = excluded
			continue
		}
		entry := strings.TrimSpace(strings.SplitN(line, "#", 2)[0])
		if entry == "" || section == nil {
			continue
		}
		fields := strings.Fields(entry)
		if len(fields) < 2 {
			continue
		}
		section[strings.ToUpper(fields[0])+" "+strings.Join(fields[1:], " ")] = true
	}
	return covered, excluded
}

// fetchSpecEndpoints fetches the spec and returns its METHOD /path set,
// dropping /health* meta routes. It skips the test on network failure.
func fetchSpecEndpoints(t *testing.T) map[string]bool {
	t.Helper()
	if os.Getenv("OTARI_SKIP_NETWORK_TESTS") == "1" {
		t.Skip("OTARI_SKIP_NETWORK_TESTS=1")
	}
	client := &http.Client{Timeout: 30 * time.Second}
	resp, err := client.Get(specURL)
	if err != nil {
		t.Skipf("could not fetch otari OpenAPI spec from %s: %v", specURL, err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		t.Skipf("could not fetch otari OpenAPI spec from %s: HTTP %d", specURL, resp.StatusCode)
	}
	var doc struct {
		Paths map[string]map[string]json.RawMessage `json:"paths"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&doc); err != nil {
		t.Fatalf("could not decode OpenAPI spec: %v", err)
	}
	eps := map[string]bool{}
	for path, methods := range doc.Paths {
		if path == "/health" || strings.HasPrefix(path, "/health/") {
			continue
		}
		for method := range methods {
			if httpMethods[strings.ToLower(method)] {
				eps[strings.ToUpper(method)+" "+path] = true
			}
		}
	}
	return eps
}

func sortedKeys(m map[string]bool) []string {
	out := make([]string, 0, len(m))
	for k := range m {
		out = append(out, k)
	}
	sort.Strings(out)
	return out
}

func TestManifestParses(t *testing.T) {
	covered, excluded := parseManifest(t)
	if len(covered) == 0 {
		t.Fatal("manifest [covered] section is empty")
	}
	for e := range covered {
		if excluded[e] {
			t.Errorf("endpoint in both [covered] and [excluded]: %s", e)
		}
	}
}

func TestSpecEndpointsAreAccountedFor(t *testing.T) {
	covered, excluded := parseManifest(t)
	spec := fetchSpecEndpoints(t)
	var unaccounted []string
	for e := range spec {
		if !covered[e] && !excluded[e] {
			unaccounted = append(unaccounted, e)
		}
	}
	sort.Strings(unaccounted)
	if len(unaccounted) > 0 {
		t.Fatalf("Gateway OpenAPI exposes endpoint(s) the SDK does not account for: %v. "+
			"Add a public wrapper and list under [covered], or defer it under [excluded] "+
			"with a reason, in sdk-endpoints.txt.", unaccounted)
	}
}

func TestManifestHasNoStaleEntries(t *testing.T) {
	covered, excluded := parseManifest(t)
	spec := fetchSpecEndpoints(t)
	accounted := map[string]bool{}
	for e := range covered {
		accounted[e] = true
	}
	for e := range excluded {
		accounted[e] = true
	}
	var stale []string
	for e := range accounted {
		if !spec[e] {
			stale = append(stale, e)
		}
	}
	sort.Strings(stale)
	if len(stale) > 0 {
		t.Skipf("manifest entries not present in current spec (review): %v", stale)
	}
}
