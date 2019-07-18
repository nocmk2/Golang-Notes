package keywords

import "fmt"

func expressionSwitches() int {
	fmt.Println("================expressionSwitches==================")
	switch x := seven(); {
	default:
		fmt.Println("default")
	case x == 1 || x == 3: // 1,3,7,7 will gotcompile error
		fmt.Println("B")
	case x == 2:
		fmt.Println("A")
	}
	return 1
}

func seven() int {
	return 7
}

func typeSwitches() {
	fmt.Println("typeSwitches")
}
