package ql

const (
	OperationPop         byte = 0
	OperationTitle       byte = 1
	OperationTag         byte = 2
	OperationExt         byte = 3
	OperationFromDate    byte = 4
	OperationToDate      byte = 5
	OperationOnDate      byte = 6
	OperationConstant    byte = 7
	OperationJump        byte = 8
	OperationJumpIfFalse byte = 9
	OperationJumpIfTrue  byte = 10
	OperationNot         byte = 11
	OperationEnd         byte = 12
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
	case OperationFromDate:
		return "OperationFromDate"
	case OperationToDate:
		return "OperationToDate"
	case OperationOnDate:
		return "OperationOnDate"
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
