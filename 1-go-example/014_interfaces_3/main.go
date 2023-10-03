/*
  - Herhangi bir yöntemi olmayan interface'e Emtpy Interface denir.
    Empty Interface her tür veriyi depolayabilir, bu nedenle genellikle dinamik parametrelere
    (tüm veri türleri) ihtiyaç duyan işlevlerin işlenmesi durumunda kullanılır.
*/
package main

import "fmt"

func func1() {
	var i interface{} // Empty interface
	i = 42
	describe1(i)
	i = "hello"
	describe1(i)
}

// we can pass all data type in to function describe
func describe1(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}

func func2() {
	var i any // Empty interface
	i = 42
	describe2(i)
	i = "hello"
	describe2(i)
}

// we can pass all data type in to function describe
func describe2(i any) {
	fmt.Printf("(%v, %T)\n", i, i)
}

func main() {
	func1()
	func2()
}
