package storage

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/madeinheaven91/patchouli/internal/shared"
)

func Delete(table string, key string, value string, ctx context.Context) error {
	conn, err := pool.Acquire(ctx)
	if err != nil {
		shared.LogError(err)
		return err
	}
	defer conn.Release()

	query := fmt.Sprintf(`DELETE FROM %s WHERE %s = $1`,
		pgx.Identifier{table}.Sanitize(),
		pgx.Identifier{key}.Sanitize())
	cmd, err := conn.Exec(ctx, query, value)
	if err != nil {
		shared.LogError(err)
		return err
	}
	if cmd.RowsAffected() == 0 {
		err = errors.New("no row found to delete")
	}
	return err
}
