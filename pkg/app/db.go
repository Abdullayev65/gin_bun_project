package app

import (
	"context"
	"database/sql"
	"github.com/Abdullayev65/gin_bun_project/pkg/models"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"github.com/uptrace/bun/extra/bundebug"
)

var resetModel = false

func database() (*bun.DB, context.Context) {
	dsn := "postgres://postgres:@localhost:5432/postgres?sslmode=disable"
	pgdb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))

	db := bun.NewDB(pgdb, pgdialect.New())

	db.AddQueryHook(bundebug.NewQueryHook(bundebug.WithVerbose(true)))

	ctx := context.Background()

	ms := []interface{}{(*models.User)(nil),
		(*models.Attachment)(nil),
		(*models.PostAttachment)(nil),
		(*models.Post)(nil),
		(*models.Comment)(nil),
	}
	if resetModel {
		db.ResetModel(ctx, ms...)
	} else {
		db.RegisterModel(ms...)
	}
	return db, ctx
}
