package ql

type Kind byte

func (k *Kind) String() string {
	switch *k {
	case KindAnd:
		return "And"
	case KindOr:
		return "Or"
	case KindTitle:
		return "Title"
	case KindTag:
		return "Tag"
	case KindExt:
		return "Ext"
	case KindFromDate:
		return "FromDate"
	case KindToDate:
		return "ToDate"
	case KindOnDate:
		return "OnDate"
	case KindString:
		return "String"
	case KindLeftParens:
		return "LeftParens"
	case KindRightParens:
		return "RightParens"
	case KindEOF:
		return "EOF"
	case KindNot:
		return "Not"
	default:
		return "Unknown"
	}
}

const (
	KindAnd Kind = iota
	KindOr
	KindTitle
	KindTag
	KindExt
	KindFromDate
	KindToDate
	KindOnDate
	KindString
	KindLeftParens
	KindRightParens
	KindEOF
	KindNot
)
