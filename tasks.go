package marathon

import (
	"time"
)

type Task struct {
	Task Tsk `json:"task"`
}

type Tasks struct {
	Tasks []Tsk `json:"tasks"`
}

type Tsk struct {
	AppID              string        `json:"appId"`
	HealthCheckResults []Result      `json:"healthCheckResults"`
	Host               string        `json:"host"`
	ID                 string        `json:"id"`
	IPAddresses        []IPAddress   `json:"ipAddresses"`
	Ports              []int         `json:"ports"`
	ServicePorts       []interface{} `json:"servicePorts"`
	SlaveID            string        `json:"slaveId"`
	State              string        `json:"state"`
	StagedAt           time.Time     `json:"stagedAt"`
	StartedAt          time.Time     `json:"startedAt"`
	Version            time.Time     `json:"version"`
	LocalVolumes       []interface{} `json:"localVolumes"`
	Role               string        `json:"role"`
}

type IPAddress struct {
	IPAddress string `json:"ipAddress"`
	Protocol  string `json:"protocol"`
}
