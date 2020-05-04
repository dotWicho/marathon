package marathon

import "time"


type Deployments []Deployment

type Deployment struct {
	ID             string          `json:"id"`
	Version        time.Time       `json:"version"`
	AffectedApps   []string        `json:"affectedApps"`
	AffectedPods   []string        `json:"affectedPods"`
	Steps          []Step          `json:"steps"`
	CurrentActions []CurrentAction `json:"currentActions"`
	CurrentStep    int             `json:"currentStep"`
	TotalSteps     int             `json:"totalSteps"`
}

type CurrentAction struct {
	Action                string                 `json:"action"`
	App                   string                 `json:"app"`
	ReadinessCheckResults []ReadinessCheckResult `json:"readinessCheckResults"`
}

type ReadinessCheckResult struct {
	TaskID       string       `json:"taskId"`
	LastResponse LastResponse `json:"lastResponse"`
	Name         string       `json:"name"`
	Ready        bool         `json:"ready"`
}

type LastResponse struct {
	Body        string `json:"body"`
	ContentType string `json:"contentType"`
	Status      int    `json:"status"`
}

type Step struct {
	Actions []Action `json:"actions"`
}

type Action struct {
	Action string `json:"action"`
	App    string `json:"app"`
}


type Response struct {
	ID             string          `json:"deploymentId"`
	Version        time.Time       `json:"version"`
}