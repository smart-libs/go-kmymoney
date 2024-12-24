package reposecurities

import (
	"context"
	"database/sql"
	_ "embed"
	"errors"
	"fmt"
	"github.com/smart-libs/go-kmymoney/impl/sqlite/pkg/repo"
	"strconv"
)

type (
	GetNextSecurityIDService interface {
		NextSecurityID(ctx context.Context) string
	}

	getNextSecurityID struct {
		repo.FindOne
	}
)

//go:embed svc_get_last_security.sql
var sqlLastIDSQL string

func (g getNextSecurityID) NextSecurityID(ctx context.Context) string {
	var lastID sql.NullString
	if g.FindOne.Run(ctx, nil, []any{&lastID}) {
		if lastID.Valid {
			lastInt, err := strconv.Atoi(lastID.String[1:])
			if err != nil {
				panic(errors.Join(err, fmt.Errorf("cannot convert lastID to int: %s", lastID)))
			}
			lastInt++
			return fmt.Sprintf("E%06d", lastInt)
		}
	}
	return "E000001"
}

func NewGetNextSecurityIDService(provider repo.LoggerProvider) GetNextSecurityIDService {
	return getNextSecurityID{
		repo.FindOne{
			ErrMsg:         "getNextSecurityID:",
			SQLCmd:         sqlLastIDSQL,
			LoggerProvider: provider,
		},
	}
}
