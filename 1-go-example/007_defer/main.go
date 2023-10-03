/*
  - Gecikme (erteleme) akış kontrolünde oldukça yeni bir kavramdır.
    Bir komutun çağrılmasına izin verir ancak hemen yürütülmez ve
    çevredeki komutlar sonuç döndürene kadar gecikir.
  - defer anahtar sözcüğüyle çağrılan komutlar bir yığına konulur ve
    son giren ilk çıkar mekanizmasına göre çalışır.
*/
package main

import "fmt"

func main() {
	defer fmt.Println("World") // Delay print: "World"
	fmt.Println("Hello")       // Print: "Hello"
	// output : Hello World
}
