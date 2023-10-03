/*
  - Class yerine struct kullanabiliriz. Ancak yapının yalnızca özellikleri vardır ve yöntemleri yoktur.
    Yöntemi diğer nesne yönelimli diller gibi uygulamak için,
    işlevi alıcı argümanı adı verilen özel bir parametreyle bildirmemiz gerekecek.
    Alıcı argümanı, func anahtar sözcüğünün ve işlev adının ortasındadır;
    bu işlevi bir yöntem olarak uygulamak için bir tür (genellikle bir yapı) belirtir.
*/
package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

// Creat method abs() for struct Vertex(receiver argument)
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
func func1() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())
}

type Student struct {
	name string
	age  int
}

func func2() {
	s1 := Student{"Steve", 30}
	s2 := Student{"Steve", 30}
	s3 := Student{"Job", 30}
	if s1 == s2 {
		fmt.Println("s1 = s2")
	} else {
		fmt.Println("s1 != s2")
	}
	if s1 == s3 {
		fmt.Println("s1 = s3")
	} else {
		fmt.Println("s1 != s3")
	}
	if s2 == s3 {
		fmt.Println("s2 = s3")
	} else {
		fmt.Println("s2 != s3")
	}
}

func main() {
	func1()
	func2()
}
