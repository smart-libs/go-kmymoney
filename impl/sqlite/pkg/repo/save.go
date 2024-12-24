package repo

import (
	"context"
	"errors"
	"fmt"
)

type Save struct {
	ErrMsg string
	SQLCmd string
	LoggerProvider
}

func (s *Save) Run(ctx context.Context, input ...any) (affected int64) {
	endpoint := EndpointFromContext(ctx)

	// insert
	result, err := endpoint.GetDB().Exec(s.SQLCmd, input...)
	if err != nil {
		panic(errors.Join(err, fmt.Errorf(s.ErrMsg, input...)))
	}

	affected, err = result.RowsAffected()
	if err != nil {
		panic(errors.Join(err, fmt.Errorf(s.ErrMsg, input...)))
	}

	//s.LoggerProvider(ctx).Info("Save", "affected", affected)
	return affected
}
