package db

import (
	"context"
	"log"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/n4vxn/twitter-go/config"
)

type DB struct {
	Pool *pgxpool.Pool
}

func New(ctx context.Context, conf *config.Config) *DB {
	dbConf, err := pgxpool.ParseConfig(conf.Database.URL)
	if err != nil {
		log.Fatal("cannot parse postgres config: %v", err)
	}

	pool, err := pgxpool.ConnectConfig(ctx, dbConf)
	if err != nil {
		log.Fatal("error connecting to postgres: %v", err)
	}

	db := &DB{Pool: pool}

	db.Ping(ctx)

	return db
}

func (db *DB) Ping(ctx context.Context) {
	if err := db.Pool.Ping(ctx); err != nil {
		log.Fatal("cannot ping postgres; %v", err)
	}

	log.Println("postgres connected!")
}

func (db *DB) Close() {
	db.Pool.Close()
}