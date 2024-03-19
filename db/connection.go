package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/iuan95/apigo/config"
)


func Connection(ctx context.Context) error{
	cfg:= config.GetDbConfig()
	conn, err:= pgxpool.New(ctx, cfg)
	if  err != nil {
		conn.Close()
		return err
	}
	DB=conn
	return nil
}