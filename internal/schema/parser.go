package schema

func ParseTokens(tokens []string) (*CommandAST, error) {
	if len(tokens) == 0 {
		return nil, nil
	}
	return &CommandAST{
		Head:   tokens[0],
		Tokens: append([]string(nil), tokens...),
	}, nil
}

