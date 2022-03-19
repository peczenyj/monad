package either_test

import (
	"fmt"

	"github.com/peczenyj/monad/either"
)

func ExampleEither() {
	x := either.Right[error](1)

	y := either.Left[error, int](fmt.Errorf("ops"))

	fmt.Println(x.GetOrElse(0), y.GetOrElse(0))
	// Output: 1 0
}
