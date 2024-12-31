package repo

import (
	"github.com/smart-libs/go-kmymoney/spec/lib/pkg/cursor"
	"io"
)

type (
	// Rows is a facade for sql.Rows to allow unit tests in sqlCursor
	Rows interface {
		Next() bool
		Scan(dest ...any) error
		io.Closer
	}

	sqlCursor[T any] struct {
		rows  Rows
		fetch func(rows Rows) T
	}
)

// NewSQLCursor creates a new Cursor wrapping sql.Rows
func NewSQLCursor[T any](rows Rows, fetch func(rows Rows) T) cursor.Cursor[T] {
	return &sqlCursor[T]{
		rows:  rows,
		fetch: fetch,
	}
}

func (c *sqlCursor[T]) Next() bool {
	if c.rows == nil {
		return false
	}
	return c.rows.Next()
}

func (c *sqlCursor[T]) Fetch() T {
	if c.rows == nil {
		panic("FetchInto called on nil sql.Rows")
	}
	return c.fetch(c.rows)
}

func (c *sqlCursor[T]) Close() error {
	if c.rows == nil {
		return nil
	}
	return c.rows.Close()
}
