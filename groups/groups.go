package groups

import (
	"errors"
	"fmt"
	"github.com/dotWicho/marathon"
	"github.com/dotWicho/marathon/application"
	"github.com/dotWicho/marathon/deployments"
	"github.com/dotWicho/requist"
	"github.com/dotWicho/utilities"
	"time"
)

// Marathon Groups interface
type groups interface {
	Get(id string) (*Groups, error)
	Create(group *Groups) error
	Destroy() error
	Update(group *Group) error

	Scale(instances int) error
	Stop(force bool) error
	Start(instances int, force bool) error
	Restart(force bool) error
	Suspend(force bool) error
}

// Groups is a Marathon Groups implementation
type Groups struct {
	client *requist.Requist

	//
	group *Group

	//
	baseURL string
	auth    string

	//
	deploy *deployments.Response
	fail   *marathon.FailureMessage
}

// Group encapsulates the data definitions of a Marathon Group
type Group struct {
	ID           string                  `json:"id"`
	Apps         []application.App       `json:"apps"`
	Groups       []Group                 `json:"groups"`
	Pods         []interface{}           `json:"pods"`
	Dependencies []string                `json:"dependencies,omitempty"`
	Version      time.Time               `json:"version,omitempty"`
	VersionInfo  application.VersionInfo `json:"versionInfo,omitempty"`
	Executor     string                  `json:"executor,omitempty"`
	EnforceRole  bool                    `json:"enforceRole,omitempty"`
}

//=== Marathon Application methods

func NewMarathonGroups() *Groups {
	mg := &Groups{
		client:  nil,
		group:   &Group{},
		baseURL: "",
		auth:    "",
		deploy:  &deployments.Response{},
		fail:    &marathon.FailureMessage{},
	}
	return mg
}

// SetClient allows reuse of the main object client
func (mg *Groups) SetClient(client *requist.Requist) error {

	if client == nil {
		return errors.New("client reference cannot be null")
	}
	mg.client = client
	return nil
}

// Get allows to establish the internal structures to referenced id
func (mg *Groups) Get(id string) (*Groups, error) {

	if len(id) > 0 {
		path := fmt.Sprintf("%s%s", marathon.APIGroups, utilities.DelInitialSlash(id))

		if _, err := mg.client.BodyAsJSON(nil).Get(path, mg.group, mg.fail); err != nil {
			return nil, errors.New(fmt.Sprintf("unable to get add id = %s", id))
		}
		return mg, nil
	}
	return nil, errors.New("id cannot be empty")
}

// Create allows create a Marathon group into server
func (mg *Groups) Create(group *Group) error {

	if _, err := mg.client.BodyAsJSON(group).Post(marathon.APIGroups, mg.deploy, mg.fail); err != nil {
		return err
	}
	mg.group = group
	return nil
}

// Destroy erase a Marathon group from server
func (mg *Groups) Destroy() error {

	if mg.group != nil {

		path := fmt.Sprintf("%s%s", marathon.APIGroups, utilities.DelInitialSlash(mg.group.ID))

		if _, err := mg.client.BodyAsJSON(nil).Delete(path, mg.deploy, mg.fail); err != nil {
			return err
		}
		return nil
	}
	return errors.New("group cannot be null nor empty")
}

// Update allows change values into Marathon group
func (mg *Groups) Update(group *Group) error {

	if _, err := mg.client.BodyAsJSON(group).Post(marathon.APIGroups, mg.deploy, mg.fail); err != nil {
		return err
	}
	return nil
}

// Scale allows change instances numbers of a Marathon group applications
func (mg *Groups) Scale(instances int) error {

	return nil
}

// Stop sets instances of a Marathon group applications to 0
func (mg *Groups) Stop(force bool) error {

	return nil
}

// Start sets instances of a Marathon group applications to a number provided
func (mg *Groups) Start(instances int, force bool) error {

	return nil
}

// Restart use an endpoint to trigger restart for all applications in a Marathon group
func (mg *Groups) Restart(force bool) error {

	return nil
}

// Suspend is an alias for Stop
func (mg *Groups) Suspend(force bool) error {

	return nil
}
