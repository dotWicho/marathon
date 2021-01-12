package deployments

import (
	"errors"
	"fmt"
	"github.com/dotWicho/marathon"
	"github.com/dotWicho/requist"
	"github.com/dotWicho/utilities"
	"time"
)

//===

// Marathon Deployments interface
type deployments interface {
	Get() map[string]Deployment
	Exist(id string) bool
	Rollback(id string) error
	Await(id string, timeout time.Duration) error
}

// Deployments is a Marathon Deployments implementation
type Deployments struct {
	client *requist.Requist

	//
	deploy Deployment
	fail   *marathon.FailureMessage
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

// Step just that, a Step representation
type Step struct {
	Actions []Action `json:"actions"`
}

// Action just that, an action representation
type Action struct {
	Action string `json:"action"`
	App    string `json:"app"`
}

// Response is Marathon API response when launch changes via deployments
type Response struct {
	ID      string    `json:"deploymentId"`
	Version time.Time `json:"version"`
}

// New returns a new instance of Marathon deployment
func New() *Deployments {

	return &Deployments{
		client:  nil,
		deploy:  Deployment{},
		fail:    &marathon.FailureMessage{},
	}
}

// SetClient allows reuse of the main object client
func (md *Deployments) SetClient(client *requist.Requist) error {

	if client == nil {
		return errors.New("client reference cannot be null")
	}
	md.client = client
	return nil
}

// Get allows to establish the internal structures
func (md *Deployments) Get() map[string]Deployment {

	mapDeploy := make(map[string]Deployment)
	var deploys []Deployment
	if _, err := md.client.BodyAsJSON(nil).Get(marathon.APIDeployments, &deploys, md.fail); err != nil {
		return nil
	}

	if len(deploys) > 0 && len(md.fail.Message) == 0 {
		for _, deploy := range deploys {
			mapDeploy[deploy.ID] = Deployment{
				ID:             deploy.ID,
				Version:        deploy.Version,
				AffectedApps:   deploy.AffectedApps,
				AffectedPods:   deploy.AffectedPods,
				Steps:          deploy.Steps,
				CurrentActions: deploy.CurrentActions,
				CurrentStep:    deploy.CurrentStep,
				TotalSteps:     deploy.TotalSteps,
			}
		}
	}
	return mapDeploy
}

// Get allows to establish the internal structures
func (md *Deployments) Exist(id string) bool {

	return len(md.Get()[id].ID) > 0
}

// Rollback cancel a Marathon deployment
func (md *Deployments) Rollback(id string) error {

	if len(id) > 0 && md.Exist(id) {

		path := fmt.Sprintf("%s%s", marathon.APIDeployments, utilities.DelInitialSlash(id))

		if _, err := md.client.BodyAsJSON(nil).Delete(path, md.deploy, md.fail); err != nil {
			return err
		}
		return nil
	}
	return errors.New("deployment id cannot be null nor empty")
}

// Await wait a Marathon deployment finish or timeout
func (md *Deployments) Await(id string, timeout time.Duration) (err error) {

	stop := make(chan bool)

	go func(end time.Time) {
		for {
			fmt.Println(".")
			if !md.Exist(id) {
				err = nil
				stop <- true
			}
			if time.Now().After(end) {
				err = errors.New("deploy still executing. Timeout reached")
				stop <- true
			}
			select {
			case <-time.After(1 * time.Second):
			case <-stop:
				return
			}
		}
	}(time.Now().Add(timeout))

	return
}
