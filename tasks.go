package marathon

import (
	"time"
)

// Envelope structure for Taks response from Marathon server
type Task struct {
	Task Tsk `json:"task"`
}

// Array of
type Tasks struct {
	Tasks []Tsk `json:"tasks"`
}

// Tsk, Marathon Task representation
type Tsk struct {
	AppID              string              `json:"appId"`
	HealthCheckResults []HealthcheckResult `json:"healthCheckResults"`
	Host               string              `json:"host"`
	ID                 string              `json:"id"`
	IPAddresses        []IPAddress         `json:"ipAddresses"`
	Ports              []int               `json:"ports"`
	ServicePorts       []interface{}       `json:"servicePorts"`
	SlaveID            string              `json:"slaveId"`
	State              string              `json:"state"`
	StagedAt           time.Time           `json:"stagedAt"`
	StartedAt          time.Time           `json:"startedAt"`
	Version            time.Time           `json:"version"`
	LocalVolumes       []interface{}       `json:"localVolumes"`
	Role               string              `json:"role"`
}

// IP and protocol information of a Taks
type IPAddress struct {
	IPAddress string `json:"ipAddress"`
	Protocol  string `json:"protocol"`
}
