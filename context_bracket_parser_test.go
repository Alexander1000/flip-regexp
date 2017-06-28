package flip_regexp

import (
	"testing"
)

func TestGetToken_ContextBracketParser_SimpleIntervalWithSizeMinMax_Success(t *testing.T) {
	b := NewBuilder([]byte("[a-z]{5,6}"))
	token, err := b.getNextToken()

	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if token.Type != typeGroupOpen {
		t.Fatalf("Expected type %d, got: %d", typeGroupOpen, token.Type)
	}

	b.ContextParser = contextBracket
	b.Position += token.Length

	token, err = b.getNextToken()

	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if token.Type != typeInterval {
		t.Fatalf("Expected type %d, got: %d", typeInterval, token.Type)
	}

	if string(token.Stream) != "a-z" {
		t.Fatalf("Unexpected char stream, got: %v", string(token.Stream))
	}

	b.Position += token.Length
	token, err = b.getNextToken()

	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if token.Type != typeGroupClose {
		t.Fatalf("Expected type %d, got: %d", typeGroupClose, token.Type)
	}

	b.Position += token.Length
	token, err = b.getNextToken()

	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if token.Type != typeQuantifier {
		t.Fatalf("Expected type %d, got: %d", typeQuantifier, token.Type)
	}

	if string(token.Stream) != "{5,6}" {
		t.Fatalf("Unexpected char stream, got: %v", string(token.Stream))
	}

	if token.Min != 5 {
		t.Fatalf("Unexpected min size, got %v", token.Min)
	}

	if token.Max != 6 {
		t.Fatalf("Unexpected max size, got %v", token.Max)
	}
}

func TestGetToken_ContextBracketParser_SimpleIntervalWithSizeEmptyMinMax_Success(t *testing.T) {
	b := NewBuilder([]byte("[a-z]{,6}"))
	token, err := b.getNextToken()

	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if token.Type != typeGroupOpen {
		t.Fatalf("Expected type %d, got: %d", typeGroupOpen, token.Type)
	}

	b.ContextParser = contextBracket
	b.Position += token.Length

	token, err = b.getNextToken()

	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if token.Type != typeInterval {
		t.Fatalf("Expected type %d, got: %d", typeInterval, token.Type)
	}

	if string(token.Stream) != "a-z" {
		t.Fatalf("Unexpected char stream, got: %v", string(token.Stream))
	}

	b.Position += token.Length
	token, err = b.getNextToken()

	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if token.Type != typeGroupClose {
		t.Fatalf("Expected type %d, got: %d", typeGroupClose, token.Type)
	}

	b.Position += token.Length
	token, err = b.getNextToken()

	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if token.Type != typeQuantifier {
		t.Fatalf("Expected type %d, got: %d", typeQuantifier, token.Type)
	}

	if string(token.Stream) != "{,6}" {
		t.Fatalf("Unexpected char stream, got: %v", string(token.Stream))
	}

	if token.Min != 0 {
		t.Fatalf("Unexpected min size, got %v", token.Min)
	}

	if token.Max != 6 {
		t.Fatalf("Unexpected max size, got %v", token.Max)
	}
}

func TestGetToken_ContextBracketParser_SimpleIntervalWithSizeOnlyMax_Success(t *testing.T) {
	b := NewBuilder([]byte("[a-z]{9}"))
	token, err := b.getNextToken()

	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if token.Type != typeGroupOpen {
		t.Fatalf("Expected type %d, got: %d", typeGroupOpen, token.Type)
	}

	b.ContextParser = contextBracket
	b.Position += token.Length

	token, err = b.getNextToken()

	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if token.Type != typeInterval {
		t.Fatalf("Expected type %d, got: %d", typeInterval, token.Type)
	}

	if string(token.Stream) != "a-z" {
		t.Fatalf("Unexpected char stream, got: %v", string(token.Stream))
	}

	b.Position += token.Length
	token, err = b.getNextToken()

	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if token.Type != typeGroupClose {
		t.Fatalf("Expected type %d, got: %d", typeGroupClose, token.Type)
	}

	b.Position += token.Length
	token, err = b.getNextToken()

	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if token.Type != typeQuantifier {
		t.Fatalf("Expected type %d, got: %d", typeQuantifier, token.Type)
	}

	if string(token.Stream) != "{9}" {
		t.Fatalf("Unexpected char stream, got: %v", string(token.Stream))
	}

	if token.Min != 0 {
		t.Fatalf("Unexpected min size, got %v", token.Min)
	}

	if token.Max != 9 {
		t.Fatalf("Unexpected max size, got %v", token.Max)
	}
}
