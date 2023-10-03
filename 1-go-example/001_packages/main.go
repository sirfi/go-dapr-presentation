/*
  - Go'dan yazılan tüm programlar "main" paketler tarafından oluşturulur ve
    çalıştırmak için kullanılan paketler "main" paketlerdir
  - Diğer paketleri kullanmak için içe aktarmamız gerekiyor,
    örneğin konsola bir metin yazdırmak istiyorsak fmt paketini kullanmalıyız
*/
package main

import "fmt"

func main() {
	fmt.Println("Hello, world!")
}
