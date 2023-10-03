/*
  - Go'daki koşullu ifade için if, else, switch ve for döngüsünü kullanılır,
    ayrıca yuvarlak parantezlere de ihtiyacımız yoktur.
  - if ifadesi ile koşullu ifadedeki değişkeni deklare edebiliriz ve
    bu değişken sadece if veya else ifadesinin bloğunda çalışacaktır.
*/
package main

import (
	"fmt"
	"math"
)

func pow(x, n, limit float64) float64 {
	if v := math.Pow(x, n); v < limit {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, limit)
	}
	return limit
}

func main() {
	pow(3, 2, 10)
}
