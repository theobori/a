package ql

import (
	"fmt"
	"strings"
)

type Expression interface {
	String() string
}

type ExpressionBinary struct {
	Left     Expression
	Operator Token
	Right    Expression
}

func (e *ExpressionBinary) String() string {
	return fmt.Sprintf("<binary>(%s)", e.Operator.Kind.String())
}

type ExpressionFunctionCall struct {
	FunctionName string
	Arguments    []Expression
}

func (e *ExpressionFunctionCall) String() string {
	return fmt.Sprintf("<function-call>(%s)", e.FunctionName)
}

type ExpressionLiteral struct {
	Value any
}

func (e *ExpressionLiteral) String() string {
	return fmt.Sprintf("<literal>(%v)", e.Value)
}

func dumpExpression(expression Expression, indent int) string {
	ans := []string{}
	indentString := strings.Repeat("  ", indent)

	switch e := expression.(type) {
	case *ExpressionBinary:
		ans = append(ans, e.String())
		ans = append(ans, indentString+dumpExpression(e.Left, indent+2))
		ans = append(ans, indentString+dumpExpression(e.Right, indent+2))
	case *ExpressionFunctionCall:
		ans = append(ans, e.String())
		for _, argument := range e.Arguments {
			ans = append(ans, indentString+dumpExpression(argument, indent+2))
		}
	case *ExpressionLiteral:
		ans = append(ans, e.String())
	case nil:
		ans = append(ans, "<nil>")
	default:
		ans = append(ans, "Invalid")
	}

	return strings.Join(ans, "\n")
}

func DumpExpression(expression Expression) string {
	return dumpExpression(expression, 2)
}
