package cursor

import "io"

type Cursor[T any] interface {
	Next() bool
	Fetch() T
	io.Closer
}
