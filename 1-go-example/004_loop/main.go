/*
  - Go'da yalnızca 1 tür döngü kullanılır: for döngüsü.
    Kullanımı diğer dillere benzer ancak değişkenlerin bildirimi, tekrarlanan koşullar,
    yuvarlak parantez içine alınmasına gerek yoktur
  - for döngüsü yalnızca yinelenen bir koşul kullanıldığında, diğer dillerindeki while gibi davranır
*/
package main

import "fmt"

func func1() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
}

func func2() {
	sum := 0
	for sum < 10 {
		sum += sum
	}
	fmt.Println(sum)
}

func main() {
	func1()
	func2()
}
