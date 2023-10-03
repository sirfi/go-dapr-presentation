package main

import "fmt"

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}
func main() {
	f("direkt")
	// Goroutine’lerde f() fonksiyonu go func() şeklinde çağrılır.
	go f("goroutine")
	go func(msg string) {
		fmt.Println(msg)
	}("çalışıyor")

	fmt.Scanln()
	fmt.Println("tamamlandı")

	/*
	   çıktı
	   -------
	   direkt : 0
	   direkt : 1
	   direkt : 2
	   goroutine : 0
	   çalışıyor
	   goroutine : 1
	   goroutine : 2
	   <enter>
	   tamamlandı
	*/
}
