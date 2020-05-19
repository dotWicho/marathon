package marathon

import "time"

// Healthcheck represents configuration of a health check of a Marathon task
type Healthcheck struct {
	GracePeriodSeconds     int    `json:"gracePeriodSeconds"`
	IntervalSeconds        int    `json:"intervalSeconds"`
	MaxConsecutiveFailures int    `json:"maxConsecutiveFailures"`
	Path                   string `json:"path"`
	PortIndex              int    `json:"portIndex"`
	Protocol               string `json:"protocol"`
	IPProtocol             string `json:"ipProtocol"`
	TimeoutSeconds         int    `json:"timeoutSeconds"`
	DelaySeconds           int    `json:"delaySeconds"`
}

// HealthcheckResult represents response of a Marathon health check
type HealthcheckResult struct {
	Alive               bool      `json:"alive"`
	ConsecutiveFailures int       `json:"consecutiveFailures"`
	FirstSuccess        time.Time `json:"firstSuccess"`
	InstanceID          string    `json:"instanceId"`
	LastSuccess         time.Time `json:"lastSuccess"`
}
