package deletedUser

import (
	"context"
	"fmt"

	sq "github.com/Masterminds/squirrel"

	"github.com/sparhokm/go-course-ms-auth/pkg/client/db"
)

const (
	tableName = "deleted_users"
)

type repo struct {
	dbc db.Client
}

func NewRepository(dbc db.Client) *repo {
	return &repo{dbc: dbc}
}

func (r *repo) Save(ctx context.Context, id int64) error {
	builder := sq.Insert(tableName).
		PlaceholderFormat(sq.Dollar).
		Values(id)

	query, args, err := builder.ToSql()
	if err != nil {
		return err
	}

	_, err = r.dbc.DB().ExecContext(ctx, db.Query{Name: "deletedUser.save", QueryRaw: query}, args...)
	if err != nil {
		return fmt.Errorf("failed to insert deleted user: %v", err)
	}

	return nil
}
