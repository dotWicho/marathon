package marathon

import "time"

// Array of Deployment
type Deployments []Deployment

// Deployment holds Marathons deploys on course
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

// CurrentAction holds actions taken by a deployment
type CurrentAction struct {
	Action                string                 `json:"action"`
	App                   string                 `json:"app"`
	ReadinessCheckResults []ReadinessCheckResult `json:"readinessCheckResults"`
}

// ReadinessCheckResult holds results for tasks
type ReadinessCheckResult struct {
	TaskID       string       `json:"taskId"`
	LastResponse LastResponse `json:"lastResponse"`
	Name         string       `json:"name"`
	Ready        bool         `json:"ready"`
}

// LastResponse holds last response
type LastResponse struct {
	Body        string `json:"body"`
	ContentType string `json:"contentType"`
	Status      int    `json:"status"`
}

// That, a Step representation
type Step struct {
	Actions []Action `json:"actions"`
}

// That, a action representation
type Action struct {
	Action string `json:"action"`
	App    string `json:"app"`
}

// Marathon API response when launch changes via deployments
type Response struct {
	ID      string    `json:"deploymentId"`
	Version time.Time `json:"version"`
}
