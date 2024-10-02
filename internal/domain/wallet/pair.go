package wallet

type Pair string

const (
	EURPLN Pair = "EURPLN"
)

func (p Pair) AsString() string {
	return string(p)
}

func (p Pair) To() string {
	return string(p[3:6])
}

func (p Pair) From() string {
	return string(p[0:3])
}
