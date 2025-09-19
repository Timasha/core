package postgresql

import (
	"context"
	"fmt"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
)

type Connector struct {
	cfg Config
	*sqlx.DB
}

func NewConnector(cfg Config) *Connector {
	return &Connector{
		cfg: cfg,
		DB:  new(sqlx.DB),
	}
}

func (c *Connector) Start(ctx context.Context) error {
	db, err := sqlx.Connect(
		"postgres",
		fmt.Sprintf(
			"postgres://%v:%v@%v:%v/%v?sslmode=%v",
			c.cfg.Login,
			c.cfg.Password,
			c.cfg.Address,
			c.cfg.Port,
			c.cfg.DBName,
			c.cfg.sslMode(),
		),
	)
	if err != nil {
		return err
	}

	db.SetMaxOpenConns(c.cfg.MaxOpenConns)
	db.SetMaxIdleConns(c.cfg.MaxIdleConns)

	*c.DB = *db

	return nil
}

func (c *Connector) Stop(ctx context.Context) error {
	return c.Close()
}

func (c *Connector) GetName() string {
	return "postgres"
}

func (c *Connector) IsEnabled() bool {
	return c.cfg.IsEnabled
}
