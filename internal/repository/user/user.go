package user

import (
	"context"
	"fmt"
	"time"

	sq "github.com/Masterminds/squirrel"
	model "github.com/sparhokm/go-course-ms-auth/internal/model/user"
	repoModel "github.com/sparhokm/go-course-ms-auth/internal/repository/user/model"
	"github.com/sparhokm/go-course-ms-auth/pkg/client/db"
)

const (
	tableName = "users"

	idColumn           = "id"
	nameColumn         = "name"
	emailColumn        = "email"
	passwordHashColumn = "password_hash"
	roleColumn         = "role"
	updatedAtColumn    = "updated_at"
)

type repo struct {
	db db.Client
}

func NewRepository(db db.Client) *repo {
	return &repo{db: db}
}

func (r *repo) Save(ctx context.Context, info *model.Info, passwordHash string) (int64, error) {
	builder := sq.Insert(tableName).
		PlaceholderFormat(sq.Dollar).
		Columns(nameColumn, emailColumn, passwordHashColumn, roleColumn).
		Values(info.Name, info.Email, passwordHash, info.Role.GetCode()).
		Suffix("RETURNING id")

	query, args, err := builder.ToSql()
	if err != nil {
		return 0, err
	}
	var id int64
	err = r.db.DB().ScanOneContext(ctx, &id, db.Query{Name: "user.save", QueryRaw: query}, args...)
	if err != nil {
		return 0, fmt.Errorf("failed to save user: %v", err)
	}

	return id, nil
}

func (r *repo) Update(ctx context.Context, id int64, info *model.Info) error {
	builder := sq.Update(tableName).
		PlaceholderFormat(sq.Dollar).
		Set(nameColumn, info.Name).
		Set(emailColumn, info.Email).
		Set(roleColumn, info.Role.GetCode()).
		Set(updatedAtColumn, time.Now()).
		Where(sq.Eq{idColumn: id})

	query, args, err := builder.ToSql()
	if err != nil {
		return err
	}

	_, err = r.db.DB().ExecContext(ctx, db.Query{Name: "user.update", QueryRaw: query}, args...)
	if err != nil {
		return fmt.Errorf("failed to update user: %v", err)
	}

	return nil
}

func (r *repo) Delete(ctx context.Context, id int64) error {
	builder := sq.Delete(tableName).
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{idColumn: id})

	query, args, err := builder.ToSql()
	if err != nil {
		return err
	}

	_, err = r.db.DB().ExecContext(ctx, db.Query{Name: "user.delete", QueryRaw: query}, args...)
	if err != nil {
		return fmt.Errorf("failed to delete user: %v", err)
	}

	return nil
}

func (r *repo) Get(ctx context.Context, id int64) (*model.User, error) {
	builder := sq.Select("*").
		PlaceholderFormat(sq.Dollar).
		From(tableName).
		Where(sq.Eq{idColumn: id}).
		Limit(1)

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	var repoUser repoModel.User
	err = r.db.DB().ScanOneContext(ctx, &repoUser, db.Query{Name: "user.get", QueryRaw: query}, args...)
	if err != nil {
		return nil, err
	}

	return ToUserFromRepo(&repoUser)
}
