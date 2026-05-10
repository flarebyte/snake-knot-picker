// purpose: Define argv-facing value container types used by argv parser integration paths.
// responsibilities: Expose focused functions that parse, validate, transform, or register data within this file's module boundary.
// architecture notes: The implementation favors small deterministic helpers with explicit error IDs to keep behavior stable for both humans and automation.
package argv

import "github.com/flarebyte/snake-knot-picker"

// Values is the argv-layer map of parsed values keyed by flag name.
type Values map[string]picker.Value
