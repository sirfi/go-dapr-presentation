/*
  - Fonksiyonu func anahtar sözcüğünü kullanarak bildirin; giriş parametresinin,
    parametre adından sonra veri türünü de bildirdiğine dikkat edin.
  - Go'nun özel bir özelliği, fonksiyonun birçok sonuç döndürebilmesidir.
*/
package main

func add(x int, y int) int {
	return x + y
}

// fonksiyon iki değer döndürüyor
func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	total := add(3, 4)
	println(total)

	a, b := swap("hello", "world")
	println(a, b)
}
