package models

import "time"

type Timestamp struct {
	UUID      string    `json:"uuid,omitempty"`
	StartTime time.Time `json:"start_time,omitempty"`
	EndTime   time.Time `json:"end_time,omitempty"`
	Comment   string    `json:"comment"`
	Project   string    `json:"project"`
	Category  string    `json:"category"`
}
