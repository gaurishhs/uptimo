package models

import (
	"time"
  "net/url"

	"github.com/gaurishhs/uptimo/database"
	"github.com/gobuffalo/uuid"
)

type MonitorSettings struct {
  ID        uuid.UUID `db:"id"`
  MonitorID uuid.UUID `db:"monitor_id"`
  Monitor   *Monitor  `belongs_to:"monitor"`
  Type      string    `db:"type"`
  MaxRetries int `db:"max_retries"`
  RetryInterval int `db:"retry_interval"`
  Timeout int `db:"timeout"`
  AuthMethod string `db:"auth_method"`
  BasicAuthUser string `db:"basic_auth_user"`
  BasicAuthPassword string `db:"basic_auth_password"`
  TLSVerify bool `db:"tls_verify"`
  TLSCert string `db:"tls_cert"`
  TLSKey string `db:"tls_key"`
  TLSCaCert string `db:"tls_ca_cert"`
  AuthDomain string `db:"auth_domain"`
  AuthWorkstation string `db:"auth_workstation"`
}

type Monitor struct {
	ID        uuid.UUID  `db:"id"`
	URL       string     `db:"url"`
	CreatedAt time.Time  `db:"created_at"`
	UpdatedAt time.Time  `db:"updated_at"`
	Status    string     `db:"status"`
}

func NewMonitor(tx *database.Connection, monitorURL string) (*Monitor, error) {
	id, err := uuid.NewV4()
	if err != nil {
		return nil, err
	}
  _, err = url.ParseRequestURI(monitorURL)
  if err != nil {
    return nil, err
  }
  // TODO: Validate URL and fetch status
	monitor := &Monitor{
		ID:  id,
		URL: monitorURL,
	}
	if err := tx.Create(monitor); err != nil {
		return nil, err
	}
	return monitor, nil
}
