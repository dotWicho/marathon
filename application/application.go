package application

import (
	"errors"
	"fmt"
	"github.com/dotWicho/marathon"
	"github.com/dotWicho/marathon/data"
	"github.com/dotWicho/utilities"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
	"time"
)

//===

// App wraps an AppDefinition element returned by the Marathon API
type App struct {
	App AppDefinition `json:"app"`
}

//===

// Marathon Application interface
type application interface {
	Get(id string) *Application
	Set(app AppDefinition) *Application
	Create(app AppDefinition) *Application
	Destroy() error
	Update(app AppDefinition) error

	Instances() int

	Scale(instances int, force bool) error
	Stop(force bool) error
	Start(instances int, force bool) error
	Restart(force bool) error
	Suspend(force bool) error

	GetTag() (string, error)
	SetTag(tag string, force bool) error

	Env() map[string]string
	SetEnv(name, value string, force bool) error
	DelEnv(name string, force bool) error

	Cpus() float64
	SetCpus(to float64, force bool) error

	Memory() float64
	SetMemory(to float64, force bool) error

	Role() string
	SetRole(to string, force bool) error

	Container() *marathon.Container
	SetContainer(to *marathon.Container, force bool) error

	Parameters() (map[string]string, error)
	AddParameter(key, value string, force bool) error
	DelParameter(key string, force bool) error

	Versions() []string
	LastVersion() string
	Config(version string) *Application

	Load(fileName string) *Application
	Dump(fileName string) error

	Apply(force bool) error
}

// Application is a Marathon Application implementation
type Application struct {
	client  *marathon.Client
	timeout time.Duration

	//
	app *App

	//
	deploy *data.Response
	fail   *data.FailureMessage
}

//=== Marathon FilteredApps JSON Entities definition

// AppDefinition encapsulates the data definitions of a Marathon App
type AppDefinition struct {
	ID                    string                 `json:"id"`
	AcceptedResourceRoles []string               `json:"acceptedResourceRoles,omitempty"`
	BackoffFactor         float64                `json:"backoffFactor,omitempty"`
	BackoffSeconds        int                    `json:"backoffSeconds,omitempty"`
	Container             marathon.Container     `json:"container"`
	Constraints           []TaskConstraints      `json:"constraints,omitempty"`
	Cpus                  float64                `json:"cpus"`
	Disk                  float64                `json:"disk,omitempty"`
	Env                   map[string]string      `json:"env,omitempty"`
	Executor              string                 `json:"executor,omitempty"`
	Fetch                 []Fetch                `json:"fetch,omitempty"`
	HealthChecks          []marathon.Healthcheck `json:"healthChecks,omitempty"`
	Instances             int                    `json:"instances"`
	Labels                map[string]string      `json:"labels,omitempty"`
	MaxLaunchDelaySeconds int                    `json:"maxLaunchDelaySeconds,omitempty"`
	Mem                   float64                `json:"mem"`
	Gpus                  int                    `json:"gpus,omitempty"`
	Networks              []Network              `json:"networks,omitempty"`
	RequirePorts          bool                   `json:"requirePorts,omitempty"`
	UpgradeStrategy       UpgradeStrategy        `json:"upgradeStrategy,omitempty"`
	KillSelection         string                 `json:"killSelection,omitempty"`
	UnreachableStrategy   UnreachableStrategy    `json:"unreachableStrategy,omitempty"`
	Role                  string                 `json:"role,omitempty"`
}

// TaskConstraints is a simple array of strings
type TaskConstraints []string

// Fetch reflects the data used by the sub-element fetch on a Marathon App
type Fetch struct {
	URI        string `json:"uri"`
	Extract    bool   `json:"extract"`
	Executable bool   `json:"executable"`
	Cache      bool   `json:"cache"`
}

// Network reflects the data used by the sub-element network on a Marathon App
type Network struct {
	Mode string `json:"mode"`
}

// UpgradeStrategy reflects the data used by the sub-element ppgradeStrategy on a Marathon App
type UpgradeStrategy struct {
	MaximumOverCapacity   float64 `json:"maximumOverCapacity"`
	MinimumHealthCapacity float64 `json:"minimumHealthCapacity"`
}

// VersionInfo reflects the data used by the sub-element versionInfo on a Marathon App
type VersionInfo struct {
	LastScalingAt      time.Time `json:"lastScalingAt"`
	LastConfigChangeAt time.Time `json:"lastConfigChangeAt"`
}

// UnreachableStrategy reflects the data used by the sub-element unreachableStrategy on a Marathon App
type UnreachableStrategy struct {
	InactiveAfterSeconds int `json:"inactiveAfterSeconds"`
	ExpungeAfterSeconds  int `json:"expungeAfterSeconds"`
}

// AppVersions reflects the data used by the sub-element appVersions on a Marathon App
type AppVersions struct {
	Versions []string `json:"versions"`
}

// NewApplication returns a new instance of Marathon application implementation
func New(client *marathon.Client) *Application {

	if client != nil {
		return &Application{
			client: client,
			app:    &App{},
			deploy: &data.Response{},
			fail:   &data.FailureMessage{},
		}
	}
	return nil
}

//=== Marathon Application methods

// Get allows to establish the internal structures to referenced id
func (ma *Application) Get(id string) *Application {

	if len(id) > 0 {
		ma.clear()

		marathon.Logger.Debug("Application: Get %s %+v", id, ma.app.App)

		path := fmt.Sprintf("%s%s", marathon.APIApps, utilities.DelInitialSlash(id))

		if _, err := ma.client.Session.BodyAsJSON(nil).Get(path, ma.app, ma.fail); err != nil {
			marathon.Logger.Debug("Application: Get failed [%+v]", err)
			ma.clear()
		}
	}
	return ma
}

// Set allows to establish the internal structures from a given app
func (ma *Application) Set(app AppDefinition) *Application {

	if len(app.ID) > 0 {
		marathon.Logger.Debug("Application: Set id = %s", app.ID)
		ma.clear()
		ma.app.App = app
	}
	return ma
}

// Create allows create a Marathon application into server
func (ma *Application) Create(app AppDefinition) *Application {

	if len(app.ID) > 0 {
		marathon.Logger.Debug("Application: Create id = [%s] body = %+v", app.ID, app)

		ma.app.App = app
		_ = ma.Apply(true)
	}
	return ma
}

// Destroy erase a Marathon application from server
func (ma *Application) Destroy() error {

	if len(ma.app.App.ID) > 0 {
		marathon.Logger.Debug("Application: Destroy id = [%s] body = %+v", ma.app.App.ID, ma.app.App)

		path := fmt.Sprintf("%s%s", marathon.APIApps, utilities.DelInitialSlash(ma.app.App.ID))

		ma.clear()
		if _, err := ma.client.Session.BodyAsJSON(nil).Delete(path, ma.deploy, ma.fail); err != nil {
			marathon.Logger.Debug("Application: Destroy failed [%+v]", err)
			return err
		}
		return nil
	}
	return errors.New("app cannot be null nor empty")
}

// Update allows change values into Marathon application
func (ma *Application) Update(app AppDefinition) error {

	if len(app.ID) > 0 {
		marathon.Logger.Debug("Application: Update id = [%s] body = %+v", app.ID, app)

		ma.app.App = app
		return ma.Apply(true)
	}
	return errors.New("app cannot be null nor empty")
}

// Instances return actual instances of a Marathon application
func (ma *Application) Instances() int {

	if len(ma.app.App.ID) > 0 {
		return ma.app.App.Instances
	}
	return -1
}

// Scale allows change instances numbers of a Marathon application
func (ma *Application) Scale(instances int, force bool) error {

	if len(ma.app.App.ID) > 0 {
		marathon.Logger.Debug("Application: Scale %s to %d force=%v", ma.app.App.ID, instances, force)
		ma.app.App.Instances = instances

		return ma.Apply(force)
	}
	return errors.New("app cannot be null nor empty")
}

// Start sets instances of a Marathon application to a number provided
func (ma *Application) Start(instances int, force bool) error {

	return ma.Scale(instances, force)
}

// Stop sets instances of a Marathon application to 0
func (ma *Application) Stop(force bool) error {

	return ma.Scale(0, force)
}

// Restart use an endpoint to trigger a Marathon application restart
func (ma *Application) Restart(force bool) error {

	if len(ma.app.App.ID) > 0 {
		marathon.Logger.Debug("Application: Restart id = [%s] force = %v", ma.app.App.ID, force)

		path := fmt.Sprintf("%s%s/restart", marathon.APIApps, utilities.DelInitialSlash(ma.app.App.ID))

		if force {
			ma.client.Session.AddQueryParam("force", "true")
		}

		if _, err := ma.client.Session.BodyAsJSON(nil).Post(path, ma.deploy, ma.fail); err != nil {
			marathon.Logger.Debug("Application: Restart failed [%+v]", err)
			return err
		}
		return nil
	}
	return errors.New("app cannot be null nor empty")
}

// Suspend is an alias to Stop
func (ma *Application) Suspend(force bool) error {

	return ma.Stop(force)
}

// GetTag allows you to change the version of Docker image
func (ma *Application) GetTag() (string, error) {

	if len(ma.app.App.ID) > 0 {
		re := regexp.MustCompile(marathon.DockerImageRegEx)
		elements := re.FindStringSubmatch(ma.app.App.Container.Docker.Image)

		return elements[len(elements)-1], nil
	}
	return "", errors.New("app cannot be null nor empty")
}

// SetTag allows you to change the version of Docker image
func (ma *Application) SetTag(tag string, force bool) error {

	if len(ma.app.App.ID) > 0 {
		re := regexp.MustCompile(marathon.DockerImageRegEx)
		elements := re.FindStringSubmatch(ma.app.App.Container.Docker.Image)

		ma.app.App.Container.Docker.Image = fmt.Sprintf("%s%s/%s:%s", elements[1], elements[4], elements[6], tag)

		return ma.Apply(force)
	}
	return errors.New("app cannot be null nor empty")
}

// Env returns the Environment Variables of a Marathon application
func (ma *Application) Env() map[string]string {

	if len(ma.app.App.ID) > 0 {

		return ma.app.App.Env
	}
	return nil
}

// SetEnv allows set an environment variable into a Marathon application
func (ma *Application) SetEnv(name, value string, force bool) error {

	if len(ma.app.App.ID) > 0 {

		ma.app.App.Env[name] = value
		return ma.Apply(force)
	}
	return errors.New("app cannot be null nor empty")
}

// DelEnv deletes an environment variable from a Marathon application
func (ma *Application) DelEnv(name string, force bool) error {

	if len(ma.app.App.ID) > 0 {

		delete(ma.app.App.Env, name)
		return ma.Apply(force)
	}
	return errors.New("app cannot be null nor empty")
}

// Cpus returns the amount of cpus from a Marathon application
func (ma *Application) Cpus() float64 {

	if len(ma.app.App.ID) > 0 {

		return ma.app.App.Cpus
	}
	return -1
}

// SetCpus sets the amount of cpus of a Marathon application
func (ma *Application) SetCpus(to float64, force bool) error {

	if len(ma.app.App.ID) > 0 {

		ma.app.App.Cpus = to
		return ma.Apply(force)
	}
	return errors.New("app cannot be null nor empty")
}

// Memory returns the amount of memory from a Marathon application
func (ma *Application) Memory() float64 {

	if len(ma.app.App.ID) > 0 {

		return ma.app.App.Mem
	}
	return -1
}

// SetMemory sets the amount of memory of a Marathon application
func (ma *Application) SetMemory(to float64, force bool) error {

	if len(ma.app.App.ID) > 0 {

		ma.app.App.Mem = to
		return ma.Apply(force)
	}
	return errors.New("app cannot be null nor empty")
}

// Role returns task role of a Marathon application
func (ma *Application) Role() string {

	if len(ma.app.App.ID) > 0 {

		return ma.app.App.Role
	}
	return ""
}

// SetRole sets role of a Marathon application
func (ma *Application) SetRole(to string, force bool) error {

	if len(ma.app.App.ID) > 0 {

		ma.app.App.Role = to
		return ma.Apply(force)
	}
	return errors.New("app cannot be null nor empty")
}

// Container returns the Container information of a Marathon application
func (ma *Application) Container() *marathon.Container {

	if len(ma.app.App.ID) > 0 {

		return &ma.app.App.Container
	}
	return nil
}

// SetContainer sets the Container information of a Marathon application
func (ma *Application) SetContainer(to *marathon.Container, force bool) error {

	if len(ma.app.App.ID) > 0 {

		ma.app.App.Container = marathon.Container{
			Type: to.Type,
			Docker: marathon.Docker{
				ForcePullImage: to.Docker.ForcePullImage,
				Image:          to.Docker.Image,
				Parameters:     to.Docker.Parameters,
				Privileged:     to.Docker.Privileged,
			},
			Volumes:      to.Volumes,
			PortMappings: to.PortMappings,
		}
		return ma.Apply(force)
	}
	return errors.New("app cannot be null nor empty")
}

// Parameters returns all Docker parameters of a Marathon application
func (ma *Application) Parameters() (map[string]string, error) {

	if len(ma.app.App.ID) > 0 {

		if len(ma.app.App.Container.Docker.Parameters) > 0 {

			paramsMap := make(map[string]string)
			for _, val := range ma.app.App.Container.Docker.Parameters {
				paramsMap[val.Key] = val.Value
			}
			return paramsMap, nil
		}
		return nil, fmt.Errorf("the Marathon app %s has no Docker parameters", ma.app.App.ID)
	}
	return nil, errors.New("app cannot be null nor empty")
}

// AddParameter sets the key, value into parameters of a Marathon application
func (ma *Application) AddParameter(key, value string, force bool) error {

	if len(ma.app.App.ID) > 0 {

		exist := false
		index := -1

		// Are there any parameters defined?
		if len(ma.app.App.Container.Docker.Parameters) > 0 {

			// Iterate searching for our passed key
			for idx, val := range ma.app.App.Container.Docker.Parameters {
				if val.Key == key {
					exist = true
					index = idx
				}
			}
		}
		// Exists?
		if exist {
			// Chenge it
			ma.app.App.Container.Docker.Parameters[index].Value = value
		} else {
			// Create, append it
			ma.app.App.Container.Docker.Parameters = append(ma.app.App.Container.Docker.Parameters, marathon.DockerParameters{
				Key:   key,
				Value: value,
			})
			return ma.Apply(force)
		}
	}
	return errors.New("app cannot be null nor empty")
}

// DelParameter erase the parameter referenced by key
func (ma *Application) DelParameter(key string, force bool) error {

	if len(ma.app.App.ID) > 0 {
		toRemove := -1

		if len(ma.app.App.Container.Docker.Parameters) > 0 {
			for index, val := range ma.app.App.Container.Docker.Parameters {
				if val.Key == key {
					toRemove = index
				}
			}
		}
		if toRemove >= 0 {
			length := len(ma.app.App.Container.Docker.Parameters)

			ma.app.App.Container.Docker.Parameters[toRemove] = ma.app.App.Container.Docker.Parameters[length-1]
			ma.app.App.Container.Docker.Parameters[length-1] = marathon.DockerParameters{
				Key:   "",
				Value: "",
			}
			ma.app.App.Container.Docker.Parameters = ma.app.App.Container.Docker.Parameters[:length-1]

			return ma.Apply(force)
		}
		return fmt.Errorf("parameters %s dont exist in Marathon app %s", key, ma.app.App.ID)
	}
	return errors.New("app cannot be null nor empty")
}

// Versions returns all configurations versions of provided task
func (ma *Application) Versions() []string {

	if len(ma.app.App.ID) > 0 {
		marathon.Logger.Debug("Application: Versions %+v", ma.app.App)

		path := fmt.Sprintf(marathon.APIVersions, utilities.DelInitialSlash(ma.app.App.ID))

		versions := &AppVersions{Versions: make([]string, 0)}

		if _, err := ma.client.Session.BodyAsJSON(nil).Get(path, versions, ma.fail); err != nil {
			marathon.Logger.Debug("Application: Versions failed [%+v]", err)
			ma.clear()
		}

		if len(versions.Versions) > 0 {
			return versions.Versions
		}
	}
	return nil
}

// LastVersion returns last version of a provided task
func (ma *Application) LastVersion() string {

	if len(ma.app.App.ID) > 0 {
		if versions := ma.Versions(); len(versions) > 0 {
			sort.Strings(versions)
			return versions[len(versions)-1]
		}
	}
	return ""
}

// Config returns a AppDefinition based on it version
func (ma *Application) Config(version string) *Application {

	if len(ma.app.App.ID) > 0 {
		marathon.Logger.Debug("Application: Config %+v", ma.app.App)

		path := fmt.Sprintf(marathon.APIConfigByVersion, utilities.DelInitialSlash(ma.app.App.ID), version)

		if _, err := ma.client.Session.BodyAsJSON(nil).Get(path, ma.app, ma.fail); err != nil {
			marathon.Logger.Debug("Application: Config failed [%+v]", err)
			ma.clear()
		}
	}
	return ma
}

// Load allows create or update a Marathon application from file
func (ma *Application) Load(fileName string) *Application {

	var err error

	ma.clear()

	switch filepath.Ext(strings.TrimSpace(fileName)) {
	case ".json":
		err = utilities.LoadDataFromJSON(&ma.app.App, fileName)
	case ".yaml":
		err = utilities.LoadDataFromYAML(&ma.app.App, fileName)
	default:
		err = fmt.Errorf("invalid filename extension")
	}

	if err != nil {
		ma.clear()
	}
	return ma
}

// Dump allows to create a .json file with the configuration of a Marathon application
func (ma *Application) Dump(fileName string) (err error) {

	if len(ma.app.App.ID) > 0 {

		switch filepath.Ext(strings.TrimSpace(fileName)) {
		case ".json":
			err = utilities.WriteDataToJSON(ma.app.App, fileName)
		case ".yaml":
			err = utilities.WriteDataToYAML(ma.app.App, fileName)
		default:
			err = utilities.WriteDataToJSON(ma.app.App, fileName)
		}

		return
	}
	return errors.New("app cannot be null nor empty")
}

// Apply internal func, allows send all changes of a Marathon application to Marathon server
func (ma *Application) Apply(force bool) error {

	if len(ma.app.App.ID) > 0 {

		path := fmt.Sprintf("%s%s", marathon.APIApps, utilities.DelInitialSlash(ma.app.App.ID))

		marathon.Logger.Debug("Application: Apply(%v)[%+v] %s", force, ma.app, path)
		if force {
			ma.client.Session.AddQueryParam("force", "true")
		}

		if _, err := ma.client.Session.BodyAsJSON(ma.app.App).Put(path, ma.deploy, ma.fail); err != nil {
			marathon.Logger.Debug("Application: Apply StatusCode: %d [Deploy Id: %s => date: %v {%+v}{%+v}]", ma.client.StatusCode(), ma.deploy.ID, ma.deploy.Version, ma.fail, err)
			return err
		}
		// TODO: Deployment wait for ma.timeout
		marathon.Logger.Debug("Application: Apply StatusCode: %d [Deploy Id: %s => date: %v]", ma.client.StatusCode(), ma.deploy.ID, ma.deploy.Version)

		return nil
	}
	return errors.New("app cannot be null nor empty")
}

// AsRaw returns AppDefinition content of current Application
func (ma *Application) AsRaw() AppDefinition {

	if len(ma.app.App.ID) > 0 {
		return ma.app.App
	}
	return AppDefinition{}
}

// clear set internal data to his defaults
func (ma *Application) clear() {

	ma.app = nil
	ma.app = &App{}
}
