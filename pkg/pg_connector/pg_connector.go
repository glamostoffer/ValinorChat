package pg_connector

import (
	"context"
	"fmt"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
	"time"
)

type Config struct {
	User            string        `yaml:"user"`
	Port            string        `yaml:"port"`
	DBName          string        `yaml:"dbname"`
	Password        string        `yaml:"password"`
	Host            string        `yaml:"host"`
	Driver          string        `yaml:"driver"`
	SSLMode         string        `yaml:"sslMode"`
	MaxOpenConns    int           `yaml:"max_open_conns"`
	ConnMaxLifetime time.Duration `yaml:"conn_max_lifetime"`
	MaxIdleConns    int           `yaml:"max_idle_conns"`
	ConnMaxIdleTime time.Duration `yaml:"conn_max_idle_time"`
}

type Connector struct {
	cfg Config
	*sqlx.DB
}

func New(cfg Config) *Connector {
	return &Connector{
		cfg: cfg,
	}
}

func (c *Connector) Start(_ context.Context) error {
	connectionString := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		c.cfg.Host,
		c.cfg.Port,
		c.cfg.User,
		c.cfg.Password,
		c.cfg.DBName,
		c.cfg.SSLMode,
	)

	db, err := sqlx.Open(c.cfg.Driver, connectionString)
	if err != nil {
		return err
	}

	db.SetMaxOpenConns(c.cfg.MaxOpenConns)
	db.SetConnMaxLifetime(c.cfg.ConnMaxLifetime)
	db.SetMaxIdleConns(c.cfg.MaxIdleConns)
	db.SetConnMaxIdleTime(c.cfg.ConnMaxIdleTime)

	if err = db.Ping(); err != nil {
		return err
	}

	c.DB = db

	return nil
}

func (c *Connector) Stop(_ context.Context) error {
	return c.Close()
}
