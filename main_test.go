package flip_regexp

import (
	"regexp"
	"testing"
)

func TestRegexp_Success(t *testing.T) {
	pattern := "[a-z]{5,6}"
	reg := NewBuilder([]byte(pattern))
	result, _ := reg.Render()

	re := regexp.MustCompile(pattern)

	if !re.MatchString(string(result)) {
		t.Fatalf("Excpeted success, false given for pattern: %s; string: %s", pattern, result)
	} else {
		t.Logf("Generated string: %s", result)
	}
}

func TestRegexp_Complex_Success(t *testing.T) {
	pattern := "[a-z]{5,6} test [a-zA-Z]{9}"
	reg := NewBuilder([]byte(pattern))
	result, _ := reg.Render()

	re := regexp.MustCompile(pattern)

	if !re.MatchString(string(result)) {
		t.Fatalf("Excpeted success, false given for pattern: %s; string: %s", pattern, result)
	} else {
		t.Logf("Generated string: %s", result)
	}
}

func TestRegexp_PhoneComplex_Success(t *testing.T) {
	pattern := `\+7 \([489][0-9]{2}\) [0-9]{3} [0-9]{2} [0-9]{2}`
	reg := NewBuilder([]byte(pattern))
	result, _ := reg.Render()

	re := regexp.MustCompile(pattern)

	if !re.MatchString(string(result)) {
		t.Fatalf("Excpeted success, false given for pattern: %s; string: %s", pattern, result)
	} else {
		t.Logf("Generated string: %s", result)
	}
}

func TestRegexp_EscapeComplexPattern_Success(t *testing.T) {
	pattern := `\+7 \(\[489\][0-9]{2}\) [0-9]{3} \[0-9\]\{2\} [0-9]{2}`
	reg := NewBuilder([]byte(pattern))
	result, _ := reg.Render()

	re := regexp.MustCompile(pattern)

	if !re.MatchString(string(result)) {
		t.Fatalf("Excpeted success, false given for pattern: %s; string: %s", pattern, result)
	} else {
		t.Logf("Generated string: %s", result)
	}
}

func TestRegexp_OrWithPipePattern_Success(t *testing.T) {
	pattern := `(status|very|important)`
	reg := NewBuilder([]byte(pattern))
	result, _ := reg.Render()

	re := regexp.MustCompile(pattern)

	if !re.MatchString(string(result)) {
		t.Fatalf("Excpeted success, false given for pattern: %s; string: %s", pattern, result)
	} else {
		t.Logf("Generated string: %s", result)
	}
}

func TestRegexp_OrWithPipeAndQuestionPattern_Success(t *testing.T) {
	pattern := `hello (status|very|important) and (man|woman)? test`
	reg := NewBuilder([]byte(pattern))
	result, _ := reg.Render()

	re := regexp.MustCompile(pattern)

	if !re.MatchString(string(result)) {
		t.Fatalf("Excpeted success, false given for pattern: %s; string: %s", pattern, result)
	} else {
		t.Logf("Generated string: %s", result)
	}
}
