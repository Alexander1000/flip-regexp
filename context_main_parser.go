package flip_regexp

type MainContext struct {
	Builder *Builder
}

func (c *MainContext) getNextToken() (*Token, error) {
	token := Token{Type: typeInvalid, Length: 0}
	token.Stream = make([]byte, 0, 1)
	curPosition := c.Builder.Position
	escape := false

	for curPosition < len(c.Builder.Pattern) {
		letter := c.Builder.getSymbolByRelativeOffset(token.Length)
		curPosition++
		token.Length++

		if !escape && letter == tokenEscape {
			escape = true
		} else if escape {
			if c.Builder.isAlias(letter) {
				token.Type = typeAlias
			} else {
				token.Type = typeLetter
			}

			token.Stream = append(token.Stream, letter)
			break
		} else if c.Builder.isQuantifier(letter) {
			token.Type = typeQuantifier
			token.Stream = append(token.Stream, letter)
			break
		} else if letter == tokenParenthesisOpen || letter == tokenBracketOpen {
			token.Type = typeGroupOpen
			token.Stream = append(token.Stream, letter)
			break
		} else {
			token.Stream = append(token.Stream, letter)
			token.Type = typeLetter
			break
		}
	}

	return &token, nil
}
