package flip_regexp

import (
	"testing"
)

func testTokenParserInMainContext(t *testing.T, pattern string, tokenType, tokenLength int, strExpectedToken string) {
	b := NewBuilder([]byte(pattern))
	token, err := b.getNextTokenInMainContext()

	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if token.Type != tokenType {
		t.Fatalf("Expected type %d, got: %d", tokenType, token.Type)
	}

	if token.Length != tokenLength {
		t.Fatalf("Expected length %d, got: %d", tokenLength, token.Length)
	}

	if string(token.Stream) != strExpectedToken {
		t.Fatalf("Expected '%s', got '%s'", strExpectedToken, string(token.Stream))
	}
}

func TestMainContextTokenParser_Data_Success(t *testing.T) {
	dataProvider := []interface{}{
		[]interface{}{"[qwerty]", typeGroup, 1, "["},
		[]interface{}{"\\[qwerty", typeLetter, 2, "["},
		[]interface{}{"(qwerty)", typeGroup, 1, "("},
		[]interface{}{"\\d+", typeAlias, 2, "d"},
		[]interface{}{"\\w+", typeAlias, 2, "w"},
		[]interface{}{"\\s+", typeAlias, 2, "s"},
		[]interface{}{"\\D+", typeAlias, 2, "D"},
		[]interface{}{"\\W+", typeAlias, 2, "W"},
		[]interface{}{"\\S+", typeAlias, 2, "S"},
		[]interface{}{"?", typeQuantifier, 1, "?"},
		[]interface{}{"+", typeQuantifier, 1, "+"},
		[]interface{}{"*", typeQuantifier, 1, "*"},
	}

	for _, row := range dataProvider {
		data := row.([]interface{})
		testTokenParserInMainContext(t, data[0].(string), data[1].(int), data[2].(int), data[3].(string))
	}
}
