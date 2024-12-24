package repo

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
)

type FindOne struct {
	ErrMsg string
	SQLCmd string
	LoggerProvider
}

func (s *FindOne) Run(ctx context.Context, input, output []any) bool {
	endpoint := EndpointFromContext(ctx)

	// insert
	row := endpoint.GetDB().QueryRowContext(ctx, s.SQLCmd, input...)
	if row.Err() != nil {
		panic(errors.Join(row.Err(), fmt.Errorf(s.ErrMsg, input...)))
	}

	err := row.Scan(output...)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false
		}
		panic(errors.Join(err, fmt.Errorf(s.ErrMsg, input...)))
	}
	return true
}
