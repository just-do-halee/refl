package refl

import (
	"testing"

	"github.com/just-do-halee/lum"
)

func TestNameEqAny(t *testing.T) {
	type Args struct {
		name string
		any  []string
	}
	lum.Batch[Args, bool]{
		{
			Name: "1",
			Args: Args{"a", []string{"a", "b"}},
			Pass: func(c *lum.Ctx[Args, bool]) {
			},
		},
	}.Run(t, "Sume", func(args Args) bool {
		return true
	})
}
