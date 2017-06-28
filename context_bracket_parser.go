package flip_regexp

import "strconv"

type BracketContext struct {
	Builder       *Builder
	StartPosition int
	OpenBracket   bool
	// OpenBrace bool
	// comma bool
	// size bool
	finish bool
	// Size []int
}

func (b *Builder) getBracketContext() *BracketContext {
	if b.bracketContext != nil {
		return b.bracketContext
	}

	b.bracketContext = &BracketContext{Builder: b, StartPosition: b.Position, OpenBracket: true, finish: false}
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

	isComma := false
	isOpenBrace := false

	checkInterval := false

	minSize := make([]byte, 0)
	maxSize := make([]byte, 0)

	isSize := false

	for curPosition < len(bc.Builder.Pattern) {
		letter := bc.Builder.getSymbolByRelativeOffset(token.Length)
		curPosition++
		token.Length++

		if bc.OpenBracket {
			if checkInterval {
				if letter == tokenHyphen {
					token.Stream = append(token.Stream, letter)
					token.Type = typeInterval
					checkInterval = false
					continue
				} else {
					curPosition--
					token.Length--
					break
				}
			}

			if token.Type == typeInterval {
				token.Stream = append(token.Stream, letter)
				break
			}

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
				checkInterval = true
				continue
			}
		} else {
			if isOpenBrace {
				if bc.Builder.isDigit(letter) {
					isSize = true
					// token.Type = typeQuantifierSize
					token.Stream = append(token.Stream, letter)

					if !isComma {
						minSize = append(minSize, letter)
					} else {
						maxSize = append(maxSize, letter)
					}
					continue
				} else if letter == tokenBraceClose {
					// token.Type = typeQuantifierClose
					token.Stream = append(token.Stream, letter)
					bc.finish = true
					break
				} else if !isComma && letter == tokenComma {
					isComma = true
					token.Stream = append(token.Stream, letter)
					// token.Type = typeQuantifierComma
					continue
				} else {
					token.Type = typeInvalid
					break
				}
			} else {
				if letter == tokenBraceOpen {
					// isComma = false
					// bc.size = false
					// bc.OpenBrace = true
					isOpenBrace = true
					token.Type = typeQuantifier
					token.Stream = append(token.Stream, letter)
					continue
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

	if token.Type == typeQuantifier && isSize {
		if (!isComma) {
			token.Min = 0
			token.Max, _ = strconv.Atoi(string(minSize))
		} else {
			token.Min, _ = strconv.Atoi(string(minSize))
			token.Max, _ = strconv.Atoi(string(maxSize))
		}
	}

	return &token, nil
}
