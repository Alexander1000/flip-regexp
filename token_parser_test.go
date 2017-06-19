package flip_regexp

import (
	"testing"
)

func TestMainContextTokenParser_OpenBracketToken_Success(t *testing.T) {
	b := NewBuilder([]byte("[qwerty]"))
	token, err := b.getNextTokenInMainContext()

	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if token.Type != typeBracketOpen {
		t.Fatalf("Expected bracket open, given: %d", token.Type)
	}

	if token.Length != 1 {
		t.Fatalf("Expected length 1, got: %d", token.Length)
	}

	if token.Stream[0] != tokenBracketOpen {
		t.Fatalf("Expected '%s', got '%s'", string(tokenBracketOpen), string(token.Stream))
	}
}

func TestMainContextTokenParser_EscapedValueToken_Success(t *testing.T) {
	b := NewBuilder([]byte("\\[qwerty"))
	token, err := b.getNextTokenInMainContext()

	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if token.Type != typeLetter {
		t.Fatalf("Expected type letter, got: %d", token.Type)
	}

	if token.Length != 2 {
		t.Fatalf("Expected length 2, got: %d", token.Length)
	}

	if string(token.Stream) != "[" {
		t.Fatalf("Expected '%s', got '%s'", "[", string(token.Stream))
	}
}

func TestMainContextTokenParser_Data_Success(t *testing.T) {
	// todo описать с помощью dataProvider test-case-ы из 2х тестов выше и дополнить их
}
