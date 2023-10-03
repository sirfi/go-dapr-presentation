/*
  - Map, anahtar - değer olarak saklanan bir dizi öğedir. Map'teki anahtar karşılaştırılabilir ve
    yinelenmeyen bir veri türüne sahiptir. Bir harita oluşturmak için make() fonsiyonunu aşağıdaki formülle kullanın:
    "make(map[anahtar tipi]değer tipi)"
  - Map'teki elemanı silmek için delete() fonksiyonunu kullanarak silinecek elemanın anahtarını sağlıyoruz.
  - Map'teki öğeye erişmek için öğenin tuşuyla haritayı çağırıyoruz.
    Eğer o anahtar yoksa sıfır değerini alırız (veri tipine bağlı olarak).
    İlk değer yukarıdaki örnekle aynıdır, ikinci değer eğer element haritada ise true,
    element mevcut değilse false olacaktır.
*/
package main

import "fmt"

func main() {
	var demoMap map[string]int
	if demoMap == nil {
		fmt.Println("Map has nil value.")
		demoMap = make(map[string]int)
	}
	demoMap["one"] = 1
	languages := make(map[string]float32)
	languages["go"] = 0.63
	languages["java"] = 1.03
	delete(languages, "go")
	m := make(map[string]int)
	m["Answer"] = 42
	delete(m, "Answer")
	fmt.Println(m["Answer"])
	v, ok := m["Answer"]
	fmt.Println("Value of element: ", v) // v= 0
	fmt.Println("Exist or not ", ok)     // ok = false

}
