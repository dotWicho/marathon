package marathon

import (
	"errors"
	"github.com/dotWicho/requist"
	"net/url"
	"time"
)

// Marathon application interface
type client interface {
	New(base *url.URL) *Client
	Connect(baseUrl string)
	StatusCode() int
	CheckConnection() error
	SetTimeout(timeout time.Duration)
	SetBasicAuth(username, password string)

	// Marathon Info interface
	MarathonVersion() string
	MarathonLeader() string
	MarathonFramework() string
	MarathonZookeeper() string

	// Marathon AppDefinition interface
	AppCreate(app AppDefinition) error
	AppDestroy(id string) error
	AppUpdate(app AppDefinition) error

	AppScale(id string, instances int, force bool) error
	AppStop(id string, force bool) error
	AppStart(id string, instances int, force bool) error
	AppRestart(id string, force bool) error
	AppSuspend(id string, force bool) error

	AppGetTag(id string) (string, error)
	AppSetTag(id string, tag string, force bool) error

	AppEnv(id string) map[string]string
	AppSetEnv(id string, name, value string, force bool) error
	AppDelEnv(id string, name string, force bool) error

	AppCpus(id string) float64
	AppSetCpus(id string, to float64, force bool) error

	AppMemory(id string) float64
	AppSetMemory(id string, to float64, force bool) error

	AppRole(id string) string
	AppSetRole(id string, to string, force bool) error

	AppContainer(id string) *Container
	AppSetContainer(id string, to *Container, force bool) error

	AppParams(id string) (map[string]string, error)
	AppAddParameter(id string, key, value string, force bool) error
	AppDelParameter(id string, key string, force bool) error

	AppLoadFromFile(fileName string) error
	AppDumpToFile(id, fileName string) error
}

// Marathon application implementation
type Client struct {
	client  *requist.Requist
	timeout time.Duration

	//
	info *Info

	//
	ma *Application
	mg *Groups
	md *Deployment
	mt *Tasks

	//
	fail FailureMessage

	//
	auth    string
	baseUrl string
}

// NewClient returns a new Client given a Marathon server base url
func New(base string) *Client {

	baseURL, err := url.Parse(base)
	if err != nil {
		panic(err)
	}

	client := &Client{}
	return client.New(baseURL)
}

// NewClientFromURL returns a new Client given a Marathon server base url in URL type
func NewClientFromURL(base *url.URL) *Client {

	baseURL, err := url.Parse(base.String())
	if err != nil {
		panic(err)
	}

	client := &Client{}
	return client.New(baseURL)
}

//=== Marathon utilities definitions ===

// New returns a Client populated struct
func (mc *Client) New(base *url.URL) *Client {

	marathon := &Client{}
	marathon.client = requist.New(base.String())
	marathon.baseUrl = base.String()
	marathon.info = &Info{}

	if base.User.String() != "" {
		if pass, check := base.User.Password(); check {
			marathon.client.SetBasicAuth(base.User.Username(), pass)
		}
		marathon.auth = marathon.client.GetBasicAuth()
	}
	marathon.SetTimeout(defaultDeploymentTimeout)
	marathon.client.Accept("application/json")
	marathon.client.SetHeader("Cache-Control", "no-cache")
	marathon.client.SetHeader("Accept-Encoding", "identity")

	marathon.ma = NewMarathonApplication(mc.timeout)
	marathon.ma.client = marathon.client
	marathon.ma.auth = marathon.auth

	marathon.mg = nil
	marathon.md = nil
	marathon.mt = nil

	return marathon
}

// Connect sets baseUrl and prepares the Client with this
func (mc *Client) Connect(baseUrl string) {
	mc.client = nil
	nc := mc.client.New(baseUrl)
	mc.client = nc
}

// StatusCode returns last responseCode
func (mc *Client) StatusCode() int {
	return mc.client.StatusCode()
}

// CheckConnection send a request to check Marathon server connectivity
func (mc *Client) CheckConnection() error {

	if _, err := mc.client.Get(marathonApiPing, nil, nil); err != nil {
		return err
	}
	if mc.StatusCode() == 200 {
		if _, err := mc.client.Get(marathonApiInfo, mc.info, mc.fail); err != nil {
			return err
		}
	} else {
		return errors.New("unable to connect")
	}
	return nil
}

// SetBasicAuth used if we need to set login parameters
func (mc *Client) SetTimeout(timeout time.Duration) {

	mc.timeout = timeout
	mc.client.SetClientTimeout(mc.timeout)
}

// SetBasicAuth used if we need to set login parameters
func (mc *Client) SetBasicAuth(username, password string) {

	mc.client.SetBasicAuth(username, password)
}

//=== Marathon Info interface definitions ===

// MarathonVersion returns version of Marathon
func (mc *Client) MarathonVersion() string {

	return mc.info.Version
}

// MarathonLeader returns actual Marathon leader server
func (mc *Client) MarathonLeader() string {

	return mc.info.Leader
}

// MarathonFramework returns the id of this Marathon on Mesos
func (mc *Client) MarathonFramework() string {

	return mc.info.FrameworkID
}

// MarathonZookeeper return Zookeeper server(s) address
func (mc *Client) MarathonZookeeper() string {

	return mc.info.ZookeeperConfig.Zk
}

//=== Marathon AppDefinition interface definitions ===

// Marathon AppCreate calls MarathonApplication.Create
func (mc *Client) AppCreate(app AppDefinition) error {

	_, err := mc.ma.Create(app)

	return err
}

// Marathon AppDestroy calls MarathonApplication.Destroy
func (mc *Client) AppDestroy(id string) error {

	if _, err := mc.ma.Get(id); err != nil {
		return err
	}
	return mc.ma.Destroy()
}

// Marathon AppUpdate calls MarathonApplication.Update
func (mc *Client) AppUpdate(app AppDefinition) error {

	return mc.ma.Update(app)
}

// Marathon AppScale calls MarathonApplication.Scale
func (mc *Client) AppScale(id string, instances int, force bool) error {

	if _, err := mc.ma.Get(id); err != nil {
		return err
	}
	return mc.ma.Scale(instances, force)
}

// Marathon AppStop calls MarathonApplication.Stop
func (mc *Client) AppStop(id string, force bool) error {

	if _, err := mc.ma.Get(id); err != nil {
		return err
	}
	return mc.ma.Stop(force)
}

// Marathon AppStart calls MarathonApplication.Start
func (mc *Client) AppStart(id string, instances int, force bool) error {

	if _, err := mc.ma.Get(id); err != nil {
		return err
	}
	return mc.ma.Start(instances, force)
}

// Marathon AppRestart calls MarathonApplication.Restart
func (mc *Client) AppRestart(id string, force bool) error {

	if _, err := mc.ma.Get(id); err != nil {
		return err
	}
	return mc.ma.Restart(force)
}

// Marathon AppSuspend calls MarathonApplication.Suspend
func (mc *Client) AppSuspend(id string, force bool) error {

	if _, err := mc.ma.Get(id); err != nil {
		return err
	}
	return mc.ma.Suspend(force)
}

// Marathon AppGetTag calls MarathonApplication.GetTag
func (mc *Client) AppGetTag(id string) (string, error) {

	if _, err := mc.ma.Get(id); err != nil {
		return "", err
	}
	return mc.ma.GetTag()
}

// Marathon AppSetTag calls MarathonApplication.SetTag
func (mc *Client) AppSetTag(id string, tag string, force bool) error {

	if _, err := mc.ma.Get(id); err != nil {
		return err
	}
	return mc.ma.SetTag(tag, force)
}

// Marathon AppEnv calls MarathonApplication.Env
func (mc *Client) AppEnv(id string) map[string]string {

	if _, err := mc.ma.Get(id); err != nil {
		return nil
	}
	return mc.ma.Env()
}

// Marathon AppSetEnv calls MarathonApplication.SetEnv
func (mc *Client) AppSetEnv(id, name, value string, force bool) error {

	if _, err := mc.ma.Get(id); err != nil {
		return err
	}
	return mc.ma.SetEnv(name, value, force)
}

// Marathon AppDelEnv calls MarathonApplication.DelEnv
func (mc *Client) AppDelEnv(id, name string, force bool) error {

	if _, err := mc.ma.Get(id); err != nil {
		return err
	}
	return mc.ma.DelEnv(name, force)
}

// Marathon AppCpus calls MarathonApplication.Cpus
func (mc *Client) AppCpus(id string) float64 {

	if _, err := mc.ma.Get(id); err != nil {
		return 0
	}
	return mc.ma.Cpus()
}

// Marathon AppSetCpus calls MarathonApplication.SetCpus
func (mc *Client) AppSetCpus(id string, to float64, force bool) error {

	if _, err := mc.ma.Get(id); err != nil {
		return err
	}
	return mc.ma.SetCpus(to, force)
}

// Marathon AppMemory calls MarathonApplication.Memory
func (mc *Client) AppMemory(id string) float64 {

	if _, err := mc.ma.Get(id); err != nil {
		return 0
	}
	return mc.ma.Memory()
}

// Marathon AppSetMemory calls MarathonApplication.SetMemory
func (mc *Client) AppSetMemory(id string, to float64, force bool) error {

	if _, err := mc.ma.Get(id); err != nil {
		return err
	}
	return mc.ma.SetCpus(to, force)
}

// Marathon AppRole calls MarathonApplication.Role
func (mc *Client) AppRole(id string) string {

	if _, err := mc.ma.Get(id); err != nil {
		return ""
	}
	return mc.ma.Role()
}

// Marathon AppSetRole calls MarathonApplication.SetRole
func (mc *Client) AppSetRole(id, to string, force bool) error {

	if _, err := mc.ma.Get(id); err != nil {
		return err
	}
	return mc.ma.SetRole(to, force)
}

// Marathon AppContainer calls MarathonApplication.Container
func (mc *Client) AppContainer(id string) *Container {

	if _, err := mc.ma.Get(id); err != nil {
		return nil
	}
	return mc.ma.Container()
}

// Marathon AppSetContainer calls MarathonApplication.SetContainer
func (mc *Client) AppSetContainer(id string, to *Container, force bool) error {

	if _, err := mc.ma.Get(id); err != nil {
		return err
	}
	return mc.ma.SetContainer(to, force)
}

// Marathon AppParams calls MarathonApplication.Parameters
func (mc *Client) AppParams(id string) (map[string]string, error) {

	if _, err := mc.ma.Get(id); err != nil {
		return nil, err
	}
	return mc.ma.Parameters(), nil
}

// Marathon AppAddParameter calls MarathonApplication.AddParameter
func (mc *Client) AppAddParameter(id string, key, value string, force bool) error {

	if _, err := mc.ma.Get(id); err != nil {
		return err
	}
	return mc.ma.AddParameter(key, value, force)
}

// Marathon AppDelParameter calls MarathonApplication.DelParameter
func (mc *Client) AppDelParameter(id string, key string, force bool) error {

	if _, err := mc.ma.Get(id); err != nil {
		return err
	}
	return mc.ma.DelParameter(key, force)
}

// Marathon AppLoadFromFile calls MarathonApplication.LoadFromFile
func (mc *Client) AppLoadFromFile(fileName string) error {

	return mc.ma.LoadFromFile(fileName)
}

// Marathon AppDumpToFile calls MarathonApplication.DumpToFile
func (mc *Client) AppDumpToFile(id, fileName string) error {

	if _, err := mc.ma.Get(id); err != nil {
		return err
	}
	return mc.ma.DumpToFile(fileName)
}
