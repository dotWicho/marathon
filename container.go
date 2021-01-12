package marathon

// Container saves in a structured way the information of the containers of a task
type Container struct {
	Type         string        `json:"type"`
	Docker       Docker        `json:"docker"`
	Volumes      []Volume      `json:"volumes,omitempty"`
	PortMappings []PortMapping `json:"portMappings,omitempty"`
}

// Docker holds all information of a docker representation of a task
type Docker struct {
	Image          string             `json:"image"`
	Network        string             `json:"network,omitempty"`
	Privileged     bool               `json:"privileged"`
	Parameters     []DockerParameters `json:"parameters,omitempty"`
	ForcePullImage bool               `json:"forcePullImage"`
}

// Volume is a Docker Volume representation
type Volume struct {
	ContainerPath string `json:"containerPath"`
	HostPath      string `json:"hostPath"`
	Mode          string `json:"mode,omitempty"`
}

// PortMapping is a Docker PortMapping representation
type PortMapping struct {
	ContainerPort int               `json:"containerPort,omitempty"`
	HostPort      int               `json:"hostPort,omitempty"`
	Labels        map[string]string `json:"labels,omitempty"`
	Protocol      string            `json:"protocol,omitempty"`
	ServicePort   int               `json:"servicePort,omitempty"`
}

// DockerParameters Docker exec Parameters representation
type DockerParameters struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

// Queue Tasks representation
type Queue struct {
	Type   string `json:"type"`
	Docker Docker `json:"docker"`
}
