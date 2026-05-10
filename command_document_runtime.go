// purpose: Bridge document-based entry points to the compile-and-validate runtime flow.
// responsibilities: Expose focused functions that parse, validate, transform, or register data within this file's module boundary.
// architecture notes: The implementation favors small deterministic helpers with explicit error IDs to keep behavior stable for both humans and automation.
package picker

// ValidateWithDocument compiles a document and validates runtime argv against it.
func ValidateWithDocument(doc CommandDocument, argv []string) (*ParseResult, error) {
	compiled, err := CompileCommandDocument(doc)
	if err != nil {
		return nil, err
	}
	return Validate(compiled, argv)
}

// ValidateWithDocumentJSON parses a JSON document and validates runtime argv against it.
func ValidateWithDocumentJSON(data []byte, argv []string) (*ParseResult, error) {
	doc, err := ParseCommandDocumentJSON(data)
	if err != nil {
		return nil, err
	}
	return ValidateWithDocument(doc, argv)
}
