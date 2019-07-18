package main

import "fmt"

func main() {
	for i, e := range "abcdefg" {
		switch e {
		case 'a':
			fmt.Println(i)
			fmt.Println("aaaaaaaaaaaaaa")
		case 'b':
			fmt.Println(i)
			fmt.Println("bbbbbbb")
		}

	}
}

// output:
// hello 222wo44rld
//
