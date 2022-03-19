package either

import "fmt"

// Either monad interface
type Either[L, R any] interface {
	IsLeft() bool
	IsRight() bool
	Fold(func(L), func(R))
	Apply(func(R))
	FlatMap(func(R) Either[L, R]) Either[L, R]
	GetOrElse(R) R
	Get() (R, error)
}

// Left constructor.
func Left[L, R any](l L) Either[L, R] {
	return left[L, R]{l}
}

// Right constructor.
func Right[L, R any](r R) Either[L, R] {
	return right[L, R]{r}
}

// Fold function
func Fold[L, R, X any](either Either[L, R], leftFunction func(L) X, rightFunction func(R) X) X {
	var x X

	either.Fold(func(l L) {
		x = leftFunction(l)
	}, func(r R) {
		x = rightFunction(r)
	})

	return x
}

// Map function
func Map[L, R, R1 any](either Either[L, R], f func(R) R1) Either[L, R1] {
	var response Either[L, R1]
	either.Fold(func(l L) {
		response = left[L, R1]{l}
	}, func(r R) {
		r1 := f(r)
		response = right[L, R1]{r1}
	})
	return response
}

type left[L, R any] struct{ l L }

func (l left[L, R]) IsLeft() bool { return true }

func (l left[L, R]) IsRight() bool { return false }

func (l left[L, R]) Fold(lf func(L), rf func(R)) { lf(l.l) }

func (l left[L, R]) Apply(func(R)) {}

func (l left[L, R]) FlatMap(func(R) Either[L, R]) Either[L, R] {
	return l
}

func (l left[L, R]) GetOrElse(r R) R {
	return r
}

func (l left[L, R]) Get() (r R, err error) {
	return r, fmt.Errorf("There is no Right[%T] here, this is a Left[%T]", r, l.l)
}

type right[L, R any] struct{ r R }

func (r right[L, R]) IsLeft() bool { return false }

func (r right[L, R]) IsRight() bool { return true }

func (r right[L, R]) Fold(lf func(L), lr func(R)) { lr(r.r) }

func (r right[L, R]) Apply(f func(R)) { f(r.r) }

func (r right[L, R]) FlatMap(f func(R) Either[L, R]) Either[L, R] {
	return f(r.r)
}

func (r right[L, R]) GetOrElse(R) R {
	return r.r
}

func (r right[L, R]) Get() (R, error) {
	return r.r, nil
}
