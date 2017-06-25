package flip_regexp

import (
	"math/rand"
	"time"
)

const (
	letterDigit0 = byte(0x30) // 0
	letterDigit9 = byte(0x39) // 9

	letterLowerA = byte(0x61) // a
	letterLowerZ = byte(0x7A) // z
	letterUpperA = byte(0x41) // A
	letterUpperZ = byte(0x5A) // Z

	letterMinChar = byte(0x20)
	letterMaxChar = byte(0x7E)

	randomMax = 16
)

type Builder struct {
	Pattern       []byte
	Position      int
	Result        []byte
	ContextParser int
}

func NewBuilder(pattern []byte) *Builder {
	return &Builder{Pattern: pattern, Position: 0, ContextParser: contextMain}
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
	escape := false
	var prev byte
	// var curToken, prevToken *Token
	// var err error

	for b.Position < len(b.Pattern) {
		/*
			curToken, err = b.getNextToken()

			if err != nil {
				return nil, err
			}

			if curToken.Type == typeLetter {
				if prevToken != nil {
					b.Result = append(b.Result, prevToken.Stream...)
				}
				prevToken = curToken
				b.Position += curToken.Length
			} else if curToken.Type == typeQuantifier {
				// todo quantifier logics
			} else if curToken.Type == typeBracketOpen {
				// todo bracket open
			}
		*/

		letter := b.getCurrentSymbol()

		if !escape && letter == tokenEscape {
			escape = true
			b.Position++
		} else if escape {
			escape = false
			if prev != 0 {
				b.Result = append(b.Result, prev)
			}
			prev = letter
			b.Position++
		} else if letter == tokenQuestion {
			if b.randInt(0, 2) == 1 {
				b.Result = append(b.Result, prev)
			}
			b.Position++
			prev = 0
		} else if letter == tokenAsterisk {
			// todo: it is need?
			if prev != 0 {
				length := b.randInt(0, randomMax)
				if length > 0 {
					b.Result = append(b.Result, b.repeat(prev, length)...)
				}
			}
			prev = 0
			b.Position++
		} else if letter == tokenPlus {
			// todo: it is need?
			if prev != 0 {
				b.Result = append(b.Result, b.repeat(prev, b.randInt(1, randomMax))...)
			}
			prev = 0
			b.Position++
		} else {
			if prev != 0 {
				b.Result = append(b.Result, prev)
				prev = 0
			}

			if letter == tokenBracketOpen {
				b.Position++
				b.parseInBracket()
			} else if letter == tokenParenthesisOpen {
				b.Position++
				b.parseInBrace()
			} else {
				prev = letter
				b.Position++
			}
		}
	}

	if prev != 0 {
		b.Result = append(b.Result, prev)
	}

	return b.Result, nil
}

// todo: need make optimisation
func (b *Builder) repeat(char byte, length int) []byte {
	buffer := make([]byte, length)
	i := 0

	for i < length {
		buffer = append(buffer, char)
		i++
	}

	return buffer
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

func (b *Builder) randomString(length int, abc []byte) []byte {
	var str []byte
	i := 0

	if len(abc) == 0 {
		abc = append(abc, b.getIntervalLetter(letterLowerA, letterLowerZ)...)
		abc = append(abc, b.getIntervalLetter(letterUpperA, letterUpperZ)...)
		abc = append(abc, b.getIntervalLetter(letterDigit0, letterDigit9)...)
	}

	size := len(abc)

	for i < length {
		num := b.randInt(0, size)
		str = append(str, abc[num])
		i++
	}

	return str
}
