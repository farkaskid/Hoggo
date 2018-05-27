package types

type Text struct {
	Value string
}

func (text Text) Size() Number {
	return NewNumber(len(text.Value))
}

// Views
func (text Text) StringView() string {
	return text.Value
}

func (text Text) WriterView() []byte {
	return []byte(text.Value)
}
