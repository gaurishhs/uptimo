package database

import (
	"net/url"

	"github.com/gaurishhs/uptimo/config"
	"github.com/gobuffalo/pop/v5"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

type Connection struct {
	*pop.Connection
}

// Dial connects to the database and returns a connection
func Dial(config *config.Config) (*Connection, error) {
	if config.DB.Driver == "" && config.DB.URL != "" {
		u, err := url.Parse(config.DB.URL)
		if err != nil {
			return nil, errors.Wrap(err, "parsing db connection url")
		}
		config.DB.Driver = u.Scheme
	}

	db, err := pop.NewConnection(&pop.ConnectionDetails{
		Dialect: config.DB.Driver,
		URL:     config.DB.URL,
	})
	if err != nil {
		return nil, errors.Wrap(err, "opening database connection")
	}
	if err := db.Open(); err != nil {
		return nil, errors.Wrap(err, "checking database connection")
	}

	if logrus.StandardLogger().Level == logrus.DebugLevel {
		pop.Debug = true
	}

	return &Connection{db}, nil
}

func (c *Connection) Transaction(fn func(*Connection) error) error {
	if c.TX == nil {
		return c.Connection.Transaction(func(tx *pop.Connection) error {
			return fn(&Connection{tx})
		})
	}
	return fn(c)
}
