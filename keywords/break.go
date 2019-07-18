package keywords

import "fmt"

func showHand() []string {
	gotcards := make([]string, 0)
	for _, i := range "♥♠♣♦" {
		for _, j := range "AKQJ098765432" {
			gotcards = append(gotcards, string(i)+string(j))
			if j == '8' { // I wanna play show-hand
				break // break the inner loop
			}
		}
	}
	// fmt.Println(gotcards)
	return gotcards // finaly i got 28 show-hand cards
}

func topCards(n int) []string {
	gotcards := make([]string, 0)
OutterLoop:
	for _, i := range "♥♠♣♦" {
		for _, j := range "AKQJ098765432" {
			gotcards = append(gotcards, string(i)+string(j))
			if len(gotcards) == n { // I wanna get n cards
				break OutterLoop // break the Outter loop using lable
			}
		}
	}
	// fmt.Println(gotcards)
	return gotcards // finaly i got 28 show-hand cards
}

func gBoom(s rune) []string {
	boom := make([]string, 0)
	for _, i := range "♥♠♣♦" {
	second:
		for _, j := range "AKQJ098765432" {
			fmt.Println(string(i) + string(j))
			switch j {
			case s:
				boom = append(boom, string(i)+string(j))
				// single break without label only break the switch  not the second for loop
				// break only break out to the innermost "for" "switch" "select"
				break second
			}
		}
	}
	return boom
}
