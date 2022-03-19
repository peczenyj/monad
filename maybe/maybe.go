package maybe

import "fmt"

// Maybe monad interface.
type Maybe[T any] interface {
	IsEmpty() bool
	Apply(f func(T))
	Map(f func(T) T) Maybe[T]
	Get() (T, error)
	GetOrElse(t T) T
}

// Of method create Some monad as typed Maybe[T]
func Of[T any](t T) Maybe[T] {
	return some[T]{t}
}

// None create None monad as typed Maybe[T]
func None[T any]() Maybe[T] {
	return none[T]{}
}

// Map function
func Map[T, R any](mt Maybe[T], f func(T) R) Maybe[R] {
	var response Maybe[R] = None[R]()

	mt.Apply(func(t T) {
		response = Of(f(t))
	})

	return response
}

// FlatMap function
func FlatMap[T, R any](mt Maybe[T], f func(T) Maybe[R]) Maybe[R] {
	var response Maybe[R] = None[R]()

	mt.Apply(func(t T) {
		response = f(t)
	})

	return response
}

type some[T any] struct{ t T }

func (s some[T]) IsEmpty() bool { return false }

func (s some[T]) Apply(f func(T)) {
	f(s.t)
}

func (s some[T]) Map(f func(T) T) Maybe[T] {
	return some[T]{f(s.t)}
}

func (s some[T]) Get() (T, error) {
	return s.t, nil
}

func (s some[T]) GetOrElse(T) T {
	return s.t
}

type none[T any] struct{}

func (n none[T]) IsEmpty() bool {
	return true
}

func (n none[T]) Apply(func(T)) {}

func (n none[T]) Map(func(T) T) Maybe[T] {
	return n
}

func (n none[T]) Get() (t T, err error) {
	return t, fmt.Errorf("There is no object here, this is a None[%T]", t)
}

func (n none[T]) GetOrElse(t T) T {
	return t
}
