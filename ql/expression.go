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

func (e ExpressionBinary) String() string {
	return "<binary>"
}

type ExpressionFunctionCall struct {
	FunctionName string
	Arguments    []Expression
}

func (e ExpressionFunctionCall) String() string {
	return "<function-call>"
}

type ExpressionLiteral struct {
	Value any
}

func (e ExpressionLiteral) String() string {
	return "<literal>"
}

func dumpExpression(e Expression, indent int) string {
	ans := []string{}
	indentString := strings.Repeat("  ", indent)

	switch e := e.(type) {
	case ExpressionBinary:
		ans = append(ans, "Binary")
		ans = append(ans, indentString+e.Operator.String())
		ans = append(ans, indentString+dumpExpression(e.Left, indent+2))
		ans = append(ans, indentString+dumpExpression(e.Right, indent+2))
	case ExpressionFunctionCall:
		ans = append(ans, "FunctionCall")
		ans = append(ans, indentString+e.FunctionName)

		for _, argument := range e.Arguments {
			ans = append(ans, indentString+dumpExpression(argument, indent+2))
		}
	case ExpressionLiteral:
		ans = append(ans, "Literal")
		ans = append(ans, indentString+fmt.Sprintf("%v", e.Value))
	}

	return strings.Join(ans, "\n")
}

func DumpExpression(expression Expression) string {
	return dumpExpression(expression, 2)
}
