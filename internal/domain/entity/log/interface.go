package log

type IContent interface {
	Type() string
	Body() string
}
