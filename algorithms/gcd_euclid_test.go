package main

import "testing"

func TestGCD(t *testing.T) {
	type args struct {
		m int
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"case1", args{544, 119}, 17},
		{"case2", args{10, 8}, 2},
		{"case3", args{119, 544}, 17},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GCD(tt.args.m, tt.args.n); got != tt.want {
				t.Errorf("GCD(%v,%v) = %v, want %v", tt.args.m, tt.args.n, got, tt.want)
			}
		})
	}
}
