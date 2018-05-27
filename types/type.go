package types

type Type interface {
	StringView() string
	WriterView() []byte
	Size() Number
}
