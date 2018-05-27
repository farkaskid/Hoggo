package types

type Text struct {
	Raw
}

// Operations
func (value1 Text) Concat(value2 Text) Text {
	return Text{New([]byte(value1.StringView() + value2.StringView()))}
}

// Views
func (value Text) StringView() string {
	return string(value.Value())
}
