package marathon

import "time"

const (
	// Default timeout for Deployments
	defaultDeploymentTimeout = 1 * time.Second
	// marathonAPIBase is root of Marathon API Rest
	marathonAPIBase string = "/v2"
	// marathonAPIApps is FilteredApps endpoint
	marathonAPIApps string = marathonAPIBase + "/apps/"
	// marathonAPIGroups is Groups endpoint
	marathonAPIGroups string = marathonAPIBase + "/groups/"
	// marathonAPIDeployments is Deployments endpoint
	marathonAPIDeployments string = marathonAPIBase + "/deployments/"
	// marathonAPIPing to Check connection endpoint resource
	marathonAPIPing string = "/ping"
	// marathonAPIInfo is Server Info endpoint
	marathonAPIInfo string = marathonAPIBase + "/info"
	// DockerImageRegEx is a RegEx used for docker images
	DockerImageRegEx = `^(([a-zA-Z0-9-]+(?:\.[a-zA-Z0-9-]+)+)(?::(\d+))?)((?:([a-zA-Z0-9-\/]+)?))\/([a-zA-Z0-9-_]+):(.*)$`
)

// FailureMessage all failed request match with this datatype
type FailureMessage struct {
	Message string      `json:"message,omitempty"`
	Details interface{} `json:"details,omitempty"`
}
