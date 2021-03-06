package flip_regexp

const (
	contextMain    = 0
	contextBracket = 1

	tokenEscape           = byte(0x5C) // \
	tokenBracketOpen      = byte(0x5B) // [
	tokenBracketClose     = byte(0x5D) // ]
	tokenParenthesisOpen  = byte(0x28) // (
	tokenParenthesisClose = byte(0x29) // )
	tokenBraceOpen        = byte(0x7B) // {
	tokenBraceClose       = byte(0x7D) // }
	tokenPipe             = byte(0x7C) // |
	tokenDot              = byte(0x2E) // .
	tokenQuestion         = byte(0x3F) // ?
	tokenDoubleDot        = byte(0x3A) // :
	tokenPlus             = byte(0x2B) // +
	tokenHyphen           = byte(0x2D) // -
	tokenComma            = byte(0x2C) // ,
	tokenAsterisk         = byte(0x2A) // *
	tokenCircumflex       = byte(0x5E) // ^

	aliasTokenDigit = byte(0x64) // d
	aliasTokenWord  = byte(0x77) // w
	aliasTokenSpace = byte(0x73) // s

	aliasTokenNotDigit = byte(0x44) // D
	aliasTokenNotWord  = byte(0x57) // W
	aliasTokenNotSpace = byte(0x53) // S

	typeInvalid    = 0
	typeLetter     = 1
	typeGroupOpen  = 2
	typeGroupClose = 3
	typeQuantifier = 4
	typeAlias      = 5
	typeCircumflex = 6
	typeInterval   = 7
	typeParenthesisOpen = 8
	typeBracketOpen = 9
	typeBraceOpen = 10
	// {n,m}
	typeQuantifierOpen = 8
	typeQuantifierClose = 9
	typeQuantifierSize = 10
	typeQuantifierComma = 11
)

type Token struct {
	Length int
	Stream []byte
	Type   int
	Min int
	Max int
}

func (t *Token) Render() []byte {
	return t.Stream
}

type Context interface {
	getNextToken() (*Token, error)
	render() string
}

func (b *Builder) getNextToken() (*Token, error) {
	if b.ContextParser == contextMain {
		context := &MainContext{Builder: b}
		return context.getNextToken()
	} else if b.ContextParser == contextBracket {
		return b.getBracketContext().getNextToken()
	}

	return nil, nil
}

func (b *Builder) isAlias(letter byte) bool {
	return letter == aliasTokenDigit ||
		letter == aliasTokenNotDigit ||
		letter == aliasTokenWord ||
		letter == aliasTokenNotWord ||
		letter == aliasTokenSpace ||
		letter == aliasTokenNotSpace
}

func (b *Builder) getSymbolByRelativeOffset(position int) byte {
	return b.Pattern[b.Position+position]
}

func (b *Builder) getCurrentSymbol() byte {
	return b.Pattern[b.Position]
}

func (b *Builder) isDigit(letter byte) bool {
	return (letter >= letterDigit0 && letter <= letterDigit9)
}

func (b *Builder) isLetter(letter byte) bool {
	return (letter >= letterLowerA && letter <= letterLowerZ) || (letter >= letterUpperA && letter <= letterUpperZ) || b.isDigit(letter)
}

func (b *Builder) isQuantifier(letter byte) bool {
	return letter == tokenQuestion || letter == tokenPlus || letter == tokenAsterisk
}
