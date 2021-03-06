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
	// APIVersions Apps Configs versions endpoint
	APIVersions string = APIApps + "%s/versions"
	// APIConfigByVersion Apps Definition by version endpoint
	APIConfigByVersion = APIVersions + "/%s"
	// DockerImageRegEx RegEx used for docker images
	DockerImageRegEx = `^(([a-zA-Z0-9-]+(?:\.[a-zA-Z0-9-]+)+)(?::(\d+))?)((?:([a-zA-Z0-9-\/]+)?))\/([a-zA-Z0-9-_]+):(.*)$`
)
