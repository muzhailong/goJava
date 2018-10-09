package lang

import (
	"strings"
)

var TRUE = NewBoolean(true)
var FALSE = NewBoolean(false)

type Boolean struct {
	value bool
	Number
}

func NewBoolean(value bool) *Boolean {
	b := Boolean{}
	b.value = value
	return &b
}

func (this *Boolean) BooleanValue() bool {
	return this.value
}

func (this *Boolean) ToString() string {
	if this.value {
		return "true"
	} else {
		return "false"
	}
}

func (this *Boolean) CompareTo(b interface{}) int {
	tmp, _ := b.(Boolean)
	return Compare(this.value, tmp.value)
}

func Compare(x, y bool) int {
	if x == y {
		return 0
	} else if x {
		return 1
	} else {
		return -1
	}
}

func (this *Boolean) Equal(obj interface{}) bool {
	if tmp, ok := obj.(Boolean); ok {
		return tmp.value == this.value
	}
	return false
}

func ParseBoolean(s string) bool {
	return strings.ToLower(s) == "true"
}

func ValueOf(b bool) *Boolean {
	if b {
		return TRUE
	} else {
		return FALSE
	}
}
