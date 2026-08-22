package ql

import "fmt"

type Token struct {
	Kind    Kind
	Lexeme  string
	Literal any
}

func (t *Token) String() string {
	return fmt.Sprintf(
		"(%s, %s, '%v')",
		t.Kind.String(),
		t.Lexeme,
		t.Literal,
	)
}
