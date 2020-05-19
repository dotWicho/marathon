package marathon

import (
	"errors"
	"fmt"
	"github.com/dotWicho/marathon/pkg/utils"
	"github.com/dotWicho/requist"
	"regexp"
	"time"
)

//===

// App wraps an AppDefinition element returned by the Marathon API
type App struct {
	App AppDefinition `json:"app"`
}

// Apps wraps an AppDefinition array returned by the Marathon API
type Apps struct {
	Apps []AppDefinition `json:"apps"`
}

//===

// Marathon Application interface
type application interface {
	SetClient(client *requist.Requist)

	Get(id string) (*Application, error)
	Create(app AppDefinition) (*Application, error)
	Destroy() error
	Update(app AppDefinition) error

	Scale(instances int, force bool) error
	Stop(force bool) error
	Start(instances int, force bool) error
	Restart(force bool) error
	Suspend(force bool) error

	Retag(tag string) error

	Env() map[string]string
	SetEnv(name, value string) error
	DelEnv(name string) error

	Cpus() float64
	SetCpus(to float64) error

	Memory() float64
	SetMemory(to float64) error

	Role() string
	SetRole(to string) error

	Container() *Container
	SetContainer(to *Container) error

	AddParameter(key, value string) error
	DelParameter(key string) error

	LoadFromFile(fileName string) error
	DumpToFile(fileName string) error

	applyChanges() error
}

// Marathon Application implementation
type Application struct {
	client *requist.Requist
	//
	app *App

	//
	baseUrl string
	auth    string

	//
	deploy *Response
	fail   *FailureMessage
}

//=== Marathon JSON Entities definition

// AppDefinition encapsulates the data definitions of a Marathon App
type AppDefinition struct {
	ID                    string              `json:"id"`
	AcceptedResourceRoles []string            `json:"acceptedResourceRoles,omitempty"`
	BackoffFactor         float64             `json:"backoffFactor,omitempty"`
	BackoffSeconds        int                 `json:"backoffSeconds,omitempty"`
	Container             Container           `json:"container"`
	Cpus                  float64             `json:"cpus"`
	Disk                  float64             `json:"disk,omitempty"`
	Env                   map[string]string   `json:"env,omitempty"`
	Executor              string              `json:"executor,omitempty"`
	Fetch                 []Fetch             `json:"fetch,omitempty"`
	HealthChecks          []Healthcheck       `json:"healthChecks,omitempty"`
	Instances             int                 `json:"instances"`
	Labels                map[string]string   `json:"labels,omitempty"`
	MaxLaunchDelaySeconds int                 `json:"maxLaunchDelaySeconds,omitempty"`
	Mem                   float64             `json:"mem"`
	Gpus                  int                 `json:"gpus,omitempty"`
	Networks              []Network           `json:"networks,omitempty"`
	RequirePorts          bool                `json:"requirePorts,omitempty"`
	UpgradeStrategy       UpgradeStrategy     `json:"upgradeStrategy,omitempty"`
	KillSelection         string              `json:"killSelection,omitempty"`
	UnreachableStrategy   UnreachableStrategy `json:"unreachableStrategy,omitempty"`
	Role                  string              `json:"role,omitempty"`
}

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
	Versions []time.Time `json:"versions"`
}

//=== Marathon Application methods

// NewMarathonApplication returns a new instance of Marathon application
func NewMarathonApplication() *Application {
	ma := &Application{
		client:  nil,
		app:     &App{},
		baseUrl: "",
		auth:    "",
		deploy:  &Response{},
		fail:    &FailureMessage{},
	}
	return ma
}

// SetClient allows reuse of the main object client
func (ma *Application) SetClient(client *requist.Requist) {

	if client != nil {
		ma.client = client
	} else {
		panic(errors.New("client reference cannot be null"))
	}
}

// Get allows to establish the internal structures to referenced id
func (ma *Application) Get(id string) (*Application, error) {

	if len(id) > 0 {
		path := fmt.Sprintf("%s%s", marathonApiApps, utils.DelInitialSlash(id))

		if _, err := ma.client.BodyAsJSON(nil).Get(path, ma.app, ma.fail); err != nil {
			return nil, errors.New(fmt.Sprintf("unable to get add id = %s", id))
		}
		return ma, nil
	}
	return nil, errors.New("id cannot be empty")
}

// Create allows create a Marathon application into server
func (ma *Application) Create(app AppDefinition) (*Application, error) {

	if len(app.ID) > 0 {
		path := fmt.Sprintf("%s%s", marathonApiApps, utils.DelInitialSlash(app.ID))

		if _, err := ma.client.BodyAsJSON(app).Put(path, ma.deploy, ma.fail); err != nil {
			return nil, err
		}
		ma.app = &App{
			App: app,
		}
		return ma, nil
	}
	return nil, errors.New("incorrect application definition")
}

// Destroy erase a Marathon application from server
func (ma *Application) Destroy() error {

	if ma.app != nil {

		path := fmt.Sprintf("%s%s", marathonApiApps, utils.DelInitialSlash(ma.app.App.ID))

		if _, err := ma.client.BodyAsJSON(nil).Delete(path, ma.deploy, ma.fail); err != nil {
			return err
		}
		return nil
	}
	return errors.New("app cannot be null nor empty")
}

// Update allows change values into Marathon application
func (ma *Application) Update(app AppDefinition) error {

	if _, err := ma.client.BodyAsJSON(app).Post(marathonApiApps, ma.deploy, ma.fail); err != nil {
		return err
	}
	return nil
}

// Scale allows change instances numbers of a Marathon application
func (ma *Application) Scale(instances int, force bool) error {

	if ma.app != nil {
		ma.app.App.Instances = instances

		return ma.applyChanges()
	}
	return errors.New("app cannot be null nor empty")
}

// Stop sets instances of a Marathon application to 0
func (ma *Application) Stop(force bool) error {

	return ma.Scale(0, force)
}

// Start sets instances of a Marathon application to a number provided
func (ma *Application) Start(instances int, force bool) error {

	return ma.Scale(instances, force)
}

// Restart use an endpoint to trigger a Marathon application restart
func (ma *Application) Restart(force bool) error {

	if ma.app != nil {
		path := fmt.Sprintf("%s%s/restart", marathonApiApps, utils.DelInitialSlash(ma.app.App.ID))

		if force {
			ma.client.AddQueryParam("force", "true")
		}

		if _, err := ma.client.BodyAsJSON(nil).Patch(path, ma.deploy, ma.fail); err != nil {
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

// Retag allows you to change the version of Docker image
func (ma *Application) Retag(tag string) error {

	if ma.app != nil {
		path := fmt.Sprintf("%s%s", marathonApiApps, utils.DelInitialSlash(ma.app.App.ID))

		ma.client.AddQueryParam("force", "true")

		re := regexp.MustCompile(DockerImageRegEx)
		elements := re.FindStringSubmatch(ma.app.App.Container.Docker.Image)

		ma.app.App.Container.Docker.Image = fmt.Sprintf("%s%s/%s:%s", elements[1], elements[4], elements[6], tag)

		if _, err := ma.client.BodyAsJSON(ma.app.App).Patch(path, ma.deploy, ma.fail); err != nil {
			return err
		}
		return nil
	}
	return errors.New("app cannot be null nor empty")
}

// Env returns the Environment Variables of a Marathon application
func (ma *Application) Env() map[string]string {

	return nil
}

// SetEnv allows set an environment variable into a Marathon application
func (ma *Application) SetEnv(name, value string) error {

	return ma.applyChanges()
}

// DelEnv deletes an environment variable from a Marathon application
func (ma *Application) DelEnv(name string) error {

	return nil
}

// Cpus returns the amount of cpus from a Marathon application
func (ma *Application) Cpus() float64 {

	return ma.app.App.Cpus
}

// SetCpus sets the amount of cpus of a Marathon application
func (ma *Application) SetCpus(to float64) error {

	ma.app.App.Cpus = to
	return ma.applyChanges()
}

// Memory returns the amount of memory from a Marathon application
func (ma *Application) Memory() float64 {

	return ma.app.App.Mem
}

// SetMemory sets the amount of memory of a Marathon application
func (ma *Application) SetMemory(to float64) error {

	ma.app.App.Mem = to
	return ma.applyChanges()
}

// Role returns task role of a Marathon application
func (ma *Application) Role() string {

	return ma.app.App.Role
}

// SetRole sets role of a Marathon application
func (ma *Application) SetRole(to string) error {

	ma.app.App.Role = to
	return ma.applyChanges()
}

// Container returns the Container information of a Marathon application
func (ma *Application) Container() *Container {

	return &ma.app.App.Container
}

// SetContainer sets the Container information of a Marathon application
func (ma *Application) SetContainer(to *Container) error {

	ma.app.App.Container = Container{
		Type: to.Type,
		Docker: Docker{
			ForcePullImage: to.Docker.ForcePullImage,
			Image:          to.Docker.Image,
			Parameters:     to.Docker.Parameters,
			Privileged:     to.Docker.Privileged,
		},
		Volumes:      to.Volumes,
		PortMappings: to.PortMappings,
	}
	return ma.applyChanges()
}

// AddParameter sets the key, value into parameters of a Marathon application
func (ma *Application) AddParameter(key, value string) error {

	return ma.applyChanges()
}

// DelParameter erase the parameter referenced by key
func (ma *Application) DelParameter(key string) error {

	return ma.applyChanges()
}

// LoadFromFile allows create or update a Marathon application from file
func (ma *Application) LoadFromFile(fileName string) error {
	app := &AppDefinition{}

	var err error
	if err = utils.LoadDataFromJson(app, fileName); err == nil {
		_, err = ma.Create(*app)
		return err
	}
	return nil
}

// DumpToFile allows to create a .json file with the configuration of a Marathon application
func (ma *Application) DumpToFile(fileName string) error {

	if err := utils.WriteDataToJson(ma.app.App, fileName); err != nil {
		return err
	}
	return nil
}

// applyChanges internal func, allows send all changes of a Marathon application to Marathon server
func (ma *Application) applyChanges() error {

	if ma.app != nil {
		path := fmt.Sprintf("%s%s", marathonApiApps, utils.DelInitialSlash(ma.app.App.ID))

		ma.client.AddQueryParam("force", "true")

		if _, err := ma.client.BodyAsJSON(ma.app.App).Patch(path, ma.deploy, ma.fail); err != nil {
			return err
		}
		return nil
	}
	return errors.New("app cannot be null nor empty")
}
