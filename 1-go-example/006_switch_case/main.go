/*
  - Anahtardaki ifadede bir sabit kullanılamaz.
    Her durumda break komutlarına gerek yoktur.
    Bu nedenle yalnızca ilk karşılanan durum çalıştırılır (yukarıdan aşağıya).
  - Bir durumda birden fazla koşulu kullanabilir veya
    bir sonraki komutun devam etmesine izin vermek için
    fallthrough anahtar sözcüğünü kullanabilirsiniz.
*/
package main

import "fmt"

func func1() {
	num := 10
	switch {
	case num >= 0 && num <= 50:
		fmt.Printf("%d < 50 \n", num) // print: 10 < 50
		fallthrough
	case num < 100:
		fmt.Printf("%d < 100 \n", num) //print: 10 < 100
	default:
		fmt.Printf("I don't know %d", num)
	}

}

func func2() {
	switch num := 10; {
	case num < 50:
		fmt.Printf("%d < 50\n", num) // print: 10 < 50
	case num < 100:
		fmt.Printf("%d < 100\n", num) // will not be executed
	default:
		fmt.Printf("I don't know %d", num)
	}

}

func main() {
	func1()
	func2()
}
