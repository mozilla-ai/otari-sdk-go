package otari

import "github.com/mozilla-ai/otari-sdk-go/otari/client"

// SpecVersion is the gateway/spec version the generated core was built from,
// stamped into the core by the gateway codegen pipeline. It lets callers tell
// which gateway spec this SDK targets.
const SpecVersion = client.SpecVersion
