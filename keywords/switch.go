package keywords

import "fmt"

func expressionSwitches() int {
	fmt.Println("================expressionSwitches==================")
	switch x := seven(); {
	default:
		fmt.Println("default")
	case x == 1 || x == 3: // 1,3,7,7 will got compile error
		fmt.Println("B")
	case x == 2:
		fmt.Println("A")
	}

	/*Switch is generally faster than a long list of ifs because the compiler can generate a jump table. The longer the list, the better a switch statement is over a series of if statements*/

	return 1
}

func seven() int {
	return 7
}

func typeSwitches() {
	fmt.Println("typeSwitches")
}
