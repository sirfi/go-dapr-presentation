package main

import "fmt"

// Interface I has method M()
type I interface {
	M()
}
type T struct {
	S string
}

// Define method M() for struct T
func (t T) M() {
	fmt.Println(t.S)
}
func main() {
	// init variable i with type interface I
	var i I = T{"hello"}
	// call method M()
	i.M()
}
