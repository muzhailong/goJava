package lang

var MIN_VALUE byte = -128
var MAX_VALUE byte = 127

type Byte struct {
	value byte
	Number
}

func NewByte(value byte) *Byte {
	tmp := Byte{value: value}
	return &tmp
}
