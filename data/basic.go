package data

import "time"

// FailureMessage all failed request match with this datatype
type FailureMessage struct {
	Message string `json:"message,omitempty"`
}

// Response is default Marathon API response when launch changes via deployments
type Response struct {
	ID      string    `json:"deploymentId"`
	Version time.Time `json:"version"`
}
