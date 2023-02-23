package service

import (
	"context"
	_ "github.com/lib/pq"
	"github.com/uptrace/bun"
)

type Service struct {
	DB  *bun.DB
	ctx context.Context
}

func New(DB *bun.DB, ctx context.Context) *Service {
	return &Service{DB: DB, ctx: ctx}
}
