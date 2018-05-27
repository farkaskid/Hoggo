package types

type Raw struct {
	bytes []byte
	size  int
}

func (raw Raw) Value() []byte {
	return raw.bytes
}

func New(bytes []byte) Raw {
	return Raw{bytes, len(bytes)}
}
