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
	case KindString:
		return "String"
	case KindLeftParens:
		return "LeftParens"
	case KindRightParens:
		return "RightParens"
	case KindEOF:
		return "EOF"
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
	KindString
	KindLeftParens
	KindRightParens
	KindEOF
)
