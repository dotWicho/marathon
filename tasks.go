package marathon

import (
	"time"
)

// Task is an Envelope structure for Tasks response from Marathon server
type Task struct {
	Task TaskMarathon `json:"task"`
}

// Tasks Array of Task
type Tasks struct {
	Tasks []TaskMarathon `json:"tasks"`
}

// TaskMarathon is Marathon Task representation
type TaskMarathon struct {
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

// IPAddress IP and protocol information of a Task
type IPAddress struct {
	IPAddress string `json:"ipAddress"`
	Protocol  string `json:"protocol"`
}
