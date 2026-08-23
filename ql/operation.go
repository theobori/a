package ql

const (
	OperationPop         byte = 0
	OperationTitle       byte = 1
	OperationTag         byte = 2
	OperationExt         byte = 3
	OperationConstant    byte = 4
	OperationJump        byte = 5
	OperationJumpIfFalse byte = 6
	OperationJumpIfTrue  byte = 7
	OperationNot         byte = 8
	OperationEnd         byte = 9
)

func NewStringFromOperation(b byte) string {
	switch b {
	case OperationPop:
		return "OperationPop"
	case OperationTitle:
		return "OperationTitle"
	case OperationTag:
		return "OperationTag"
	case OperationExt:
		return "OperationExt"
	case OperationConstant:
		return "OperationConstant"
	case OperationJump:
		return "OperationJump"
	case OperationJumpIfFalse:
		return "OperationJumpIfFalse"
	case OperationJumpIfTrue:
		return "OperationJumpIfFalse"
	case OperationNot:
		return "OperationNot"
	case OperationEnd:
		return "OperationEnd"
	default:
		return "Unknown"
	}
}
