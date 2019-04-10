package logs

import "time"

type (
	Event struct {
		Timestamp time.Time `json:"timestamp"`
		Level     string    `json:"level"`
		Message   string    `json:"message"`
	}
)
