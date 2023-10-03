/*
  - Go'nun sözdizimi C'ye benzer ancak başka birçok nokta daha vardır;
    örneğin ifadenin sonunda noktalı virgül yoktur veya
    değişken adından sonra bildirilen veri türü yoktur.
  - Farklı veri türlerine sahip değişkenler arasında hesaplamalar yaparken
    tür dönüşümüne ihtiyacımız var: Formül T (v) burada T veri türü ve v değerdir
  - Sabit bildirim, bir değişkeni bildirmeye benzer ancak const kullanarak ve
    kısaltılmış ":=" sözdizimini kullanmaz.
*/
package main

import "fmt"

const Pi = 3.14

var message string
var c, python, csharp bool
var i, j int32 = 1, 2

func main() {
	fmt.Println(message, c, python, csharp, i, j)

	k := 3
	i := 55               // int
	j := 67.8             // float 64
	sum := k + i + int(j) // type conversion
	fmt.Println(sum)      // output: 125
}
