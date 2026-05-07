package picker

import "strings"

func ValidateWithDocument(doc CommandDocument, argv []string) (*ParseResult, error) {
	compiled, err := CompileCommandDocument(doc)
	if err != nil {
		return nil, err
	}
	return Validate(compiled, argv)
}

func ValidateWithDocumentJSON(data []byte, argv []string) (*ParseResult, error) {
	doc, err := ParseCommandDocumentJSON(data)
	if err != nil {
		return nil, err
	}
	return ValidateWithDocument(doc, argv)
}

// ValidateWithDocumentJSONArgvString is a convenience wrapper for callers that
// only have a single argv-like string. It uses strings.Fields tokenization and
// is therefore best-effort only. The primary API contract is []string argv.
func ValidateWithDocumentJSONArgvString(data []byte, argvLine string) (*ParseResult, error) {
	return ValidateWithDocumentJSON(data, strings.Fields(argvLine))
}
