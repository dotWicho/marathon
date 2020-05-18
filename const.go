package marathon

const (
	marathonApiBase   string = "/v2"
	marathonApiApps   string = marathonApiBase + "/apps/"
	marathonApiGroups string = marathonApiBase + "/groups/"
	marathonApiPing   string = "/ping"

	DockerImageRegEx = `^(([a-zA-Z0-9-]+(?:\.[a-zA-Z0-9-]+)+)(?::(\d+))?)((?:([a-zA-Z0-9-\/]+)?))\/([a-zA-Z0-9-_]+):(.*)$`
)

// FailureMessage all failed request match with this datatype
type FailureMessage struct {
	Message string `json:"message,omitempty"`
}
