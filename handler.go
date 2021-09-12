package examples

type echo struct{}

func (e *echo) Echo(msg string) string {
	return msg
}

func NewEcho() Echo {
	return &echo{}
}
