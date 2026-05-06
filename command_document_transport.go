package picker

import "encoding/json"

func ParseCommandDocumentJSON(data []byte) (CommandDocument, error) {
	var doc CommandDocument
	if err := json.Unmarshal(data, &doc); err != nil {
		return CommandDocument{}, NewSchemaError(ErrorIDSchemaInvalidValue, map[string]string{"field": "json"})
	}
	return doc, nil
}

func (d CommandDocument) ToJSON() ([]byte, error) {
	return json.Marshal(d)
}
