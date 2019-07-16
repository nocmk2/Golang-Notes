package keywords

import "testing"

func Test_playPoker(t *testing.T) {
	tests := []struct {
		name  string
		input string
		num   int
		want  int
	}{
		// TODO: Add test cases.
		{"case1", "showHand", 0, 28},
		{"case2", "topCars", 10, 10},
		// {"case2", 0, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := len(playPoker(tt.input, tt.num)); got != tt.want {
				t.Errorf("len(playPoker()) = %v, want %v for input %v", got, tt.want, tt.input)
			}
		})
	}
}
