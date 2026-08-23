package ql

type Chunk struct {
	Constants []any // 255 max
	Code      []byte
}
