package main

import "fmt"

func add(a, b int) int {
	return a + b
}

func main() {
	// Değişken tanımlama ve kullanma
	var message string = "Hello, world!"
	// message := "Hello, world!"
	fmt.Println(message)

	// Döngüler
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// Diziler
	numbers := []int{1, 2, 3, 4, 5}
	fmt.Println(numbers)

	// Fonksiyonlar
	result := add(3, 4)
	fmt.Println(result)

}
