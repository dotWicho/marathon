package marathon

type Container struct {
	Type         string        `json:"type"`
	Docker       Docker        `json:"docker"`
	Volumes      []Volume      `json:"volumes,omitempty"`
	PortMappings []PortMapping `json:"portMappings,omitempty"`
}

type Docker struct {
	Image          string             `json:"image"`
	Network        string             `json:"network,omitempty"`
	Privileged     bool               `json:"privileged"`
	Parameters     []DockerParameters `json:"parameters,omitempty"`
	ForcePullImage bool               `json:"forcePullImage"`
}

type Volume struct {
	ContainerPath string `json:"containerPath"`
	HostPath      string `json:"hostPath"`
	Mode          string `json:"mode,omitempty"`
}

type PortMapping struct {
	ContainerPort int               `json:"containerPort,omitempty"`
	HostPort      int               `json:"hostPort,omitempty"`
	Labels        map[string]string `json:"labels,omitempty"`
	Protocol      string            `json:"protocol,omitempty"`
	ServicePort   int               `json:"servicePort,omitempty"`
}

type DockerParameters struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type Queue struct {
	Type   string `json:"type"`
	Docker Docker `json:"docker"`
}
