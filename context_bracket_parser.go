package flip_regexp

type BracketContext struct {
	Builder       *Builder
	StartPosition int
	OpenBracket   bool
}

func (b *Builder) getBracketContext() *BracketContext {
	if b.bracketContext != nil {
		return b.bracketContext
	}

	b.bracketContext = &BracketContext{Builder: b, StartPosition: b.Position, OpenBracket: true}
	return b.bracketContext
}

func (bc *BracketContext) getNextToken() (*Token, error) {
	token := Token{Type: typeInvalid, Length: 0}
	token.Stream = make([]byte, 0, 1)
	curPosition := bc.Builder.Position
	escape := false
	first := curPosition == bc.StartPosition

	for curPosition < len(bc.Builder.Pattern) {
		letter := bc.Builder.getSymbolByRelativeOffset(token.Length)
		curPosition++
		token.Length++

		if !escape && letter == tokenEscape {
			escape = true
		} else if escape {
			if bc.Builder.isAlias(letter) {
				token.Type = typeAlias
			} else {
				token.Type = typeLetter
			}

			token.Stream = append(token.Stream, letter)
			break
		} else if first && letter == tokenCircumflex {
			token.Type = typeCircumflex
			token.Stream = append(token.Stream, letter)
			break
		} else if bc.OpenBracket && letter == tokenBracketClose {
			bc.OpenBracket = false
			token.Type = typeGroupClose
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
