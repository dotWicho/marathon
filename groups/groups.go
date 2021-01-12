package groups

import (
	"errors"
	"fmt"
	"github.com/dotWicho/marathon"
	"github.com/dotWicho/marathon/application"
	"github.com/dotWicho/marathon/data"
	"github.com/dotWicho/utilities"
	"path/filepath"
	"strings"
	"time"
)

// Marathon Groups interface
type groups interface {
	Get(id string) *Groups
	Create(group *Group) error
	Destroy() error
	Update(group *Group) error

	Scale(instances int, force bool) error
	Stop(force bool) error
	Start(instances int, force bool) error
	Restart(force bool) error
	Suspend(force bool) error

	Load(fileName string) *Groups
	Dump(fileName string) (err error)

	traverseGroupsWithAppID(group *Group, callbackFunc CallBackFuncsWithAppID) (err error)
	traverseGroupsWithAppDefinition(group *Group, callbackFunc CallBackFuncsWithAppDef) (err error)
}

// Groups is Marathon Groups implementation
type Groups struct {
	client *marathon.Client

	//
	group *Group

	//
	deploy *data.Response
	fail   *data.FailureMessage
}

// Group encapsulates the data definitions of a Marathon Group
type Group struct {
	ID           string                      `json:"id"`
	Apps         []application.AppDefinition `json:"apps"`
	Groups       []Group                     `json:"groups"`
	Pods         []interface{}               `json:"pods"`
	Dependencies []string                    `json:"dependencies,omitempty"`
	Version      time.Time                   `json:"version,omitempty"`
	VersionInfo  application.VersionInfo     `json:"versionInfo,omitempty"`
	Executor     string                      `json:"executor,omitempty"`
	EnforceRole  bool                        `json:"enforceRole,omitempty"`
}

// CallBackFuncsWithAppID function type
type CallBackFuncsWithAppID func(appID string) error

// CallBackFuncsWithAppDef function type
type CallBackFuncsWithAppDef func(app application.AppDefinition) error

//=== Marathon Application methods

// NewGroups returns a new instance of Marathon groups implementation
func NewGroups(client *marathon.Client) *Groups {

	if client != nil {
		return &Groups{
			client: client,
			group:  &Group{},
			deploy: &data.Response{},
			fail:   &data.FailureMessage{},
		}
	}
	return nil
}

// Get allows to establish the internal structures to referenced id
func (mg *Groups) Get(id string) *Groups {

	if len(id) > 0 {
		mg.clear()

		path := fmt.Sprintf("%s%s", marathon.APIGroups, utilities.DelInitialSlash(id))

		if _, err := mg.client.Session.BodyAsJSON(nil).Get(path, mg.group, mg.fail); err != nil {
			mg.clear()
		}
	}
	return mg
}

// Create allows create a Marathon group into server
func (mg *Groups) Create(group *Group) error {

	if mg.group != nil && len(group.ID) > 0 {

		path := fmt.Sprintf("%s%s", marathon.APIGroups, utilities.DelInitialSlash(group.ID))

		if _, err := mg.client.Session.BodyAsJSON(group).Post(path, mg.deploy, mg.fail); err != nil {
			return err
		}
		mg.group = group
		return nil
	}
	return errors.New("group cannot be null nor empty")
}

// Destroy erase a Marathon group from server
func (mg *Groups) Destroy() error {

	if mg.group != nil && len(mg.group.ID) > 0 {

		path := fmt.Sprintf("%s%s", marathon.APIGroups, utilities.DelInitialSlash(mg.group.ID))

		if _, err := mg.client.Session.BodyAsJSON(nil).Delete(path, mg.deploy, mg.fail); err != nil {
			return err
		}
		mg.clear()
		return nil
	}
	return errors.New("group cannot be null nor empty")
}

// Update allows change values into Marathon group
func (mg *Groups) Update(group *Group) error {

	if mg.group != nil && len(group.ID) > 0 {

		path := fmt.Sprintf("%s%s", marathon.APIGroups, utilities.DelInitialSlash(group.ID))

		if _, err := mg.client.Session.BodyAsJSON(group).Post(path, mg.deploy, mg.fail); err != nil {
			return err
		}
		mg.group = group
		return nil
	}
	return errors.New("group cannot be null nor empty")
}

// Scale allows change instances numbers of a Marathon group filteredApps
func (mg *Groups) Scale(instances int, force bool) error {

	if mg.group != nil && len(mg.group.ID) > 0 {
		if appClient := application.New(mg.client); appClient != nil {

			callbackFunc := func(appID string) error {

				if err := appClient.Get(appID).Scale(instances, force); err != nil {
					return err
				}
				return nil
			}
			return mg.traverseGroupsWithAppID(mg.group, callbackFunc)
		}
		return fmt.Errorf("unnable to connect")
	}
	return errors.New("group cannot be null nor empty")
}

// Stop sets instances of a Marathon group filteredApps to 0
func (mg *Groups) Stop(force bool) error {

	if mg.group != nil && len(mg.group.ID) > 0 {
		if appClient := application.New(mg.client); appClient != nil {

			callbackFunc := func(appID string) error {

				if err := appClient.Get(appID).Stop(force); err != nil {
					return err
				}
				return nil
			}
			return mg.traverseGroupsWithAppID(mg.group, callbackFunc)
		}
		return fmt.Errorf("unnable to connect")
	}
	return errors.New("group cannot be null nor empty")
}

// Start sets instances of a Marathon group filteredApps to a number provided
func (mg *Groups) Start(instances int, force bool) error {

	if mg.group != nil && len(mg.group.ID) > 0 {
		if appClient := application.New(mg.client); appClient != nil {

			callbackFunc := func(appID string) error {

				if err := appClient.Get(appID).Start(instances, force); err != nil {
					return err
				}
				return nil
			}
			return mg.traverseGroupsWithAppID(mg.group, callbackFunc)
		}
		return fmt.Errorf("unnable to connect")
	}
	return errors.New("group cannot be null nor empty")
}

// Restart use an endpoint to trigger restart for all filteredApps in a Marathon group
func (mg *Groups) Restart(force bool) error {

	if mg.group != nil && len(mg.group.ID) > 0 {
		if appClient := application.New(mg.client); appClient != nil {

			callbackFunc := func(appID string) error {

				if err := appClient.Get(appID).Restart(force); err != nil {
					return err
				}
				return nil
			}
			return mg.traverseGroupsWithAppID(mg.group, callbackFunc)
		}
		return fmt.Errorf("unnable to connect")
	}
	return errors.New("group cannot be null nor empty")
}

// Suspend is an alias for Stop
func (mg *Groups) Suspend(force bool) error {

	return mg.Stop(force)
}

// Apply uses the content of mg.group.Apps to apply the configuration
func (mg *Groups) Apply(force bool) error {

	if mg.group != nil && len(mg.group.ID) > 0 {
		if appClient := application.New(mg.client); appClient != nil {

			callbackFunc := func(app application.AppDefinition) error {

				if err := appClient.Set(app).Apply(force); err != nil {
					return err
				}
				return nil
			}
			return mg.traverseGroupsWithAppDefinition(mg.group, callbackFunc)
		}
		return fmt.Errorf("unnable to connect")
	}
	return errors.New("group cannot be null nor empty")
}

// Load permit read group information from a file
func (mg *Groups) Load(fileName string) *Groups {

	var err error

	mg.clear()

	switch filepath.Ext(strings.TrimSpace(fileName)) {
	case ".json":
		err = utilities.LoadDataFromJSON(mg.group, fileName)
	default:
		err = fmt.Errorf("invalid filename extension")
	}

	if err != nil {
		mg.clear()
	}
	return mg
}

// Dump permit write group information to a file
func (mg *Groups) Dump(fileName string) (err error) {

	if len(mg.group.ID) > 0 {

		switch filepath.Ext(strings.TrimSpace(fileName)) {
		case ".json":
			err = utilities.WriteDataToJSON(mg.group, fileName)
		default:
			err = fmt.Errorf("invalid filename extension")
		}

		return err
	}
	return errors.New("group cannot be null nor empty")
}

// traverseGroupsWithAppID cross the group structure executing a CallBackFuncsWithAppID
func (mg *Groups) traverseGroupsWithAppID(group *Group, callbackFunc CallBackFuncsWithAppID) (err error) {

	marathon.Logger.Debug("traverseGroups: GROUP ID => %s", group.ID)
	if len(group.Apps) > 0 {
		for _, app := range group.Apps {

			err = callbackFunc(app.ID)
		}
	}
	if len(group.Groups) > 0 {
		for _, grp := range group.Groups {
			for _, app := range grp.Apps {
				err = callbackFunc(app.ID)
			}
			err = mg.traverseGroupsWithAppID(&grp, callbackFunc)
		}
	}
	return nil
}

// traverseGroupsWithAppDefinition cross the group structure executing a CallBackFuncsWithAppDef
func (mg *Groups) traverseGroupsWithAppDefinition(group *Group, callbackFunc CallBackFuncsWithAppDef) (err error) {

	marathon.Logger.Debug("traverseGroups: GROUP ID => %s", group.ID)
	if len(group.Apps) > 0 {
		for _, app := range group.Apps {

			err = callbackFunc(app)
		}
	}
	if len(group.Groups) > 0 {
		for _, grp := range group.Groups {
			for _, app := range grp.Apps {
				err = callbackFunc(app)
			}
			err = mg.traverseGroupsWithAppDefinition(&grp, callbackFunc)
		}
	}
	return nil
}

// clean clear internal structures
func (mg *Groups) clear() {

	mg.group = &Group{}
}
