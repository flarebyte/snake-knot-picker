// purpose: Convert command documents to and from JSON for storage and transport boundaries.
// responsibilities: Expose focused functions that parse, validate, transform, or register data within this file's module boundary.
// architecture notes: The implementation favors small deterministic helpers with explicit error IDs to keep behavior stable for both humans and automation.
package picker

import "encoding/json"

// ParseCommandDocumentJSON decodes a JSON payload into a CommandDocument.
func ParseCommandDocumentJSON(data []byte) (CommandDocument, error) {
	var doc CommandDocument
	if err := json.Unmarshal(data, &doc); err != nil {
		return CommandDocument{}, NewSchemaError(ErrorIDSchemaInvalidValue, map[string]string{"field": "json"})
	}
	return doc, nil
}

// ToJSON encodes the command document as JSON.
func (d CommandDocument) ToJSON() ([]byte, error) {
	return json.Marshal(d)
}
