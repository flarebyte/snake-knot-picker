// purpose: Extract and normalize primary schema shape details needed by document validation and compilation.
// responsibilities: Expose focused functions that parse, validate, transform, or register data within this file's module boundary.
// architecture notes: The implementation favors small deterministic helpers with explicit error IDs to keep behavior stable for both humans and automation.
package picker

import "strconv"

type primarySchemaShape struct {
	Head      string
	Operator  string
	TupleSize int
}

func parsePrimarySchema(tokens []string) (primarySchemaShape, error) {
	shape := primarySchemaShape{
		Head:      tokens[0],
		Operator:  tokens[1],
		TupleSize: 0,
	}
	if shape.Operator != "tuple" {
		return shape, nil
	}

	sizeRaw := ""
	for i := 2; i < len(tokens); i++ {
		if tokens[i] == "--size" {
			if i+1 >= len(tokens) {
				return primarySchemaShape{}, NewSchemaError(ErrorIDSchemaMissingValue, map[string]string{"flag": "--size"})
			}
			sizeRaw = tokens[i+1]
			break
		}
	}
	if sizeRaw == "" {
		return primarySchemaShape{}, NewSchemaError(ErrorIDSchemaInvalidValue, map[string]string{"field": "tuple.size"})
	}
	size, err := strconv.Atoi(sizeRaw)
	if err != nil || size < 0 {
		return primarySchemaShape{}, NewSchemaError(ErrorIDSchemaInvalidValue, map[string]string{"field": "tuple.size"})
	}
	shape.TupleSize = size
	return shape, nil
}

func isSupportedKind(kind string) bool {
	switch kind {
	case "boolean", "string", "number", "tuple":
		return true
	default:
		return false
	}
}
