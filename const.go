package marathon

import "time"

const (
	// DeploymentTimeout Default timeout for Deployments
	DeploymentTimeout = 10 * time.Second
	// APIBase root of Marathon API Rest
	APIBase string = "/v2"
	// APIApps Applications endpoint
	APIApps string = APIBase + "/apps/"
	// APIGroups Groups endpoint
	APIGroups string = APIBase + "/groups/"
	// APIDeployments Deployments endpoint
	APIDeployments string = APIBase + "/deployments/"
	// APIPing Check connection endpoint resource
	APIPing string = "/ping"
	// APIInfo Server Info endpoint
	APIInfo string = APIBase + "/info"
	// DockerImageRegEx RegEx used for docker images
	DockerImageRegEx = `^(([a-zA-Z0-9-]+(?:\.[a-zA-Z0-9-]+)+)(?::(\d+))?)((?:([a-zA-Z0-9-\/]+)?))\/([a-zA-Z0-9-_]+):(.*)$`
)

// FailureMessage all failed request match with this datatype
type FailureMessage struct {
	Message string `json:"message,omitempty"`
}
