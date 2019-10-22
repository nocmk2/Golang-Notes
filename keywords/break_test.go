package keywords

import (
	"reflect"
	"testing"
)

func Test_gBoom(t *testing.T) {
	type args struct {
		s rune
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{"case1", args{'0'}, []string{"♥K", "♠K", "♣K", "♦K"}},
		// {"case2", args{'Q'}, []string{"♥Q", "♠Q", "♣Q", "♦Q"}},
		// {"case3", args{'9'}, []string{"♥9", "♠9", "♣9", "♦9"}},
		// {"case3", args{'0'}, []string{"♥0", "♠0", "♣0", "♦0"}},
		// {"case4", args{'2'}, []string{"♥2", "♠2", "♣2", "♦2"}},
		// {"case5", args{'1'}, []string{}}, // no 1 in cards
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := gBoom(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("gBoom() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_topCards(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"case1", args{2}, 2},
		{"case1", args{20}, 20},
		{"case1", args{100}, 52},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := topCards(tt.args.n); len(got) != tt.want {
				t.Errorf("topCards() = %v, len = %v,  want %v", got, len(got), tt.want)
			}
		})
	}
}

func Test_showHand(t *testing.T) {
	tests := []struct {
		name string
		want []string
	}{
		{"case1", []string{"♥A", "♥K", "♥Q", "♥J", "♥0", "♥9", "♥8", "♠A", "♠K", "♠Q", "♠J", "♠0", "♠9", "♠8", "♣A", "♣K", "♣Q", "♣J", "♣0", "♣9", "♣8", "♦A", "♦K", "♦Q", "♦J", "♦0", "♦9", "♦8"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := showHand(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("showHand() = %v, want %v", got, tt.want)
			}
		})
	}
}
