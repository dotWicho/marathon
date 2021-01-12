package filtered

import (
	"fmt"
	"github.com/dotWicho/marathon"
	"github.com/dotWicho/marathon/application"
	"github.com/dotWicho/marathon/data"
	"github.com/dotWicho/utilities"
	"path/filepath"
	"strings"
)

// FilterFunction is a type to create callback functions
type FilterFunction func(app application.AppDefinition) bool

// Apps wraps an AppDefinition array returned by the Marathon API
type apps struct {
	Apps []application.AppDefinition `json:"apps"`
}

// AppSummary holds a resumed format of Application
type AppSummary struct {
	ID     string            `json:"id,omitempty"`
	Cpus   float64           `json:"cpus,omitempty"`
	Mem    float64           `json:"mem,omitempty"`
	Env    map[string]string `json:"env,omitempty"`
	Labels map[string]string `json:"labels,omitempty"`
	Image  string            `json:"image,omitempty"`
}

// filteredApps Marathon Application FilteredApps interface
type filteredApps interface {
	Get(filter string) *Apps
	Scale(instances int, force bool) error
	Stop(force bool) error
	Start(instances int, force bool) error
	Restart(force bool) error
	Suspend(force bool) error

	Load(fileName, filter string) *Apps
	Dump(fileName string) (err error)
	DumpSingly(baseName string) (err error)

	FilterBy(filterFunc FilterFunction) *Apps

	AsMap() map[string]AppSummary
	AsRaw() []application.AppDefinition
}

// FilteredApps is a Marathon Applications by filter implementation
type Apps struct {
	client  *marathon.Client
	//
	apps *apps

	//
	deploy *data.Response
	fail   *data.FailureMessage
}

// NewFilteredApps returns a new instance of Marathon filteredApps implementation
func NewFilteredApps(client *marathon.Client) *Apps {

	if client != nil {
		return &Apps{
			client: client,
			apps:   &apps{},
			deploy: &data.Response{},
			fail:   &data.FailureMessage{},
		}
	}
	return nil
}

// Get allows to establish the internal structures to referenced id
func (fa *Apps) Get(filter string) *Apps {

	if len(filter) > 0 {

		marathon.Logger.Debug("FilteredApps: Get (%s)", filter)
		_apps := &apps{}

		if _, err := fa.client.Session.BodyAsJSON(nil).Get(marathon.APIApps, _apps, fa.fail); err != nil {
			fa.apps.Apps = nil
			return fa
		}

		for _, app := range _apps.Apps {
			if strings.HasPrefix(app.ID, filter) {
				marathon.Logger.Debug("FilteredApps: Get %s match", app.ID)
				fa.apps.Apps = append(fa.apps.Apps, app)
			}
		}

		marathon.Logger.Debug("FilteredApps: Get (%s) found %d apps", filter, len(fa.apps.Apps))
	}
	return fa
}

// Scale allows change instances numbers of a Marathon filteredApps
func (fa *Apps) Scale(instances int, force bool) error {

	if fa.apps != nil && len(fa.apps.Apps) > 0 {

		marathon.Logger.Debug("FilteredApps: Scale %d %v [%+v] %d", instances, force, fa.apps.Apps, len(fa.apps.Apps))

		appHandler := application.New(fa.client)
		for index, app := range fa.apps.Apps {

			if err := appHandler.Get(app.ID).Scale(instances, force); err != nil {
				return err
			}
			fa.apps.Apps[index].Instances = instances
		}
		return nil
	}
	return fmt.Errorf("filteredApps Scale was called with an empty set")
}

// Stop sets instances of a Marathon filteredApps to 0
func (fa *Apps) Stop(force bool) error {

	if fa.apps != nil && len(fa.apps.Apps) > 0 {

		marathon.Logger.Debug("FilteredApps: Stop %v [%+v] %d", force, fa.apps.Apps, len(fa.apps.Apps))

		appHandler := application.New(fa.client)
		for index, app := range fa.apps.Apps {

			if err := appHandler.Get(app.ID).Stop(force); err != nil {
				return err
			}
			fa.apps.Apps[index].Instances = 0
		}
		return nil
	}
	return fmt.Errorf("filteredApps Stop was called with an empty set")
}

// Start sets instances of a Marathon filteredApps to a number provided
func (fa *Apps) Start(instances int, force bool) error {

	if fa.apps != nil && len(fa.apps.Apps) > 0 {

		marathon.Logger.Debug("FilteredApps: Start %v [%+v] %d", force, fa.apps.Apps, len(fa.apps.Apps))

		appHandler := application.New(fa.client)
		for index, app := range fa.apps.Apps {
			if err := appHandler.Get(app.ID).Start(instances, force); err != nil {
				return err
			}
			fa.apps.Apps[index].Instances = instances
		}
		return nil
	}
	return fmt.Errorf("filteredApps Start was called with an empty set")
}

// Restart use an endpoint to trigger a Marathon filteredApps restart
func (fa *Apps) Restart(force bool) error {

	if fa.apps != nil && len(fa.apps.Apps) > 0 {

		marathon.Logger.Debug("FilteredApps: Restart %v [%+v] %d", force, fa.apps.Apps, len(fa.apps.Apps))

		appHandler := application.New(fa.client)
		for _, app := range fa.apps.Apps {
			if err := appHandler.Get(app.ID).Restart(force); err != nil {
				return err
			}
		}
		return nil
	}
	return fmt.Errorf("filteredApps Restart was called with an empty set")
}

// Suspend is an alias to Stop
func (fa *Apps) Suspend(force bool) error {

	return fa.Stop(force)
}

// Load allows create or update a Marathon filteredApps from file
func (fa *Apps) Load(fileName, filter string) *Apps {

	var err error
	_apps := &apps{}

	switch filepath.Ext(strings.TrimSpace(fileName)) {
	case ".json":
		err = utilities.LoadDataFromJSON(_apps, fileName)
	default:
		err = fmt.Errorf("invalid filename extension")
	}

	if err == nil {
		for _, app := range _apps.Apps {
			if strings.HasPrefix(app.ID, filter) {
				marathon.Logger.Debug("FilteredApps: Load %s match", app.ID)
				fa.apps.Apps = append(fa.apps.Apps, app)
			}
		}
		marathon.Logger.Debug("FilteredApps: Load (%s) found %d apps", filter, len(fa.apps.Apps))
	} else {
		fa.apps.Apps = nil
		fa.apps.Apps = []application.AppDefinition{}
	}

	return fa
}

// Dump allows to create a file with the configuration of filteredApps
func (fa *Apps) Dump(fileName string) (err error) {

	if fa.apps != nil && len(fa.apps.Apps) > 0 {

		switch filepath.Ext(strings.TrimSpace(fileName)) {
		case ".json":
			err = utilities.WriteDataToJSON(fa.apps.Apps, fileName)
		default:
			err = fmt.Errorf("invalid filename extension")
		}

		return
	}
	return fmt.Errorf("filteredApps Dump was called with an empty set")
}

// Dump allows to create a file with the configuration of filteredApps
func (fa *Apps) DumpSingly(baseName string) (err error) {

	if fa.apps != nil && len(fa.apps.Apps) > 0 {

		switch filepath.Ext(strings.TrimSpace(baseName)) {
		case ".json":

			baseName = strings.TrimSuffix(baseName, filepath.Ext(baseName))

			for _, app := range fa.apps.Apps {

				appFileName := baseName + utilities.BaseName(app.ID) + ".json"
				err = utilities.WriteDataToJSON(app, appFileName)
				if err != nil {
					marathon.Logger.Error("Writing app %s [%+v]", app.ID, err)
				}
			}

		default:
			err = fmt.Errorf("invalid filename extension")
		}

		return
	}
	return fmt.Errorf("filteredApps Dump was called with an empty set")
}

// FilterBy make a new apps.Apps just with those match filterFunc
func (fa *Apps) FilterBy(filterFunc FilterFunction) *Apps {

	if fa.apps != nil && len(fa.apps.Apps) > 0 {
		var filtered []application.AppDefinition

		for _, app := range fa.apps.Apps {
			if filterFunc(app) {
				filtered = append(filtered, app)
			}
		}
		fa.apps.Apps = filtered
	}
	return fa
}

// AsMap returns a map of Summary Info
func (fa *Apps) AsMap() map[string]AppSummary {

	if fa.apps != nil && len(fa.apps.Apps) > 0 {

		mapApps := make(map[string]AppSummary)
		for _, app := range fa.apps.Apps {
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
	return nil
}

// AsRaw returns a pointer of Application Info
func (fa *Apps) AsRaw() []application.AppDefinition {

	if fa.apps != nil && len(fa.apps.Apps) > 0 {
		return fa.apps.Apps
	}
	return nil
}
