package keywords

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
		for cnt, j := range "AKQJ098765432" {
			gotcards = append(gotcards, string(i)+string(j))
			if cnt == n-1 { // I wanna get n cards
				break OutterLoop // break the Outter loop using lable
			}
		}
	}
	// fmt.Println(gotcards)
	return gotcards // finaly i got 28 show-hand cards
}

func playPoker(t string, n int) []string {
	if t == "showHand" {
		return showHand()
	} else if t == "topCards" {
		return topCards(n)
	} else {
		panic("wrong poker game type:  showHand or topCards")
	}
}
