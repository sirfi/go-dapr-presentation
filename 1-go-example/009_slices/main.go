/*
  - Slice, Diziye bir referanstır; Dizinin bir kısmını (veya tamamını) açıklar.
    Dinamik boyuta sahip olduğundan genellikle Array'den daha fazla kullanılır.
  - Dizideki öğeyi bulmak için 2 dizin (düşük ve yüksek) sağlanarak bir Diziden slice oluşturulabilir.
  - Bir slice'ın iki özelliği olacaktır: uzunluk ve kapasite. Uzunluk, Dilimdeki öğelerin sayısıdır;
    kapasite ise Dilimin ifade ettiği Dizideki öğelerin sayısıdır (Dilimdeki ilk öğeden başlayarak).
    Uzunluğu elde etmek için len()'ı, kapasiteyi elde etmek için ise cap()'ı kullanırız.
  - Slice'lar yalnızca Diziye referans olduğundan, slice'ların değerini değiştirmek
    Dizinin değerini değiştirecektir. Bir Diziye başvuran birden fazla slice varsa,
    bir slice'ın değerini değiştirmek diğer slice'ların değerini değiştirebilir.
*/
package main

import "fmt"

func main() {
	s := []int{2, 3, 5, 7, 11, 13}
	s = s[0:0] // s = [], len(s) = 0, cap(s) = 6
	fmt.Println(s)
	s = s[0:4] // s = [2, 3, 5, 7], len(s) = 4, cap(s) = 6
	fmt.Println(s)
	s = s[2:4] // s = [5, 7], len(s) = 2, cap(s) = 4
	fmt.Println(s)
}
