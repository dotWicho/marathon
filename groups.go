package marathon

import (
	"errors"
	"fmt"
	"github.com/dotWicho/marathon/pkg/utils"
	"github.com/dotWicho/requist"
	"time"
)

type groups interface {
	SetClient(client *requist.Requist)

	Get(id string) (*Groups, error)
	Create(group *Group) error
	Destroy() error
	Update(group *Group) error

	Scale(instances int) bool
	Stop(force bool) bool
	Start() bool
	Restart(force bool) bool
	Suspend(force bool) bool
}

type Groups struct {
	client *requist.Requist

	//
	group *Group

	//
	baseUrl string
	auth    string

	//
	deploy *Response
	fail   *FailureMessage
}

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

func NewMarathonGroups() *Groups {
	mg := &Groups{
		client:  nil,
		group:   &Group{},
		baseUrl: "",
		auth:    "",
		deploy:  &Response{},
		fail:    &FailureMessage{},
	}
	return mg
}

func (mg *Groups) SetClient(client *requist.Requist) {

	if client != nil {
		mg.client = client
	} else {
		panic(errors.New("client reference cannot be null"))
	}
}

func (mg *Groups) Get(id string) (*Groups, error) {

	if len(id) > 0 {
		path := fmt.Sprintf("%s%s", marathonApiGroups, utils.DelInitialSlash(id))

		if _, err := mg.client.BodyAsJSON(nil).Get(path, mg.group, mg.fail); err != nil {
			return nil, errors.New(fmt.Sprintf("unable to get add id = %s", id))
		}
		return mg, nil
	}
	return nil, errors.New("id cannot be empty")
}

func (mg *Groups) Create(group *Group) error {

	if _, err := mg.client.BodyAsJSON(group).Post(marathonApiGroups, mg.deploy, mg.fail); err != nil {
		return err
	}
	mg.group = group
	return nil
}

func (mg *Groups) Destroy() error {

	if mg.group != nil {

		path := fmt.Sprintf("%s%s", marathonApiGroups, utils.DelInitialSlash(mg.group.ID))

		if _, err := mg.client.BodyAsJSON(nil).Delete(path, mg.deploy, mg.fail); err != nil {
			return err
		}
		return nil
	}
	return errors.New("group cannot be null nor empty")
}

func (mg *Groups) Update(group *Group) error {

	if _, err := mg.client.BodyAsJSON(group).Post(marathonApiGroups, mg.deploy, mg.fail); err != nil {
		return err
	}
	return nil
}

func (mg *Groups) Scale(instances int) error {

	return nil
}

func (mg *Groups) Stop(force bool) error {

	return nil
}

func (mg *Groups) Start() error {

	return nil
}

func (mg *Groups) Restart(force bool) error {

	return nil
}

func (mg *Groups) Suspend(force bool) error {

	return nil
}
