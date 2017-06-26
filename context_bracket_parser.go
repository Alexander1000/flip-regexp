package flip_regexp

type BracketContext struct {
	Builder       *Builder
	StartPosition int
	OpenBracket   bool
	OpenBrace bool
	comma bool
	size bool
	finish bool
	// Size []int
}

func (b *Builder) getBracketContext() *BracketContext {
	if b.bracketContext != nil {
		return b.bracketContext
	}

	b.bracketContext = &BracketContext{Builder: b, StartPosition: b.Position, OpenBracket: true, OpenBrace: false, finish: false}
	return b.bracketContext
}

func (bc *BracketContext) getNextToken() (*Token, error) {
	token := Token{Type: typeInvalid, Length: 0}
	token.Stream = make([]byte, 0, 1)

	if bc.finish {
		return &token, nil
	}

	curPosition := bc.Builder.Position
	escape := false
	first := curPosition == bc.StartPosition

	for curPosition < len(bc.Builder.Pattern) {
		letter := bc.Builder.getSymbolByRelativeOffset(token.Length)
		curPosition++
		token.Length++

		if bc.OpenBracket {
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
		} else {
			if bc.OpenBrace {
				if bc.Builder.isDigit(letter) {
					token.Type = typeQuantifierSize
					token.Stream = append(token.Stream, letter)
					continue
				} else if letter == tokenBraceClose {
					// token.Type = typeQuantifierClose
					// token.Stream = append(token.Stream, letter)
					bc.finish = true
					break
				} else if !bc.comma && letter == tokenComma {
					bc.comma = true
					token.Type = typeQuantifierComma
					break
				} else {
					token.Type = typeInvalid
					break
				}
			} else {
				if letter == tokenBraceOpen {
					bc.comma = false
					// bc.size = false
					bc.OpenBrace = true
					token.Type = typeQuantifierOpen
					token.Stream = append(token.Stream, letter)
					break
				} else if letter == tokenQuestion {
					token.Type = typeQuantifier
					token.Stream = append(token.Stream, letter)
					bc.finish = true
					break
				} else if letter == tokenAsterisk {
					token.Type = typeQuantifier
					token.Stream = append(token.Stream, letter)
					bc.finish = true
					break
				} else if letter == tokenPlus {
					token.Type = typeQuantifier
					token.Stream = append(token.Stream, letter)
					bc.finish = true
					break
				} else {
					bc.finish = true
					token.Type = typeInvalid
					break
				}
			}
		}
	}

	return &token, nil
}
