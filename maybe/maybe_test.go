package maybe_test

import (
	"fmt"

	"github.com/peczenyj/monad/maybe"
)

func ExampleMaybe() {
	x := maybe.Of(1)

	y := maybe.None[int]()

	fmt.Println(x.GetOrElse(0), y.GetOrElse(0))
	// Output: 1 0
}
