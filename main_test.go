package flip_regexp

import (
	"regexp"
	"testing"
)

func testRegexpGenerator(t *testing.T, pattern string) {
	i := 0

	for i < 10 {
		reg := NewBuilder([]byte(pattern))
		result, _ := reg.Render()

		re := regexp.MustCompile(pattern)

		if !re.MatchString(string(result)) {
			t.Fatalf("Excpeted success, false given for pattern: %s; string: %s", pattern, result)
		} else {
			t.Logf("Generated string: %s", result)
		}

		i++
	}
}

func TestRegexp_SetPatterns_Success(t *testing.T) {
	patternSet := []string{
		"[a-z]{5,6}",
		"[a-z]{5,6} test [a-zA-Z]{9}",
		`\+7 \([489][0-9]{2}\) [0-9]{3} [0-9]{2} [0-9]{2}`,
		`\+7 \(\[489\][0-9]{2}\) [0-9]{3} \[0-9\]\{2\} [0-9]{2}`,
		`(status|very|important)`,
		`hello (status|very|important) and (man|woman)? test`,
		"[A-F]{3}hellow [^0-9a-zA-Z]{7} moto",
		"[a-z]+",
		"hello [alexander]? it is your [qwerty90]* name?",
	}

	for _, pattern := range patternSet {
		testRegexpGenerator(t, pattern)
	}
}
