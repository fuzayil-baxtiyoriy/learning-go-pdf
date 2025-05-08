package main

type Sentinel string

func (s Sentinel) Error() string {
	return string(s)
}

const (
	ErrFoo = Sentinel("foo error")
	ErrBar = Sentinel("bar error")
)

type Status int 

const (
	InValidLogin Status = iota + 1
	NotFound
)

type StatusErr struct {
	Status Status
	Message string
}

func (se StatusErr) Error() string {
	return se.Message
}
