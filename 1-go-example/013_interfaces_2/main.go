package main

import "fmt"

type I interface {
	M()
	N()
}
type T struct {
	S string
}

func (t T) M() {
	fmt.Println(t.S)
}

func (t T) N() {
}

func main() {
	var i I = T{"hello"}
	i.M()
	// The result will be an error because of struct T implement interface I,but there are not enough declared methods (missing method N ())
}
