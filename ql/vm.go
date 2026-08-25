package ql

import (
	"fmt"
	"time"

	"github.com/theobori/a/internal"
)

const timeLayout string = time.DateOnly

type VM struct {
	chunk *Chunk
	ip    int
	stack internal.Stack[any]
}

func isFalsey(vAny any) bool {
	switch v := vAny.(type) {
	case bool:
		return v == false
	case int8, int16, int32, int64,
		uint8, uint16, uint32, uint64,
		float32, float64:
		return v == 0
	case string:
		return v == ""
	case nil:
		return false
	default:
		return true
	}
}

func isTrue(v any) bool {
	return !isFalsey(v)
}

func NewVM(chunk *Chunk) *VM {
	return &VM{
		chunk: chunk,
		ip:    0,
	}
}

func (vm *VM) readByte() byte {
	ans := vm.chunk.Code[vm.ip]
	// yes i cannot put v.ip++ between brackets...
	vm.ip++

	return ans
}

func (vm *VM) run(ctx *Context) (bool, error) {
	for {
		b := vm.readByte()

		switch b {
		case OperationEnd:
			v, err := vm.stack.Pop()
			if err != nil {
				return false, err
			}

			return isTrue(v), nil
		case OperationConstant:
			index := vm.readByte()
			constant := vm.chunk.Constants[index]
			vm.stack.Push(constant)
		case OperationNot:
			v, err := vm.stack.Pop()
			if err != nil {
				return false, err
			}

			vm.stack.Push(!isTrue(v))
		case OperationJumpIfFalse:
			jumpIndex := vm.readByte()

			if isFalsey(vm.stack.Top()) {
				vm.ip = int(jumpIndex)
			}
		case OperationJumpIfTrue:
			jumpIndex := vm.readByte()

			if isTrue(vm.stack.Top()) {
				vm.ip = int(jumpIndex)
			}
		case OperationPop:
			_, err := vm.stack.Pop()
			if err != nil {
				return false, err
			}
		case OperationTitle:
			v, err := vm.stack.Pop()
			if err != nil {
				return false, err
			}

			s, ok := v.(string)
			if !ok {
				return false, fmt.Errorf("The function 'title' argument '%v' must be a string.", v)
			}

			vm.stack.Push(ctx.FilenameMetadata.Title == s)
		case OperationTag:
			v, err := vm.stack.Pop()
			if err != nil {
				return false, err
			}

			s, ok := v.(string)
			if !ok {
				return false, fmt.Errorf("The function 'tag' argument '%v' must be a string.", v)
			}

			vm.stack.Push(ctx.FilenameMetadata.Tags.Exists(s))
		case OperationExt:
			v, err := vm.stack.Pop()
			if err != nil {
				return false, err
			}

			s, ok := v.(string)
			if !ok {
				return false, fmt.Errorf("The function 'ext' argument '%v' must be a string.", v)
			}

			vm.stack.Push(ctx.FilenameMetadata.Ext == s)
		case OperationFromDate:
			v, err := vm.stack.Pop()
			if err != nil {
				return false, err
			}

			dateString, ok := v.(string)
			if !ok {
				return false, fmt.Errorf("The function 'from-date' argument '%v' must be a string.", v)
			}
			t, err := time.Parse(timeLayout, dateString)
			if err != nil {
				return false, err
			}

			vm.stack.Push(ctx.Date.After(t))
		case OperationToDate:
			v, err := vm.stack.Pop()
			if err != nil {
				return false, err
			}

			dateString, ok := v.(string)
			if !ok {
				return false, fmt.Errorf("The function 'to-date' argument '%v' must be a string.", v)
			}
			t, err := time.Parse(timeLayout, dateString)
			if err != nil {
				return false, err
			}

			vm.stack.Push(internal.DateDayEqual(ctx.Date, t) || ctx.Date.Before(t))
		case OperationOnDate:
			v, err := vm.stack.Pop()
			if err != nil {
				return false, err
			}

			dateString, ok := v.(string)
			if !ok {
				return false, fmt.Errorf("The function 'on-date' argument '%v' must be a string.", v)
			}
			t, err := time.Parse(timeLayout, dateString)
			if err != nil {
				return false, err
			}

			vm.stack.Push(internal.DateDayEqual(ctx.Date, t))
		default:
			return false, fmt.Errorf("Invalid operation '%b'", b)
		}
	}
}

func (vm *VM) Interpret(ctx *Context) (bool, error) {
	return vm.run(ctx)
}
