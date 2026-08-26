package ql

import "fmt"

type Parser struct {
	tokens  []Token
	current int
}

func NewParser(tokens []Token) *Parser {
	return &Parser{
		tokens: tokens,
	}
}

func (p *Parser) peekWithIndex(index int) Token {
	return p.tokens[min(max(index, 0), len(p.tokens)-1)]
}

func (p *Parser) peek() Token {
	return p.peekWithIndex(p.current)
}

func (p *Parser) peekPrevious() Token {
	return p.peekWithIndex(max(p.current-1, 0))
}

func (p *Parser) advance() Token {
	token := p.peek()

	p.current += 1

	return token
}

func (p *Parser) consume(kind Kind) (Token, error) {
	token := p.advance()

	if token.Kind != kind {
		return Token{}, fmt.Errorf("The kind '%s' does not match '%s'", token.Kind.String(), kind.String())
	}

	return token, nil
}

func (p *Parser) parseBinary() (*ExpressionBinary, error) {
	previousToken := p.peekPrevious()

	binary := &ExpressionBinary{nil, previousToken, nil}
	curr := binary

	for {
		token := p.peek()
		if token.Kind == KindRightParens {
			break
		}

		expression, err := p.parseExpression()
		if err != nil {
			return nil, err
		}

		if curr.Left == nil {
			curr.Left = expression
		} else if curr.Right == nil {
			curr.Right = expression
		} else {
			curr.Right = &ExpressionBinary{curr.Right, previousToken, expression}
			curr = curr.Right.(*ExpressionBinary)
		}
	}

	_, err := p.consume(KindRightParens)
	if err != nil {
		return nil, err
	}

	// I want and to returns true by default on every branch for 'and'.
	// The opposite for 'or'.
	//
	// It works because atm there are only 2 binary operators.
	if binary.Left == nil {
		binary.Left = &ExpressionLiteral{Value: previousToken.Kind == KindAnd}
	}
	if binary.Right == nil {
		binary.Right = &ExpressionLiteral{Value: previousToken.Kind == KindAnd}
	}

	return binary, nil
}

func (p *Parser) parseFunctionCall() (*ExpressionFunctionCall, error) {
	previousToken := p.peekPrevious()

	arguments := []Expression{}
	for {
		token := p.peek()
		if token.Kind == KindRightParens {
			break
		}

		expression, err := p.parseExpression()
		if err != nil {
			return nil, err
		}

		arguments = append(arguments, expression)
	}

	_, err := p.consume(KindRightParens)
	if err != nil {
		return nil, err
	}

	return &ExpressionFunctionCall{
		Arguments:    arguments,
		FunctionName: previousToken.Lexeme,
	}, nil
}

func (p *Parser) parseString() (*ExpressionLiteral, error) {
	previous := p.peekPrevious()

	return &ExpressionLiteral{
		Value: previous.Literal,
	}, nil
}

func (p *Parser) parseNot() (*ExpressionNot, error) {
	expression, err := p.parseExpression()
	if err != nil {
		return nil, err
	}

	_, err = p.consume(KindRightParens)
	if err != nil {
		return nil, err
	}

	return &ExpressionNot{
		Expr: expression,
	}, nil
}

func (p *Parser) parseIf(isWhen bool) (*ExpressionIf, error) {
	condition, err := p.parseExpression()
	if err != nil {
		return nil, err
	}
	ifBranch, err := p.parseExpression()
	if err != nil {
		return nil, err
	}

	var elseBranch Expression
	if isWhen || p.peek().Kind == KindRightParens {
		elseBranch = &ExpressionLiteral{Value: false}
	} else {
		elseBranch, err = p.parseExpression()
		if err != nil {
			return nil, err
		}
	}

	_, err = p.consume(KindRightParens)
	if err != nil {
		return nil, err
	}

	return &ExpressionIf{
		Condition:  condition,
		IfBranch:   ifBranch,
		ElseBranch: elseBranch,
	}, nil
}

func (p *Parser) parseFilter() (Expression, error) {
	token := p.advance()

	switch token.Kind {
	case KindAnd, KindOr:
		return p.parseBinary()
	case KindIf:
		return p.parseIf(false)
	case KindWhen:
		return p.parseIf(true)
	case KindUnless:
		i, err := p.parseIf(true)
		if err != nil {
			return nil, err
		}

		i.Condition = &ExpressionNot{Expr: i.Condition}

		return i, nil
	case KindNot:
		return p.parseNot()
	case KindTitle, KindTag, KindExt,
		KindOnDate, KindFromDate, KindToDate:
		return p.parseFunctionCall()
	default:
		return nil, fmt.Errorf("The token '%s' is invalid.", token.String())
	}
}

func (p *Parser) parseExpression() (Expression, error) {
	token := p.advance()

	switch token.Kind {
	case KindLeftParens:
		return p.parseFilter()
	case KindString:
		return p.parseString()
	default:
		return nil, fmt.Errorf("Invalid token '%s'", token.String())
	}
}

// This function will always return a single expression
func (p *Parser) Parse() (Expression, error) {
	expression, err := p.parseExpression()
	if err != nil {
		return nil, err
	}

	_, err = p.consume(KindEOF)
	if err != nil {
		return nil, err
	}

	return expression, nil
}

func NewExpressionFromSource(source string) (Expression, error) {
	scanner := NewScanner(source)
	tokens, err := scanner.Scan()
	if err != nil {
		return nil, err
	}

	parser := NewParser(tokens)
	expression, err := parser.Parse()
	if err != nil {
		return nil, err
	}

	return expression, nil
}
