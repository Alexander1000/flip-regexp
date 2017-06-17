package flip_regexp

import (
	"testing"
)

func TestMainContextTokenParser_Success(t *testing.T) {
	b := NewBuilder([]byte("[fd]"))
	token, err := b.getNextTokenInMainContext()

	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if token.Type != typeBracketOpen {
		t.Fatalf("Expected bracket open, given: %v", token.Type)
	}

	if token.Length != 1 {
		t.Fatalf("Expected length 1, got: %v", token.Length)
	}

	if token.Stream[0] != tokenBracketOpen {
		t.Fatalf("Expected '%s', got '%s'", string(tokenBracketOpen), string(token.Stream))
	}
}

