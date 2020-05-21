package marathon

import (
	"errors"
	"fmt"
	"github.com/dotWicho/marathon/pkg/utils"
	"github.com/dotWicho/requist"
	"time"
)

//===

// Marathon Deployments interface
type deployments interface {
	SetClient(client *requist.Requist)
	Get() (*Deployments, error)
	Destroy(id string) error
	Await(id string, timeout time.Duration) error
}

// Marathon Deployments implementation
type Deployments struct {
	client *requist.Requist
	//
	deployments *deployment

	//
	baseUrl string
	auth    string

	//
	deploy *Response
	fail   *FailureMessage
}

//=== Marathon Deployments JSON Entities definition

// Array of Deployment
type deployment []Deployment

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

// SetClient allows reuse of the main object client
func (md *Deployments) SetClient(client *requist.Requist) {

	if client != nil {
		md.client = client
	} else {
		panic(errors.New("client reference cannot be null"))
	}
}

// Get allows to establish the internal structures
func (md *Deployments) Get() (*Deployments, error) {

	if _, err := md.client.BodyAsJSON(nil).Get(marathonApiDeployments, md.deployments, md.fail); err != nil {
		return nil, errors.New("unable to get deployments")
	}
	return md, nil
}

// Destroy cancel a Marathon deployment
func (md *Deployments) Destroy(id string) error {

	if md.deployments != nil {

		path := fmt.Sprintf("%s%s", marathonApiDeployments, utils.DelInitialSlash(id))

		if _, err := md.client.BodyAsJSON(nil).Delete(path, md.deploy, md.fail); err != nil {
			return err
		}
		return nil
	}
	return errors.New("deployment id cannot be null nor empty")
}

// Await wait a Marathon deployment finish or timeout
func (md *Deployments) Await(id string, timeout time.Duration) error {

}
