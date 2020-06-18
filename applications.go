package marathon

import (
	"errors"
	"fmt"
	"github.com/dotWicho/marathon/pkg/utils"
	"strings"
	"time"
)

// Apps wraps an AppDefinition array returned by the Marathon API
type apps struct {
	Apps []AppDefinition `json:"apps"`
}

//
type AppSummary struct {
	ID     string            `json:"id,omitempty"`
	Cpus   float64           `json:"cpus,omitempty"`
	Mem    float64           `json:"mem,omitempty"`
	Env    map[string]string `json:"env,omitempty"`
	Labels map[string]string `json:"labels,omitempty"`
	Image  string            `json:"image,omitempty"`
}

// Marathon Application interface
type applications interface {
	Get(filter string) (*Applications, error)
	Scale(instances int, force bool) error
	Stop(force bool) error
	Start(instances int, force bool) error
	Restart(force bool) error
	Suspend(force bool) error

	LoadFromFile(fileName, fileType string) error
	DumpToFile(fileName, fileType string) error

	Apply() error

	AsMap() map[string]AppSummary
	AsRaw() []AppDefinition
}

// Marathon Application implementation
type Applications struct {
	marathon *Client
	timeout  time.Duration
	//
	apps *apps

	//
	baseUrl string
	auth    string

	//
	deploy *Response
	fail   *FailureMessage
}

// NewApplications returns a new instance of Marathon applications implementation
func NewApplications(marathon *Client) *Applications {
	_applications := &Applications{
		marathon: marathon,
		timeout:  marathon.timeout,
		apps:     &apps{},
		baseUrl:  marathon.baseUrl,
		auth:     marathon.auth,
		deploy:   &Response{},
		fail:     &FailureMessage{},
	}
	return _applications
}

// Get allows to establish the internal structures to referenced id
func (ma *Applications) Get(filter string) (*Applications, error) {

	if len(filter) > 0 {

		_apps := &apps{}

		if _, err := ma.marathon.Session.BodyAsJSON(nil).Get(marathonApiApps, _apps, ma.fail); err != nil {
			return nil, errors.New(fmt.Sprintf("unable to get apps with filter = %s", filter))
		}

		for _, app := range _apps.Apps {
			if strings.HasPrefix(app.ID, filter) {
				ma.apps.Apps = append(ma.apps.Apps, app)
			}
		}

		return ma, nil
	}
	return nil, errors.New("filter cannot be null nor empty")
}

// Scale allows change instances numbers of a Marathon applications
func (ma *Applications) Scale(instances int, force bool) error {

	if ma.apps != nil && len(ma.apps.Apps) > 0 {
		appHandler := NewApplication(ma.marathon)
		for _, app := range ma.apps.Apps {
			if _, err := appHandler.Get(app.ID); err == nil {
				_ = appHandler.Scale(instances, force)
			}
		}
	}
	return nil
}

// Stop sets instances of a Marathon applications to 0
func (ma *Applications) Stop(force bool) error {

	if ma.apps != nil && len(ma.apps.Apps) > 0 {
		appHandler := NewApplication(ma.marathon)
		for _, app := range ma.apps.Apps {
			if _, err := appHandler.Get(app.ID); err == nil {
				_ = appHandler.Stop(force)
			}
		}
	}
	return nil
}

// Start sets instances of a Marathon applications to a number provided
func (ma *Applications) Start(instances int, force bool) error {

	if ma.apps != nil && len(ma.apps.Apps) > 0 {
		appHandler := NewApplication(ma.marathon)
		for _, app := range ma.apps.Apps {
			if _, err := appHandler.Get(app.ID); err == nil {
				_ = appHandler.Start(instances, force)
			}
		}
	}
	return nil
}

// Restart use an endpoint to trigger a Marathon applications restart
func (ma *Applications) Restart(force bool) error {

	if ma.apps != nil && len(ma.apps.Apps) > 0 {
		appHandler := NewApplication(ma.marathon)
		for _, app := range ma.apps.Apps {
			if _, err := appHandler.Get(app.ID); err == nil {
				_ = appHandler.Restart(force)
			}
		}
	}
	return nil
}

// Suspend is an alias to Stop
func (ma *Applications) Suspend(force bool) error {

	return ma.Stop(force)
}

// LoadFromFile allows create or update a Marathon applications from file
func (ma *Applications) LoadFromFile(fileName, fileType string) error {

	var err error

	switch fileType {
	case "json":
		err = utils.LoadDataFromJson(ma.apps.Apps, fileName)
	case "yaml":
		err = utils.LoadDataFromYaml(ma.apps.Apps, fileName)
	}

	return err
}

// DumpToFile allows to create a file with the configuration of applications
func (ma *Applications) DumpToFile(fileName, fileType string) error {

	var err error

	switch fileType {
	case "json":
		err = utils.WriteDataToJson(ma.apps.Apps, fileName)
	case "yaml":
		err = utils.WriteDataToYaml(ma.apps.Apps, fileName)
	}

	return err
}

// Apply make a patch to each app
func (ma *Applications) Apply() error {

	if ma.apps != nil {
		if len(ma.apps.Apps) > 0 {
			for _, app := range ma.apps.Apps {
				path := fmt.Sprintf("%s%s", marathonApiApps, utils.DelInitialSlash(app.ID))

				ma.marathon.Session.AddQueryParam("force", "true")

				if _, err := ma.marathon.Session.BodyAsJSON(app).Patch(path, ma.deploy, ma.fail); err != nil {
					return err
				}
				// TODO: Deployment wait for ma.timeout
				fmt.Printf("Deploy Id: %s => date: %v\n", ma.deploy.ID, ma.deploy.Version)
			}
			return nil
		}
	}
	return errors.New("apps cannot be null nor empty")
}

// AsMap returns a map of Summary Info
func (ma *Applications) AsMap() map[string]AppSummary {

	mapApps := make(map[string]AppSummary)
	for _, app := range ma.apps.Apps {
		mapApps[app.ID] = AppSummary{
			ID:     app.ID,
			Cpus:   app.Cpus,
			Mem:    app.Mem,
			Env:    app.Env,
			Labels: app.Labels,
			Image:  app.Container.Docker.Image,
		}
	}
	return mapApps
}

// AsRaw returns a pointer of Application Info
func (ma *Applications) AsRaw() []AppDefinition {

	return ma.apps.Apps
}
