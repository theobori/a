package ql

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestParser(t *testing.T) {
	tests := []struct {
		source             string
		expectedExpression Expression
	}{
		{
			source: "(tag \"tag1\")",
			expectedExpression: ExpressionFunctionCall{
				FunctionName: "tag",
				Arguments: []Expression{
					ExpressionLiteral{
						Value: "tag1",
					},
				},
			},
		},
		{
			source: `(and
						(or
						    (and (title "hello")))
					 (tag "tag1"))`,
			expectedExpression: ExpressionBinary{
				Operator: Token{
					Kind:   KindAnd,
					Lexeme: "and",
				},
				Left: ExpressionBinary{
					Operator: Token{
						Kind:   KindOr,
						Lexeme: "or",
					},
					Left: ExpressionBinary{
						Operator: Token{
							Kind:   KindAnd,
							Lexeme: "and",
						},
						Left: ExpressionFunctionCall{
							FunctionName: "title",
							Arguments: []Expression{
								ExpressionLiteral{
									Value: "hello",
								},
							},
						},
						Right: nil,
					},
					Right: nil,
				},
				Right: ExpressionFunctionCall{
					FunctionName: "tag",
					Arguments: []Expression{
						ExpressionLiteral{
							Value: "tag1",
						},
					},
				},
			},
		},
		// TODO: add tests
	}

	for _, test := range tests {
		expression, err := NewExpressionFromSource(test.source)
		if err != nil {
			t.Fatal(err)
		}

		if !cmp.Equal(expression, test.expectedExpression) {
			fmt.Println("Below the produced expression.")
			fmt.Println(DumpExpression(expression))

			fmt.Println("Below the expected expression.")
			fmt.Println(DumpExpression(test.expectedExpression))

			t.Fatalf("The produced expression %v should be equal to %v.", expression, test.expectedExpression)
		}
	}
}
