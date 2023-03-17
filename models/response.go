package models

import "github.com/gobuffalo/uuid"

type Response struct {
	ID uuid.UUID `db:"id"`
	// Status is the HTTP status code of the response.
	Status int `json:"status"`
	// Latency is the time taken to complete the request.
	Latency   int64     `json:"latency"`
	MonitorID uuid.UUID `json:"-" db:"monitor_id"`
	Monitor   *Monitor  `json:"monitor,omitempty" belongs_to:"monitor"`
}
