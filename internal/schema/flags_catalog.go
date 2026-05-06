package schema

var schemaFlagArity = map[string]int{
	"--required":            0,
	"--int":                 0,
	"--secure":              0,
	"--allow-query":         0,
	"--allow-alpha":         0,
	"--alphabetic":          0,
	"--hexa":                0,
	"--whitespace":          0,
	"--lowercase":           0,
	"--uppercase":           0,
	"--punctuation":         0,
	"--blank":               0,
	"--unicode-letter":      0,
	"--unicode-number":      0,
	"--unicode-punctuation": 0,
	"--unicode-symbol":      0,
	"--unicode-separator":   0,
	"--latin":               0,
	"--han":                 0,
	"--devanagari":          0,
	"--arabic":              0,
	"--hiragana":            0,
	"--katakana":            0,
	"--hangul":              0,
	"--tamil":               0,
	"--gujarati":            0,
	"--ethiopic":            0,
	"--email":               0,
	"--datetime":            0,
	"--duration":            0,
	"--base64":              0,
	"--boolean":             0,
	"--color":               0,
	"--date":                0,
	"--time":                0,
	"--uri":                 0,
	"--arn":                 0,
	"--allow-timezone":      0,
	"--enum":                1,
	"--enum-separator":      1,
	"--size":                1,
	"--tuple":               1,
	"--min-length":          1,
	"--max-length":          1,
	"--scheme":              1,
	"--allow-domains":       1,
	"--allow-partition":     1,
	"--allow-service":       1,
	"--allow-region":        1,
	"--allow-account-id":    1,
	"--allow-resource":      1,
	"--country":             1,
	"--layout":              1,
	"--location":            1,
	"--min-duration":        1,
	"--max-duration":        1,
	"--starts-with":         1,
	"--format":              1,
	"--min":                 1,
	"--max":                 1,
	"--multiple-of":         1,
	"--codepoint-range":     2,
}

func allowedFlagsForHead(head string) map[string]struct{} {
	if head == "custom" {
		return map[string]struct{}{
			"--required": {},
			"--country":  {},
		}
	}
	base := make(map[string]struct{}, len(schemaFlagArity))
	for flag := range schemaFlagArity {
		base[flag] = struct{}{}
	}
	return base
}
