package core

type Event interface {
	GetType() string
	GetDescription() string
	GetPath() string
}
