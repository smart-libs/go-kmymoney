package repo

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockRows struct {
	mock.Mock
}

func (m *mockRows) Next() bool {
	args := m.Called()
	return args.Bool(0)
}

func (m *mockRows) Scan(dest ...any) error {
	args := m.Called(dest)
	return args.Error(0)
}

func (m *mockRows) Close() error {
	args := m.Called()
	return args.Error(0)
}

func TestSQLCursor(t *testing.T) {
	t.Run("Next returns true when rows have more data", func(t *testing.T) {
		mockRows := &mockRows{}
		mockRows.On("Next").Return(true).Once()

		cursor := NewSQLCursor(mockRows, func(rows Rows) int { return 0 })
		assert.True(t, cursor.Next())

		mockRows.AssertExpectations(t)
	})

	t.Run("Next returns false when rows have no more data", func(t *testing.T) {
		mockRows := &mockRows{}
		mockRows.On("Next").Return(false).Once()

		cursor := NewSQLCursor(mockRows, func(rows Rows) int { return 0 })
		assert.False(t, cursor.Next())

		mockRows.AssertExpectations(t)
	})

	t.Run("Fetch returns correct value using fetch function", func(t *testing.T) {
		mockRows := &mockRows{}
		mockRows.On("Scan", mock.Anything).Run(func(args mock.Arguments) {
			dest := args.Get(0).([]any)
			*dest[0].(*int) = 42
		}).Return(nil).Once()

		cursor := NewSQLCursor(mockRows, func(rows Rows) int {
			var value int
			assert.NoError(t, rows.Scan(&value))
			return value
		})

		mockRows.On("Next").Return(true).Once()
		assert.True(t, cursor.Next())

		value := cursor.Fetch()
		assert.Equal(t, 42, value)

		mockRows.AssertExpectations(t)
	})

	t.Run("Fetch panics on nil rows", func(t *testing.T) {
		cursor := NewSQLCursor(nil, func(rows Rows) int { return 0 })
		assert.Panics(t, func() {
			cursor.Fetch()
		})
	})

	t.Run("Close closes the rows", func(t *testing.T) {
		mockRows := &mockRows{}
		mockRows.On("Close").Return(nil).Once()

		cursor := NewSQLCursor(mockRows, func(rows Rows) int { return 0 })
		assert.NoError(t, cursor.Close())

		mockRows.AssertExpectations(t)
	})

	t.Run("Close handles nil rows", func(t *testing.T) {
		cursor := NewSQLCursor(nil, func(rows Rows) int { return 0 })
		assert.NoError(t, cursor.Close())
	})
}
