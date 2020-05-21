package marathon

const (
	// root of Marathon API Rest
	marathonApiBase string = "/v2"
	// Applications endpoint
	marathonApiApps string = marathonApiBase + "/apps/"
	// Groups endpoint
	marathonApiGroups string = marathonApiBase + "/groups/"
	// Deployments endpoint
	marathonApiDeployments string = marathonApiBase + "/deployments/"
	// Check connection endpoint resource
	marathonApiPing string = "/ping"
	// RegEx used for docker images
	DockerImageRegEx = `^(([a-zA-Z0-9-]+(?:\.[a-zA-Z0-9-]+)+)(?::(\d+))?)((?:([a-zA-Z0-9-\/]+)?))\/([a-zA-Z0-9-_]+):(.*)$`
)

// FailureMessage all failed request match with this datatype
type FailureMessage struct {
	Message string `json:"message,omitempty"`
}
