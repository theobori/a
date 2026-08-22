package ql

import (
	"slices"
	"testing"
)

func TestScanner(t *testing.T) {
	tests := []struct {
		source         string
		expectedTokens []Token
	}{
		{
			source: "(((tag \"tag1\")))",
			expectedTokens: []Token{
				{Kind: KindLeftParens, Lexeme: "("},
				{Kind: KindLeftParens, Lexeme: "("},
				{Kind: KindLeftParens, Lexeme: "("},
				{Kind: KindTag, Lexeme: "tag"},
				{Kind: KindString, Lexeme: "\"tag1\"", Literal: "tag1"},
				{Kind: KindRightParens, Lexeme: ")"},
				{Kind: KindRightParens, Lexeme: ")"},
				{Kind: KindRightParens, Lexeme: ")"},
				{Kind: KindEOF, Lexeme: ""},
			},
		},
		{
			source: "(or (ext \".txt\") (title \"test-.*\"))",
			expectedTokens: []Token{
				{Kind: KindLeftParens, Lexeme: "("},
				{Kind: KindOr, Lexeme: "or"},
				{Kind: KindLeftParens, Lexeme: "("},
				{Kind: KindExt, Lexeme: "ext"},
				{Kind: KindString, Lexeme: "\".txt\"", Literal: ".txt"},
				{Kind: KindRightParens, Lexeme: ")"},
				{Kind: KindLeftParens, Lexeme: "("},
				{Kind: KindTitle, Lexeme: "title"},
				{Kind: KindString, Lexeme: "\"test-.*\"", Literal: "test-.*"},
				{Kind: KindRightParens, Lexeme: ")"},
				{Kind: KindRightParens, Lexeme: ")"},
				{Kind: KindEOF, Lexeme: ""},
			},
		},
		// TODO: add tests
	}

	for _, test := range tests {
		scanner := NewScanner(test.source)

		tokens, err := scanner.Scan()
		if err != nil {
			t.Fatal(err)
		}

		if !slices.Equal(tokens, test.expectedTokens) {
			t.Fatalf("The produced tokens %v shoud be equal to %v.", tokens, test.expectedTokens)
		}
	}
}
