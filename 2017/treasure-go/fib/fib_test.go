package main

import (
	"testing"
)

func TestFib(t *testing.T) {
	type Case struct {
		n, out int
	}
	cases := []Case{
		{0, 0},
		{1, 1},
		{10, 55},
	}
	for i, c := range cases {
		if got := fib(c.n); got != c.out {
			t.Errorf("#%d: fib(%d) want %d, got %d\n", i, c.n, c.out, got)
		}
	}
}
