package marathon

import (
	"errors"
	"fmt"
	"time"
)

//===

// Marathon Deployments interface
type deployments interface {
	Get() (*Deployments, error)
	Rollback(id string) error
	Await(id string, timeout time.Duration) error
}

// Deployments is Marathon Deployments implementation
type Deployments struct {
	marathon *Client
	//
	deployments []Deployment

	//
	baseURL string
	auth    string

	//
	deploy *Response
	fail   *FailureMessage
}

//=== Marathon Deployments JSON Entities definition

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

// Step on actions representation
type Step struct {
	Actions []Action `json:"actions"`
}

// Action representation
type Action struct {
	Action string `json:"action"`
	App    string `json:"app"`
}

// Response is default Marathon API response when launch changes via deployments
type Response struct {
	ID      string    `json:"deploymentId"`
	Version time.Time `json:"version"`
}

// NewDeployments returns a new instance of Marathon deployments implementation
func NewDeployments(marathon *Client) *Deployments {

	if marathon != nil {
		_deployment := &Deployments{
			marathon:    marathon,
			deployments: nil,
			baseURL:     marathon.baseURL,
			auth:        marathon.auth,
			deploy:      &Response{},
			fail:        &FailureMessage{},
		}
		return _deployment
	}
	return nil
}

// Get allows to establish the internal structures
func (md *Deployments) Get() (*Deployments, error) {

	if _, err := md.marathon.Session.BodyAsJSON(nil).Get(marathonAPIDeployments, &md.deployments, md.fail); err != nil {
		return md, errors.New("unable to get deployments")
	}
	return md, nil
}

// Rollback cancel a Marathon deployment
func (md *Deployments) Rollback(id string) error {

	if len(id) > 0 {

		path := fmt.Sprintf("%s%s", marathonAPIDeployments, id)

		if _, err := md.marathon.Session.BodyAsJSON(nil).Delete(path, md.deploy, md.fail); err != nil {
			return err
		}
		return nil
	}
	return errors.New("deployment id cannot be null nor empty")
}

// Await wait a Marathon deployment finish or timeout
func (md *Deployments) Await(id string, timeout time.Duration) error {

	// define break condition
	var found bool

	// Start time is Now
	start := time.Now()

	// Finish is Now + timeout
	finish := start.Add(timeout)

	Logger.Debug("Checking for deployment Id = %s", id)

	// iterate while deploy exists or timeout don't reached
	for {
		// Deployment not found by default
		found = false

		if _, err := md.Get(); err != nil {
			break
		}

		for _, deploy := range md.deployments {
			if id == deploy.ID {
				found = true
			}
		}

		if !found || time.Now().After(finish) {
			break
		}
		time.Sleep(1 * time.Second)
	}

	if found {
		return fmt.Errorf("exit by timeout... deployment still existing")
	}
	return nil
}
