package otari

// Endpoint-coverage manifest checks.
//
// sdk-endpoints.txt records which gateway endpoints this SDK surfaces
// ([covered]) and which it deliberately does not ([excluded]). The file is a
// generated artifact: the gateway's codegen workflow pushes it here alongside
// the generated core, from the canonical copy at
// scripts/sdk_codegen/sdk-endpoints.txt in mozilla-ai/otari.
//
// The drift gate itself lives in the gateway, where the manifest is validated
// against docs/public/openapi.json from the same commit. It used to live here
// and fetch the spec from main over the network at test time, which made the
// result depend on when the test ran rather than on what the commit contained:
// an unchanged commit passed one day and failed the next, and because CI only
// runs on push and pull_request, main sat red unnoticed for over two weeks
// (mozilla-ai/otari#438). What remains here is offline and deterministic.

import (
	"os"
	"path/filepath"
	"runtime"
	"sort"
	"strings"
	"testing"
)

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
	if len(excluded) == 0 {
		t.Fatal("manifest [excluded] section is empty")
	}
}

func TestManifestSectionsAreDisjoint(t *testing.T) {
	covered, excluded := parseManifest(t)
	var both []string
	for e := range covered {
		if excluded[e] {
			both = append(both, e)
		}
	}
	sort.Strings(both)
	if len(both) > 0 {
		t.Errorf("endpoint(s) in both [covered] and [excluded]: %v", both)
	}
}

func TestManifestEntriesAreWellFormed(t *testing.T) {
	covered, excluded := parseManifest(t)
	methods := map[string]bool{
		"GET": true, "POST": true, "PUT": true, "PATCH": true, "DELETE": true,
	}
	all := map[string]bool{}
	for _, set := range []map[string]bool{covered, excluded} {
		for e := range set {
			all[e] = true
		}
	}
	for _, entry := range sortedKeys(all) {
		method, path, found := strings.Cut(entry, " ")
		if !found {
			t.Errorf("manifest entry is not %q: %q", "METHOD /path", entry)
			continue
		}
		if !methods[method] {
			t.Errorf("manifest entry has an unknown method: %q", entry)
		}
		if !strings.HasPrefix(path, "/") {
			t.Errorf("manifest entry path does not start with %q: %q", "/", entry)
		}
	}
}
