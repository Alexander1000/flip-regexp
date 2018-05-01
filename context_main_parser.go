package flip_regexp

type MainContext struct {
	Builder *Builder
	TokenList []*Token
}

func NewMainContext(builder *Builder) *MainContext {
	return &MainContext{
		Builder: builder,
		TokenList: make([]*Token, 0, 10),
	}
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

	c.TokenList = append(c.TokenList, &token)
	return &token, nil
}

func (c *MainContext) render() string {
	printResult := make([]string, 0)
	result := ""
	isLetter := false
	isAlias := false
	// aliasType := byte(0)
	aliasAbc := ""
	for _, token := range c.TokenList {
		tokenStream := token.Render()

		if token.Type == typeLetter {
			result += string(tokenStream)
			isLetter = true
			isAlias = false
		} else if token.Type == typeQuantifier {
			if isLetter && len(result) > 1 {
				printResult = append(printResult, result[0:len(result)-1])
				result = result[len(result)-1:]
			}

			if isLetter {
				if tokenStream[0] == tokenAsterisk {
					value := c.Builder.randInt(0, randomMax)
					tResult := c.Builder.randomString(value, []byte(result))
					printResult = append(printResult, string(tResult))
					result = ""
				} else if tokenStream[0] == tokenPlus {
					value := c.Builder.randInt(1, randomMax)
					tResult := c.Builder.randomString(value, []byte(result))
					printResult = append(printResult, string(tResult))
					result = ""
				} else if tokenStream[0] == tokenQuestion {
					value := c.Builder.randInt(0, 1)
					if value == 1 {
						printResult = append(printResult, result)
					}
					result = ""
				}
			} else if isAlias {
				if tokenStream[0] == tokenAsterisk {
					value := c.Builder.randInt(0, randomMax)
					tResult := c.Builder.randomString(value, []byte(aliasAbc))
					printResult = append(printResult, string(tResult))
					result = ""
				} else if tokenStream[0] == tokenPlus {
					value := c.Builder.randInt(1, randomMax)
					tResult := c.Builder.randomString(value, []byte(aliasAbc))
					printResult = append(printResult, string(tResult))
					result = ""
				} else if tokenStream[0] == tokenQuestion {
					value := c.Builder.randInt(0, 1)
					if value == 1 {
						tResult := c.Builder.randomString(value, []byte(aliasAbc))
						printResult = append(printResult, string(tResult))
					}
					result = ""
				}
			}

			isLetter = false
			isAlias = false
		} else if token.Type == typeAlias {
			if isLetter {
				printResult = append(printResult, result)
				result = ""
			}

			isAlias = true

			if tokenStream[1] == aliasTokenDigit {
				// aliasType = aliasTokenDigit
				aliasAbc = "0123456789"
			} else if tokenStream[1] == aliasTokenNotDigit {
				// aliasType = aliasTokenNotDigit
				aliasAbc = "0123456789qwertyuiopasdfghjklzxcvbnm "
			} else if tokenStream[1] == aliasTokenNotSpace {
				// aliasType = aliasTokenNotSpace
				aliasAbc = "0123456789qwertyuiopasdfghjklzxcvbnm"
			} else if tokenStream[1] == aliasTokenSpace {
				// aliasType = aliasTokenSpace
				aliasAbc = " "
			} else if tokenStream[1] == aliasTokenWord {
				// aliasType = aliasTokenWord
				aliasAbc = "qwertyuiopasdfghjklzxcvbnm"
			} else if tokenStream[1] == aliasTokenNotWord {
				// aliasType = aliasTokenNotWord
				aliasAbc = "0123456789"
			}
		}
	}

	result = ""
	for _, str := range printResult {
		result += str
	}

	return result
}
