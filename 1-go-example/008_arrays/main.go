/*
  - Go'daki dizi diğer dillere benzer, ancak sabit bir boyuta sahiptir ve
    tüm öğelerin aynı veri tipinde olması gerekir.
  - Diğer çoğu dilin aksine, Go'daki Array bir referans tipi değil,
    bir değer tipidir. Değerini yeni bir değişkene atarken,
    eski Dizinin bir kopyasını oluşturur ve
    yeni Dizide yapılan herhangi bir değişiklik eski Diziyi etkilemez.
*/
package main

import "fmt"

func main() {
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)
	primes_copy := primes
	primes_copy[0] = 1
	fmt.Println(primes_copy)
}
