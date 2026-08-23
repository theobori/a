package ql

import (
	"fmt"
)

type Compiler struct {
	chunk Chunk
}

func NewCompiler() *Compiler {
	return &Compiler{
		chunk: Chunk{
			Constants: []any{},
			Code:      []byte{},
		},
	}
}

func (c *Compiler) emitByte(b byte) {
	c.chunk.Code = append(c.chunk.Code, b)
}

func (c *Compiler) addConstant(value any) (byte, error) {
	n := len(c.chunk.Constants)
	if n >= 256 {
		return 0, fmt.Errorf("You cannot have more than 256 constants.")
	}

	c.chunk.Constants = append(c.chunk.Constants, value)

	return byte(n), nil
}

func (c *Compiler) compileBinaryGeneric(binary *ExpressionBinary, operationJump byte) error {
	err := c.compile(binary.Left)
	if err != nil {
		return err
	}

	c.emitByte(operationJump)
	c.emitByte(0)
	c.emitByte(OperationPop)
	jumpIndex := len(c.chunk.Code) - 2

	err = c.compile(binary.Right)
	if err != nil {
		return err
	}

	if len(c.chunk.Code) >= 256 {
		return fmt.Errorf("Cannot jump above 255.")
	}
	c.chunk.Code[jumpIndex] = byte(len(c.chunk.Code))

	return nil
}

func (c *Compiler) compileAnd(binary *ExpressionBinary) error {
	return c.compileBinaryGeneric(binary, OperationJumpIfFalse)
}

func (c *Compiler) compileOr(binary *ExpressionBinary) error {
	return c.compileBinaryGeneric(binary, OperationJumpIfTrue)
}

func (c *Compiler) compileBinary(binary *ExpressionBinary) error {
	kind := binary.Operator.Kind

	switch kind {
	case KindAnd:
		return c.compileAnd(binary)
	case KindOr:
		return c.compileOr(binary)
	default:
		return fmt.Errorf("Invalid binary op %s", binary.String())
	}
}

func (c *Compiler) compileTitle(functionCall *ExpressionFunctionCall) error {
	if len(functionCall.Arguments) != 1 {
		return fmt.Errorf("The function 'title' should have only one argument.")
	}

	c.compile(functionCall.Arguments[0])
	c.emitByte(OperationTitle)

	return nil
}

func (c *Compiler) compileTag(functionCall *ExpressionFunctionCall) error {
	if len(functionCall.Arguments) != 1 {
		return fmt.Errorf("The function 'tag' should have only one argument.")
	}

	c.compile(functionCall.Arguments[0])
	c.emitByte(OperationTag)

	return nil
}

func (c *Compiler) compileExt(functionCall *ExpressionFunctionCall) error {
	if len(functionCall.Arguments) != 1 {
		return fmt.Errorf("The function 'ext' should have only one argument.")
	}

	c.compile(functionCall.Arguments[0])
	c.emitByte(OperationExt)

	return nil
}

func (c *Compiler) compileFunctionCall(functionCall *ExpressionFunctionCall) error {
	// Since its only native functions,
	// I prefer to check the functions arity earlier at compile time.
	switch functionCall.FunctionName {
	case "title":
		return c.compileTitle(functionCall)
	case "tag":
		return c.compileTag(functionCall)
	case "ext":
		return c.compileExt(functionCall)
	default:
		return fmt.Errorf("The function '%s' is unknown.", functionCall.FunctionName)
	}
}

func (c *Compiler) compileLiteral(literal *ExpressionLiteral) error {
	index, err := c.addConstant(literal.Value)
	if err != nil {
		return err
	}

	c.emitByte(OperationConstant)
	c.emitByte(index)

	return nil
}

func (c *Compiler) compile(expression Expression) error {
	switch e := expression.(type) {
	case *ExpressionBinary:
		return c.compileBinary(e)
	case *ExpressionFunctionCall:
		return c.compileFunctionCall(e)
	case *ExpressionLiteral:
		return c.compileLiteral(e)
	default:
		return fmt.Errorf("Invalid expressions type '%v'", e)
	}
}

func (c *Compiler) Compile(expression Expression) (*Chunk, error) {
	err := c.compile(expression)
	if err != nil {
		return nil, err
	}

	c.emitByte(OperationEnd)

	return &c.chunk, nil
}

func NewChunkFromSource(source string) (*Chunk, error) {
	expression, err := NewExpressionFromSource(source)
	if err != nil {
		return nil, err
	}

	compiler := NewCompiler()
	chunk, err := compiler.Compile(expression)
	if err != nil {
		return nil, err
	}

	return chunk, nil
}
