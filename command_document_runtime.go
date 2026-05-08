package picker

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
