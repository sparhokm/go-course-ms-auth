package pg

import (
	"github.com/sparhokm/go-course-ms-auth/internal/client/db"

	"github.com/jackc/pgx/v5/pgxpool"
)

type pgClient struct {
	masterDBC db.DB
}

func New(dbc *pgxpool.Pool) db.Client {
	return &pgClient{
		masterDBC: NewDB(dbc),
	}
}

func (c *pgClient) DB() db.DB {
	return c.masterDBC
}

func (c *pgClient) Close() error {
	if c.masterDBC != nil {
		c.masterDBC.Close()
	}

	return nil
}
