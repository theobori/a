package ql

const (
	// TODO: Remove unused
	OperationPush        byte = 0
	OperationPop         byte = 1
	OperationAnd         byte = 2
	OperationOr          byte = 3
	OperationTitle       byte = 4
	OperationTag         byte = 5
	OperationExt         byte = 6
	OperationConstant    byte = 7
	OperationJump        byte = 8
	OperationJumpIfFalse byte = 9
	OperationJumpIfTrue  byte = 10
	OperationNot         byte = 11
	OperationEnd         byte = 12
)

func NewStringFromOperation(b byte) string {
	switch b {
	case OperationPush:
		return "OperationPush"
	case OperationPop:
		return "OperationPop"
	case OperationAnd:
		return "OperationAnd"
	case OperationOr:
		return "OperationOr"
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
