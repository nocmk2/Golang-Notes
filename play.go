package main

import (
	"fmt"
	"reflect"
)

// GCD greatest common divisor
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return b
}

func main() {
	// naturals := make(chan int)
	// squares := make(chan int)

	// go func() {
	// 	for x := 0; x < 4; x++ {
	// 		time.Sleep(time.Second)
	// 		naturals <- x
	// 	}
	// 	close(naturals)
	// }()

	// go func() {
	// 	for {
	// 		for x := range naturals {
	// 			squares <- x * x
	// 		}

	// 	}
	// }()

	// for x := range squares {
	// 	fmt.Println(x)
	// }
	// fmt.Println(GCD(3, 5))

	a := 3.5566
	b := fmt.Sprintf("%f", a)
	fmt.Println(reflect.TypeOf(b))

}
