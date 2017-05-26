package flip_regexp

import (
	"strconv"
	"math/rand"
)

const (
	WORD_DIGITS = "qwertyuiopasdfghjklzxcvbnm0123456789 "

	TOKEN_BRACKET_OPEN    = `[`
	TOKEN_BRACKET_CLOSE   = `]`
	TOKEN_C_BRACKET_OPEN  = `(`
	TOKEN_C_BRACKET_CLOSE = `)`
	TOKEN_PIPE            = `|`
	TOKEN_QUESTION        = `?`
	TOKEN_PLUS            = `+`
	TOKEN_MINUS           = `-`
	TOKEN_ESCAPE          = `\`
)

type Builder struct {
	Pattern  string
	Position int
	Result   string
}

func NewBuilder(pattern string) *Builder {
	return &Builder{Pattern: pattern}
}

func (b *Builder) getCurrentSymbol() byte {
	return b.Pattern[b.Position]
}

func (b *Builder) Render() (string, error) {
	b.Result = ""
	b.Position = 0
	escape := false

	for b.Position < len(b.Pattern) {
		letter := b.getCurrentSymbol()

		if !escape && letter == []byte(TOKEN_ESCAPE)[0] {
			escape = true
			b.Position++
		} else if escape {
			escape = false
			b.Result += string(letter)
			b.Position++
		} else if letter == []byte(TOKEN_BRACKET_OPEN)[0] {
			b.Position++
			b.parseInBracket()
			continue
		} else if letter == []byte(TOKEN_C_BRACKET_OPEN)[0] {
			b.Position++
			b.parseInCBracket()
		} else {
			b.Result += string(letter)
			b.Position++
		}
	}

	return b.Result, nil
}

func (b *Builder) getIntervalLetter(begin byte, end byte) string {
	str := ""
	curSymbol := begin

	for curSymbol <= end {
		str += string(curSymbol)
		curSymbol++
	}

	return str
}

func (b *Builder) parseInCBracket() {
	var words []string
	str := ""

	for b.Position < len(b.Pattern) {
		letter := b.getCurrentSymbol()
		b.Position++

		if b.isLetter(letter) {
			str += string(letter)
		} else if letter == []byte(TOKEN_PIPE)[0] {
			words = append(words, str)
			str = ""
		} else if letter == []byte(TOKEN_C_BRACKET_CLOSE)[0] {
			if len(str) > 0 {
				words = append(words, str)
				str = ""
			}

			break
		} else {
			str += string(letter)
		}
	}

	if len(words) == 0 {
		return
	}

	if b.Position < len(b.Pattern) {
		letter := b.getCurrentSymbol()

		if letter == []byte(TOKEN_QUESTION)[0] {
			b.Position++
			if rand.Intn(2) == 1 {
				b.Result += words[rand.Intn(len(words))]
				return
			} else {
				return
			}
		}
	}

	b.Result += words[rand.Intn(len(words))]
}

// генерация строки для шаблона: []{,}
func (b *Builder) parseInBracket() {
	alfa := ""
	var prev byte
	interval := false

	for b.Position < len(b.Pattern) {
		letter := b.getCurrentSymbol()
		b.Position++

		if b.isLetter(letter) {
			if !interval {
				alfa += string(letter)
				prev = letter
			} else {
				// prev = letter
				alfa += b.getIntervalLetter(prev+1, letter)
				interval = false
			}
		} else if letter == []byte(TOKEN_MINUS)[0] {
			if interval {
				alfa += "-"
			} else {
				interval = true
			}
		} else if letter == []byte(TOKEN_BRACKET_CLOSE)[0] {
			interval = false
			break
		}
	}

	// var strMin, strMax, str string
	var str string
	min := 1
	max := 1
	var size []int
	var bracket bool

	for b.Position < len(b.Pattern) {
		letter := b.getCurrentSymbol()

		if letter == []byte(`{`)[0] {
			bracket = true
		} else if letter == []byte(`}`)[0] {
			bracket = false
		} else if bracket && b.isDigit(letter) {
			str += string(letter)
		} else if bracket && letter == []byte(`,`)[0] {
			tSize, _ := strconv.Atoi(str)
			size = append(size, tSize)
			str = ""
		} else {
			break
		}

		b.Position++
	}

	if len(str) > 0 {
		tSize, _ := strconv.Atoi(str)
		size = append(size, tSize)
		str = ""
	}

	if len(size) == 1 {
		min = size[0]
		max = size[0]
	} else if len(size) == 2 {
		min = size[0]
		max = size[1]
	} else {
		min = 1
		max = 1
	}

	var length int

	if min != max {
		length = min + rand.Intn(max-min)
	} else {
		length = min
	}

	b.Result += b.randomString(length, alfa)
}

func (b *Builder) isDigit(letter byte) bool {
	return (letter >= []byte(`0`)[0] && letter <= []byte(`9`)[0])
}

func (b *Builder) isLetter(letter byte) bool {
	return (letter >= []byte(`a`)[0] && letter <= []byte(`z`)[0]) || (letter >= []byte(`A`)[0] && letter <= []byte(`Z`)[0]) || b.isDigit(letter)
}

func (b *Builder) randomString(length int, alpha string) string {
	str := ""
	i := 0

	if alpha == "" {
		alpha = WORD_DIGITS
	}

	size := len(alpha)

	for i < length {
		num := rand.Intn(size)
		str += string(alpha[num])
		i++
	}

	return str
}

