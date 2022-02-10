package main

import "fmt"

func main() {
	// n := 15

	// for x := 1; x <= n; x++ {
	// 	if x%3 == 0 && x%5 == 0 {
	// 		fmt.Println("FizzBazz")
	// 	} else if x%5 == 0 {
	// 		fmt.Println("Bazz")
	// 	} else if x%3 == 0 {
	// 		fmt.Println("Fizz")
	// 	} else {
	// 		fmt.Println(x)
	// 	}
	// }

	test1 := 1221
	test2 := 12421
	test3 := 2212

	palindrom(test1)
	palindrom(test2)
	palindrom(test3)

}

func palindrom(data int) {
	// x := make(map[int]int)

	// for y := 0; y < 5; y++ {
	// 	x[data[y]] = data[y]
	// }

	// for idx,item := range x {
	// 	if
	// }

	//make map
	//cek panjang data
	// //mengisi data ke map kosong
	// for y := data; y >= 5; y++ {

	// 	// index berisi angka asli dan data berisi angka yg dibalik
	// 	x[data[y]] = data[y]
	// }
	fmt.Println(data)
}
