package reader

import (
	"testing"
)

func TestLexer(t *testing.T) {

	text := "state q1 match a write b next q2\n"
	l := NewLexer(text)

	expected := []struct {
		Type    TokenType
		Literal string
	}{
		{Type: KEYWORD, Literal: "state"},
		{Type: SYMBOL, Literal: "q1"},
		{Type: KEYWORD, Literal: "match"},
		{Type: CHAR, Literal: "a"},
		{Type: KEYWORD, Literal: "write"},
		{Type: CHAR, Literal: "b"},
		{Type: KEYWORD, Literal: "next"},
		{Type: SYMBOL, Literal: "q2"},
		{Type: EOL, Literal: ""},
		{Type: EOT, Literal: ""},
	}

	for i, e := range expected {
		tk := l.nextToken()
		if tk.Type != e.Type {
			t.Fatalf("No.%v expected type=%q, got=%q", i, e.Type, tk.Type)
		}
		if tk.Literal != e.Literal {
			t.Fatalf("No.%v expected literal=%q, got=%q", i, e.Literal, tk.Literal)
		}
	}
}
