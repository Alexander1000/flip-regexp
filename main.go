package flip_regexp

import (
	"math/rand"
	"strconv"
	"time"
)

const (
	WORD_DIGITS = "qwertyuiopasdfghjklzxcvbnm0123456789 "

	tokenEscape           = byte(0x5C)  // \
	tokenBracketOpen      = byte(0x5B)  // [
	tokenBracketClose     = byte(0x5D)  // ]
	tokenParenthesisOpen  = byte(0x28)  // (
	tokenParenthesisClose = byte(0x29)  // )
	tokenBraceOpen        = byte(0x7B)  // {
	tokenBraceClose       = byte(0x7D)  // }
	tokenPipe             = byte(0x7C)  // |
	tokenDot              = byte(0x2E)  // .
	tokenQuestion         = byte(0x3F)  // ?
	tokenDoubleDot        = byte(0x3A)  // :
	tokenPlus             = byte(0x2B)  // +
	tokenHyphen           = byte(0x2D)  // -
	tokenComma            = byte(0x2C)  // ,
	tokenAsterisk         = byte(0x2A)  // *

	letterDigit0 = byte(0x30) // 0
	letterDigit9 = byte(0x39) // 9

	letterLowerA = byte(0x61) // a
	letterLowerZ = byte(0x7A) // z
	letterUpperA = byte(0x41) // A
	letterUpperZ = byte(0x5A) // Z
)

type Builder struct {
	Pattern  []byte
	Position int
	Result   []byte
}

func NewBuilder(pattern []byte) *Builder {
	return &Builder{Pattern: pattern}
}

func (b *Builder) getCurrentSymbol() byte {
	return b.Pattern[b.Position]
}

func (b *Builder) randInt(min, max int) int {
	if min == max {
		return min
	}

	rand.Seed(time.Now().UnixNano())
	return min + rand.Intn(max-min)
}

func (b *Builder) Render() ([]byte, error) {
	b.Result = make([]byte, 10)
	b.Position = 0
	escape := false

	for b.Position < len(b.Pattern) {
		letter := b.getCurrentSymbol()

		if !escape && letter == tokenEscape {
			escape = true
			b.Position++
		} else if escape {
			escape = false
			b.Result = append(b.Result, letter)
			b.Position++
		} else if letter == tokenBracketOpen {
			b.Position++
			b.parseInBracket()
			continue
		} else if letter == tokenParenthesisOpen {
			b.Position++
			b.parseInBrace()
		} else {
			b.Result = append(b.Result, letter)
			b.Position++
		}
	}

	return b.Result, nil
}

func (b *Builder) getIntervalLetter(begin byte, end byte) []byte {
	var str []byte
	curSymbol := begin

	for curSymbol <= end {
		str = append(str, curSymbol)
		curSymbol++
	}

	return str
}

func (b *Builder) parseInBrace() {
	var words [][]byte
	var str []byte

	for b.Position < len(b.Pattern) {
		letter := b.getCurrentSymbol()
		b.Position++

		if b.isLetter(letter) {
			str = append(str, letter)
		} else if letter == tokenPipe {
			words = append(words, str)
			str = []byte{}
		} else if letter == tokenParenthesisClose {
			if len(str) > 0 {
				words = append(words, str)
				str = []byte{}
			}

			break
		} else {
			str = append(str, letter)
		}
	}

	if len(words) == 0 {
		return
	}

	if b.Position < len(b.Pattern) {
		letter := b.getCurrentSymbol()

		if letter == tokenQuestion {
			b.Position++
			if b.randInt(0, 2) == 1 {
				b.Result = append(b.Result, words[b.randInt(0, len(words))]...)
				return
			} else {
				return
			}
		}
	}

	b.Result = append(b.Result, words[b.randInt(0, len(words))]...)
}

// генерация строки для шаблона: []{,}
func (b *Builder) parseInBracket() {
	var abc []byte
	var prev byte
	interval := false

	for b.Position < len(b.Pattern) {
		letter := b.getCurrentSymbol()
		b.Position++

		if b.isLetter(letter) {
			if !interval {
				abc = append(abc, letter)
				prev = letter
			} else {
				abc = append(abc, b.getIntervalLetter(prev+1, letter)...)
				interval = false
			}
		} else if letter == tokenHyphen {
			if interval {
				abc = append(abc, tokenHyphen)
			} else {
				interval = true
			}
		} else if letter == tokenBracketClose {
			interval = false
			break
		}
	}

	// var strMin, strMax, str string
	var str []byte
	min := 1
	max := 1
	var size []int
	var bracket bool

	for b.Position < len(b.Pattern) {
		letter := b.getCurrentSymbol()

		if letter == tokenBraceOpen {
			bracket = true
		} else if letter == tokenBraceClose {
			bracket = false
		} else if bracket && b.isDigit(letter) {
			str = append(str, letter)
		} else if bracket && letter == tokenComma {
			tSize, _ := strconv.Atoi(string(str))
			size = append(size, tSize)
			str = []byte{}
		} else {
			break
		}

		b.Position++
	}

	if len(str) > 0 {
		tSize, _ := strconv.Atoi(string(str))
		size = append(size, tSize)
		str = []byte{}
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
	length = b.randInt(min, max)

	b.Result = append(b.Result, b.randomString(length, abc)...)
}

func (b *Builder) isDigit(letter byte) bool {
	return (letter >= letterDigit0 && letter <= letterDigit9)
}

func (b *Builder) isLetter(letter byte) bool {
	return (letter >= letterLowerA && letter <= letterLowerZ) || (letter >= letterUpperA && letter <= letterUpperZ) || b.isDigit(letter)
}

func (b *Builder) randomString(length int, abc []byte) []byte {
	var str []byte
	i := 0

	if len(abc) == 0 {
		abc = []byte(WORD_DIGITS)
	}

	size := len(abc)

	for i < length {
		num := b.randInt(0, size)
		str = append(str, abc[num])
		i++
	}

	return str
}
