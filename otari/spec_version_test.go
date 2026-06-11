package otari_test

import (
	"testing"

	"github.com/mozilla-ai/otari-sdk-go/otari"
	"github.com/mozilla-ai/otari-sdk-go/otari/client"
)

// The gateway codegen stamps the spec version into the generated core
// (client.SpecVersion); the shell re-exports it as otari.SpecVersion so callers
// can tell which gateway spec this SDK targets. This guards the wiring (a stale
// literal in the shell would diverge from the generated marker).
func TestSpecVersionIsSurfaced(t *testing.T) {
	if otari.SpecVersion != client.SpecVersion {
		t.Fatalf("otari.SpecVersion = %q, want %q", otari.SpecVersion, client.SpecVersion)
	}
	if otari.SpecVersion == "" {
		t.Fatal("otari.SpecVersion is empty")
	}
}
