package marathon

import (
	"errors"
	"github.com/dotWicho/logger"
	"github.com/dotWicho/requist"
	"net/url"
	"strings"
	"time"
)

// Logger default
var Logger = logger.NewLogger(false)

// Actions commons actions
type Actions interface {
	SetClient(*requist.Requist) error
}

// Marathon application interface
type client interface {
	New(base *url.URL) *Marathon
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
}

// Marathon application implementation
type Marathon struct {
	Client  *requist.Requist
	timeout time.Duration

	//
	info *Info

	//
	fail FailureMessage

	//
	auth    string
	baseURL string
}

// NewClient returns a new Marathon given a Marathon server base url
func New(base string) *Marathon {

	baseURL, err := url.Parse(base)

	if len(base) == 0 || err != nil {
		Logger.Debug("[marathon] Invalid baseURL")
		return nil
	}

	return new(Marathon).New(baseURL)
}

// NewClientFromURL returns a new Marathon given a Marathon server base url in URL type
func NewClientFromURL(base *url.URL) *Marathon {

	if baseStr := base.String(); len(baseStr) > 0 {
		Logger.Debug("[marathon] Creating Marathon Marathon from url.URL = %s", base.String())
		baseURL, err := url.Parse(baseStr)
		if (baseURL != nil && !strings.HasPrefix(baseURL.Scheme, "http")) || err != nil {
			Logger.Debug("[marathon] Invalid baseURL")
			return nil
		}
		return new(Marathon).New(baseURL)
	}
	return nil
}

//=== Marathon utilities definitions ===

// New returns a Marathon populated struct
func (mc *Marathon) New(base *url.URL) *Marathon {

	mc.Client = requist.New(base.String())
	mc.baseURL = base.String()
	mc.info = &Info{}

	if base.User.String() != "" {
		if pass, check := base.User.Password(); check {
			mc.Client.SetBasicAuth(base.User.Username(), pass)
		}
		mc.auth = mc.Client.GetBasicAuth()
	}
	mc.SetTimeout(DeploymentTimeout)
	mc.Client.Accept("application/json")
	mc.Client.SetHeader("Cache-Control", "no-cache")
	mc.Client.SetHeader("Accept-Encoding", "identity")

	return mc
}

// Connect sets baseURL and prepares the Marathon with this
func (mc *Marathon) Connect(baseUrl string) {
	mc.Client = requist.New(baseUrl)
}

// StatusCode returns last responseCode
func (mc *Marathon) StatusCode() int {
	return mc.Client.StatusCode()
}

// CheckConnection send a request to check Marathon server connectivity
func (mc *Marathon) CheckConnection() error {

	if _, err := mc.Client.Get(APIPing, nil, nil); err != nil {
		return err
	}
	if mc.StatusCode() == 200 {
		if _, err := mc.Client.Get(APIInfo, mc.info, mc.fail); err != nil {
			return err
		}
	} else {
		return errors.New("unable to connect")
	}
	return nil
}

// SetBasicAuth used if we need to set login parameters
func (mc *Marathon) SetTimeout(timeout time.Duration) {

	mc.timeout = timeout
	mc.Client.SetClientTimeout(mc.timeout)
}

// SetBasicAuth used if we need to set login parameters
func (mc *Marathon) SetBasicAuth(username, password string) {

	mc.Client.SetBasicAuth(username, password)
}

//=== Marathon Info interface definitions ===

// MarathonVersion returns version of Marathon
func (mc *Marathon) MarathonVersion() string {

	return mc.info.Version
}

// MarathonLeader returns actual Marathon leader server
func (mc *Marathon) MarathonLeader() string {

	return mc.info.Leader
}

// MarathonFramework returns the id of this Marathon on Mesos
func (mc *Marathon) MarathonFramework() string {

	return mc.info.FrameworkID
}

// MarathonZookeeper return Zookeeper server(s) address
func (mc *Marathon) MarathonZookeeper() string {

	return mc.info.ZookeeperConfig.Zk
}
