package socket

type Socket interface {
	GetName() string
}

type RowSocket struct {
	name     string
	protocol string
}

func (r *RowSocket) GetName() string {
	return r.name
}

func NewRawSocket() *RowSocket {
	return &RowSocket{}
}
