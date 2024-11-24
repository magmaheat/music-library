package postgres

import (
	"context"
	"fmt"
	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5/pgxpool"
	log "github.com/sirupsen/logrus"
	"time"
)

type Postgres struct {
	maxPoolSize  int
	connAttempts int
	connTimeout  time.Duration

	Builder squirrel.StatementBuilderType
	Pool    *pgxpool.Pool
}

func NewPostgres() *Postgres {
	return &Postgres{
		maxPoolSize:  1,
		connAttempts: 20,
		connTimeout:  time.Second,
		Builder:      squirrel.StatementBuilderType{},
		Pool:         nil,
	}
}

func New(storagePath string, opts ...Options) (*Postgres, error) {
	pg := NewPostgres()

	for _, opt := range opts {
		opt(pg)
	}

	pg.Builder = squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)

	poolConfig, err := pgxpool.ParseConfig(storagePath)
	if err != nil {
		return nil, fmt.Errorf("pgdb - New - pgxpool.ParseConfig: %v", err)
	}

	poolConfig.MaxConns = int32(pg.maxPoolSize)

	for pg.connAttempts > 0 {
		pg.Pool, err = pgxpool.NewWithConfig(context.Background(), poolConfig)
		if err == nil {
			break
		}

		log.Infof("Postgres is trying connect: %d", pg.connAttempts)
		time.Sleep(pg.connTimeout)
		pg.connAttempts--
	}

	if err != nil {
		return nil, fmt.Errorf("pgdb - New - pgxpool.WithConfig: %v", err)
	}

	return pg, nil
}

func (p *Postgres) Close() {
	if p.Pool != nil {
		p.Pool.Close()
	}
}
