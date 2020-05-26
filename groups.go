package marathon

import (
	"errors"
	"fmt"
	"github.com/dotWicho/marathon/pkg/utils"
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

// Marathon Groups implementation
type Groups struct {
	marathon *Client

	//
	group *Group

	//
	baseUrl string
	auth    string

	//
	deploy *Response
	fail   *FailureMessage
}

// Group encapsulates the data definitions of a Marathon Group
type Group struct {
	ID           string        `json:"id"`
	Apps         []App         `json:"apps"`
	Groups       []Group       `json:"groups"`
	Pods         []interface{} `json:"pods"`
	Dependencies []string      `json:"dependencies,omitempty"`
	Version      time.Time     `json:"version,omitempty"`
	VersionInfo  VersionInfo   `json:"versionInfo,omitempty"`
	Executor     string        `json:"executor,omitempty"`
	EnforceRole  bool          `json:"enforceRole,omitempty"`
}

//=== Marathon Application methods

// NewGroups returns a new instance of Marathon groups implementation
func NewGroups(marathon *Client) *Groups {
	_groups := &Groups{
		marathon: marathon,
		group:    &Group{},
		baseUrl:  marathon.baseUrl,
		auth:     marathon.auth,
		deploy:   &Response{},
		fail:     &FailureMessage{},
	}
	return _groups
}

// Get allows to establish the internal structures to referenced id
func (mg *Groups) Get(id string) (*Groups, error) {

	if len(id) > 0 {
		path := fmt.Sprintf("%s%s", marathonApiGroups, utils.DelInitialSlash(id))

		if _, err := mg.marathon.Session.BodyAsJSON(nil).Get(path, mg.group, mg.fail); err != nil {
			return nil, errors.New(fmt.Sprintf("unable to get group id = %s", id))
		}
		return mg, nil
	}
	return nil, errors.New("id cannot be null nor empty")
}

// Create allows create a Marathon group into server
func (mg *Groups) Create(group *Group) error {

	if _, err := mg.marathon.Session.BodyAsJSON(group).Post(marathonApiGroups, mg.deploy, mg.fail); err != nil {
		return err
	}
	mg.group = group
	return nil
}

// Destroy erase a Marathon group from server
func (mg *Groups) Destroy() error {

	if mg.group != nil {

		path := fmt.Sprintf("%s%s", marathonApiGroups, utils.DelInitialSlash(mg.group.ID))

		if _, err := mg.marathon.Session.BodyAsJSON(nil).Delete(path, mg.deploy, mg.fail); err != nil {
			return err
		}
		return nil
	}
	return errors.New("group cannot be null nor empty")
}

// Update allows change values into Marathon group
func (mg *Groups) Update(group *Group) error {

	if _, err := mg.marathon.Session.BodyAsJSON(group).Post(marathonApiGroups, mg.deploy, mg.fail); err != nil {
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
