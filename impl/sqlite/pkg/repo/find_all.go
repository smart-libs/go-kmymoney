package repo

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
)

type FindAll struct {
	ErrMsg string
	SQLCmd string
	LoggerProvider
}

func (s *FindAll) Run(ctx context.Context, inputs ...any) *sql.Rows {
	endpoint := EndpointFromContext(ctx)

	// insert
	rows, err := endpoint.GetDB().QueryContext(ctx, s.SQLCmd, inputs...)
	if err != nil {
		panic(errors.Join(err, fmt.Errorf(s.ErrMsg, inputs...)))
	}

	return rows
}
