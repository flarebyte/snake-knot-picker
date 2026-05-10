// purpose: Provide production logic for the snake-knot-picker validation and schema pipeline.
// responsibilities: Expose focused functions that parse, validate, transform, or register data within this file's module boundary.
// architecture notes: The implementation favors small deterministic helpers with explicit error IDs to keep behavior stable for both humans and automation.
package picker

import (
	"encoding/json"
	"os"
	"testing"
)

func TestSecurityRuntimeFixtures(t *testing.T) {
	raw, err := os.ReadFile("testdata/fixtures/security/runtime-security.json")
	if err != nil {
		t.Fatalf("read security fixture: %v", err)
	}
	var cases []struct {
		Name        string   `json:"name"`
		Argv        []string `json:"argv"`
		WantErrorID string   `json:"want_error_id"`
	}
	if err := json.Unmarshal(raw, &cases); err != nil {
		t.Fatalf("decode security fixture: %v", err)
	}

	docRaw, err := os.ReadFile("doc/design-meta/examples/args-command.json")
	if err != nil {
		t.Fatalf("read args fixture: %v", err)
	}

	for _, tc := range cases {
		tc := tc
		t.Run(tc.Name, func(t *testing.T) {
			_, err := ValidateWithDocumentJSON(docRaw, tc.Argv)
			assertDocErrID(t, err, tc.WantErrorID)
		})
	}
}
