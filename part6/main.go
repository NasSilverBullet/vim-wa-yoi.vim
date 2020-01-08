package main

import "fmt"

/*
GoImpl

gr LspRename

GoRun
*/

func main() {
	fmt.Println("part6")
}

type User interface {
	GetName() string
	GetAge() int
	IsVIP() bool
}

type UserImpl struct {
}

func (u *UserImpl) GetName() string {
	panic("not implemented") // TODO: Implement
}

func (u *UserImpl) GetAge() int {
	panic("not implemented") // TODO: Implement
}

func (u *UserImpl) IsVIP() bool {
	panic("not implemented") // TODO: Implement
}
