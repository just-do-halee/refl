package refl

import (
	"strings"
	"testing"
)

type Field[A any, R comparable] struct {
	Name string
	Args A
	Want R
}

func (f Field[A, R]) Run(t *testing.T, fnName string, fn func(a A) R) {
	t.Run(f.Name, func(t *testing.T) {
		if got := fn(f.Args); got != f.Want {
			t.Errorf("%s(%v) => %v, want %v", fnName, f.Args, got, f.Want)
		}
	})
}

type Batch[A any, R comparable] []Field[A, R]

func (b Batch[A, R]) Run(t *testing.T, fn func(a A) R) {
	testName, _, _ := strings.Cut(t.Name(), "/")
	testName = testName[4:]
	for _, test := range b {
		test.Run(t, testName, fn)
	}
}

func TestNameEqAny(t *testing.T) {
	type Args struct {
		a, b any
	}
	b := Batch[Args, bool]{
		{
			Name: "it should be true",
			Args: Args{1, 1},
			Want: true,
		},
		{
			Name: "It should be false",
			Args: Args{1, true},
			Want: false,
		},
	}
	b.Run(t, func(a Args) bool {
		return NameEqAny(a.a, a.b)
	})
}
