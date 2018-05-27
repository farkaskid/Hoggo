package types

import (
	"bytes"
	"encoding/binary"
	"strconv"
)

type Number struct {
	value int64
}

func NewNumber(value int) Number {
	return Number{int64(value)}
}

func (number Number) StringView() string {
	return strconv.Itoa(int(number.value))
}

func (number Number) WriterView() []byte {
	buffer := new(bytes.Buffer)

	binary.Write(buffer, binary.BigEndian, number.value)

	return buffer.Bytes()
}
