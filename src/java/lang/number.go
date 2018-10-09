package lang

type Number struct {
}

func (this *Number) IntValue() int {
	return 0
}

func (this *Number) LongValue() int64 {
	return 0
}

func (this *Number) FloatValue() float32 {
	return 0
}

func (this *Number) DoubleValue() float64 {
	return 0
}

func (this *Number) ByteValue() byte {
	return byte(this.IntValue())
}

func (this *Number) ShortValue() int16 {
	return int16(this.IntValue())
}
