package main

import (
	"fmt"
	"testing"
)

var words = []string{"super", "califragi", "listic", "expiali", "docius"}
var wlen = len(words)

func BenchmarkSwitch(b *testing.B) {
	m := 0
	for i := 0; i < b.N; i++ {
		switch words[i%wlen] {
		case "super":
			m++
		case "califragi":
			m++
		case "listic":
			m++
		case "expiali":
			m++
		case "docius":
			m++
		}
	}
	fmt.Println(m)
}

func BenchmarkIf(b *testing.B) {
	m := 0
	for i := 0; i < b.N; i++ {
		w := words[i%wlen]
		if w == "super" {
			m++
		} else if w == "califragi" {
			m++
		} else if w == "listic" {
			m++
		} else if w == "expiali" {
			m++
		} else if w == "docius" {
			m++
		}
	}
	fmt.Println(m)
}
