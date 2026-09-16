package ql

import (
	"fmt"

	"github.com/theobori/a/internal"
)

const EOFByte = 0x00

var ReservedKeywords = map[string]Kind{
	"and":       KindAnd,
	"or":        KindOr,
	"title":     KindTitle,
	"tag":       KindTag,
	"ext":       KindExt,
	"from-date": KindFromDate,
	"to-date":   KindToDate,
	"on-date":   KindOnDate,
	"not":       KindNot,
	"if":        KindIf,
	"when":      KindWhen,
	"unless":    KindUnless,
}

func isIdentifierChar(b byte) bool {
	return internal.IsAlphaNumeric(b) || b == '-' || b == '_'
}

type Scanner struct {
	source string
	tokens []Token

	current int
	start   int
}

func NewScanner(source string) *Scanner {
	return &Scanner{
		source:  source,
		tokens:  []Token{},
		current: 0,
		start:   0,
	}
}

func (s *Scanner) isAtEnd() bool {
	return s.current >= len(s.source)
}

func (s *Scanner) peek() byte {
	if s.isAtEnd() {
		return EOFByte
	}

	return s.source[s.current]
}

func (s *Scanner) advance() byte {
	b := s.peek()
	if b == EOFByte {
		return EOFByte
	}

	s.current += 1

	return b
}

func (s *Scanner) scanString(encloseChar byte) error {
loop:
	for {
		ch := s.peek()
		switch ch {
		case EOFByte:
			return fmt.Errorf("Unterminated string")
		case encloseChar:
			break loop
		}

		s.advance()
	}

	for s.peek() == encloseChar {
		s.advance()
	}

	literal := s.source[s.start+1 : s.current-1]
	s.pushToken(KindString, literal)

	return nil
}

func (s *Scanner) scanReservedKeyword() error {
	for isIdentifierChar(s.peek()) {
		s.advance()
	}

	keyword := s.source[s.start:s.current]
	kind, ok := ReservedKeywords[keyword]
	if !ok {
		return fmt.Errorf("The keyword '%s' is not allowed.", keyword)
	}

	s.pushToken(kind, nil)

	return nil
}

func (s *Scanner) scanComment() error {
	secondCh := s.advance()
	if secondCh != ';' {
		return fmt.Errorf("Invalid comment syntax")
	}

	for {
		ch := s.advance()
		if ch == '\n' || ch == EOFByte {
			break
		}
	}

	return nil
}

func (s *Scanner) scan() error {
	ch := s.advance()

	switch ch {
	case '(':
		s.pushToken(KindLeftParens, nil)
	case ')':
		s.pushToken(KindRightParens, nil)
	case '\t', '\r', '\n', ' ':
		break
	case '"', '\'':
		return s.scanString(ch)
	case ';':
		return s.scanComment()
	default:
		if isIdentifierChar(ch) {
			return s.scanReservedKeyword()
		}

		return fmt.Errorf("Invalid character '%c'", ch)
	}

	return nil
}

func (s *Scanner) pushToken(kind Kind, literal any) {
	token := Token{
		Kind:    kind,
		Lexeme:  s.source[s.start:s.current],
		Literal: literal,
	}

	s.tokens = append(s.tokens, token)
}

func (s *Scanner) Scan() ([]Token, error) {
	for !s.isAtEnd() {
		s.start = s.current

		err := s.scan()
		if err != nil {
			return []Token{}, err
		}
	}

	eofToken := Token{Kind: KindEOF, Lexeme: "", Literal: nil}
	s.tokens = append(s.tokens, eofToken)

	return s.tokens, nil
}
