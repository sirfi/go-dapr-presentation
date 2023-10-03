/*
  - Interface uygulama tipinin, interfacede tanımlanan yöntemlerin
    tam bir kümesine sahip olması gerekir.
*/
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
	// Sonuç, struct T I interface'i implementasyonu nedeniyle bir hata olacaktır, ancak yeterince bildirilen metotlar yoktur (eksik metot N ())
}
