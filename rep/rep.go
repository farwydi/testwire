package rep

func NewRep() *Rep {
	return &Rep{}
}

type Rep struct {
}

func (r Rep) X() {}
