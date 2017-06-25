package flip_regexp

import (
	"testing"
)

func testTokenParserInMainContext(t *testing.T, context int, pattern string, tokenType, tokenLength int, strExpectedToken string) {
	b := NewBuilder([]byte(pattern))
	b.ContextParser = context
	token, err := b.getNextToken()

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
		[]interface{}{contextMain, "[", typeGroup, 1, "["},
		[]interface{}{contextMain, "\\[", typeLetter, 2, "["},
		[]interface{}{contextMain, "(", typeGroup, 1, "("},
		[]interface{}{contextMain, "\\d", typeAlias, 2, "d"},
		[]interface{}{contextMain, "\\w", typeAlias, 2, "w"},
		[]interface{}{contextMain, "\\s", typeAlias, 2, "s"},
		[]interface{}{contextMain, "\\D", typeAlias, 2, "D"},
		[]interface{}{contextMain, "\\W", typeAlias, 2, "W"},
		[]interface{}{contextMain, "\\S", typeAlias, 2, "S"},
		[]interface{}{contextMain, "?", typeQuantifier, 1, "?"},
		[]interface{}{contextMain, "+", typeQuantifier, 1, "+"},
		[]interface{}{contextMain, "*", typeQuantifier, 1, "*"},
		[]interface{}{contextMain, "a", typeLetter, 1, "a"},
		[]interface{}{contextMain, "", typeInvalid, 0, ""},
	}

	for _, row := range dataProvider {
		data := row.([]interface{})
		testTokenParserInMainContext(t, data[0].(int), data[1].(string), data[2].(int), data[3].(int), data[4].(string))
	}
}
