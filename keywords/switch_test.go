package keywords

import "testing"

func Test_expressionSwitches(t *testing.T) {
	tests := []struct {
		name string
		want int
	}{
		{"case1", 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := expressionSwitches(); got != tt.want {
				t.Errorf("expressionSwitches() = %v, want %v", got, tt.want)
			}
		})
	}
}
