package flip_regexp

import (
	"strconv"
)

// генерация строки для шаблона: []{,}
func (b *Builder) parseInBracket() {
	var abc []byte
	var prev byte
	interval := false
	firstChar := true
	negative := false

	for b.Position < len(b.Pattern) {
		letter := b.getCurrentSymbol()
		b.Position++

		if letter == tokenCircumflex && firstChar {
			negative = true
		} else if b.isLetter(letter) {
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

		firstChar = false
	}

	var str []byte
	min := 1
	max := 1
	var size []int
	var bracket, plus, asterisk, question bool

	// сканируем потенциально возможные управляющие конструкции
	for b.Position < len(b.Pattern) {
		letter := b.getCurrentSymbol()

		if !bracket && letter == tokenBraceOpen {
			bracket = true
		} else if bracket && letter == tokenBraceClose {
			bracket = false
			b.Position++
			break
		} else if bracket && b.isDigit(letter) {
			str = append(str, letter)
		} else if bracket && letter == tokenComma {
			tSize, _ := strconv.Atoi(string(str))
			size = append(size, tSize)
			str = []byte{}
		} else if !bracket && letter == tokenPlus {
			plus = true
			b.Position++
			break
		} else if !bracket && letter == tokenAsterisk {
			asterisk = true
			b.Position++
			break
		} else if !bracket && letter == tokenQuestion {
			question = true
			b.Position++
			break
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

	if plus {
		min = 1
		max = randomMax
	}

	if asterisk {
		min = 1
		max = randomMax
	}

	if !plus && !asterisk {
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
	}

	var length int
	length = b.randInt(min, max)

	if length <= 0 || (question && b.randInt(0, 2) == 1) {
		return
	}

	if negative {
		i := letterMinChar
		lengthAntiSlice := int(letterMaxChar-letterMinChar) - len(abc)

		if lengthAntiSlice <= 0 {
			return
		}

		antiAbc := make([]byte, lengthAntiSlice)

		for i <= letterMaxChar || len(antiAbc) < lengthAntiSlice {
			exists := false

			for _, letter := range abc {
				if letter == i {
					exists = true
					break
				}
			}

			if !exists {
				antiAbc = append(antiAbc, i)
			}

			i++
		}

		if len(antiAbc) > 0 {
			abc = antiAbc
		} else {
			return
		}
	}

	b.Result = append(b.Result, b.randomString(length, abc)...)
}
