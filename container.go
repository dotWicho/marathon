package marathon

type Container struct {
	Type         string        `json:"type"`
	Docker       Docker        `json:"docker"`
	Volumes      []Volume      `json:"volumes"`
	PortMappings []PortMapping `json:"portMappings"`
}

type Docker struct {
	ForcePullImage bool          `json:"forcePullImage"`
	Image          string        `json:"image"`
	Parameters     []interface{} `json:"parameters"`
	Privileged     bool          `json:"privileged"`
}

type Volume struct {
	ContainerPath string `json:"containerPath"`
	HostPath      string `json:"hostPath"`
	Mode          string `json:"mode"`
}

type PortMapping struct {
	ContainerPort int `json:"containerPort"`
	HostPort      int `json:"hostPort"`
	Labels        struct {
		VIP0 string `json:"VIP_0"`
	} `json:"labels"`
	Protocol    string `json:"protocol"`
	ServicePort int    `json:"servicePort"`
}

type Queue struct {
	Type   string      `json:"type"`
	Docker DockerQueue `json:"docker"`
}

type DockerQueue struct {
	Image          string `json:"image"`
	Network        string `json:"network"`
	Privileged     bool   `json:"privileged"`
	ForcePullImage bool   `json:"forcePullImage"`
}
