package access

import (
	"context"
	"log"

	sq "github.com/Masterminds/squirrel"

	"github.com/sparhokm/go-course-ms-auth/internal/model/user"
	"github.com/sparhokm/go-course-ms-auth/pkg/client/db"
)

const (
	tableName = "endpoint_accesses"

	roleColumn     = "role"
	endpointColumn = "endpoint"
)

type repo struct {
	dbc db.Client
}

func NewRepository(dbc db.Client) *repo {
	return &repo{dbc: dbc}
}

func (r *repo) Get(ctx context.Context, endpoint string) ([]user.Role, error) {
	builder := sq.Select(roleColumn).
		PlaceholderFormat(sq.Dollar).
		From(tableName).
		Where(sq.Eq{endpointColumn: endpoint})

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := r.dbc.DB().QueryContext(ctx, db.Query{Name: "access.get", QueryRaw: query}, args...)
	if err != nil {
		return nil, err
	}

	var roles []user.Role
	for rows.Next() {
		var role user.Role

		err = rows.Scan(&role)
		if err != nil {
			log.Printf("failed to scan access: %v", err)
		}

		roles = append(roles, role)

		log.Printf("access endpoint: %s\n", endpoint)
	}

	return roles, nil
}
