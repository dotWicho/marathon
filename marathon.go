package marathon

import (
	"fmt"
	"github.com/dotWicho/logger"
	"github.com/dotWicho/marathon/data"
	"github.com/dotWicho/requist"
	"net/url"
	"time"
)

// Logger default
var Logger *logger.StandardLogger = logger.NewLogger(true)

// client Marathon application interface
type client interface {
	New(base *url.URL) *Client
	Connect(baseURL string)
	StatusCode() int
	CheckConnection() error
	SetTimeout(timeout time.Duration)
	SetBasicAuth(username, password string)

	// Marathon Info interface
	Version() string
	Leader() string
	Framework() string
	Zookeeper() string
}

// Client is implementation of Marathon application interface
type Client struct {
	Session *requist.Requist
	timeout time.Duration

	//
	info *data.Info

	//
	fail *data.FailureMessage

	//
	auth    string
	baseURL string
}

// New returns a new Client given a Marathon server base url
func New(base string) *Client {

	Logger.Debug("Creating Marathon Client with baseURL = %s", base)
	baseURL, err := url.Parse(base)
	if len(base) == 0 || err != nil {
		Logger.Debug("Invalid baseURL")
		return nil
	}

	_client := &Client{}
	return _client.New(baseURL)
}

// NewFromURL returns a new Client given a Marathon server base url in URL type
func NewFromURL(base *url.URL) *Client {

	if baseStr := base.String(); len(baseStr) > 0 {

		Logger.Debug("Creating Marathon Client from url.URL = %s", base.String())
		baseURL, err := url.Parse(baseStr)
		if err != nil {
			Logger.Debug("Invalid baseURL")
			return nil
		}

		_client := &Client{}
		return _client.New(baseURL)
	}
	return nil
}

//=== Marathon utilities definitions ===

// New returns a Client populated struct
func (mc *Client) New(base *url.URL) *Client {

	marathon := mc
	marathon.Session = requist.New(base.String())

	if marathon.Session != nil {
		requist.Logger = Logger
		marathon.baseURL = base.String()
		marathon.info = &data.Info{}
		marathon.fail = &data.FailureMessage{}

		if base.User.String() != "" {
			if pass, check := base.User.Password(); check {
				marathon.Session.SetBasicAuth(base.User.Username(), pass)
			}
			marathon.auth = marathon.Session.GetBasicAuth()
		}
		marathon.SetTimeout(DeploymentTimeout)
		marathon.Session.Accept("application/json")
		marathon.Session.SetHeader("Cache-Control", "no-cache")
		marathon.Session.SetHeader("Accept-Encoding", "identity")

		Logger.Debug("Marathon Client = %+v", marathon)
		return marathon
	}
	return nil
}

// Connect sets baseURL and prepares the Client with this
func (mc *Client) Connect(baseURL string) {
	mc.Session = requist.New(baseURL)
}

// StatusCode returns last responseCode
func (mc *Client) StatusCode() int {
	return mc.Session.StatusCode()
}

// CheckConnection send a request to check Marathon server connectivity
func (mc *Client) CheckConnection() error {

	if _, err := mc.Session.Get(APIPing, nil, nil); err != nil {
		Logger.Debug("CheckConnection unable to connect to Marathon server")
		return fmt.Errorf("unable to connect to Marathon server %s", mc.baseURL)
	}
	if mc.StatusCode() == 200 {
		Logger.Debug("CheckConnection successful")
		if _, err := mc.Session.Get(APIInfo, mc.info, mc.fail); err != nil {
			return fmt.Errorf("unable to get info from Marathon server %s", mc.baseURL)
		}
		Logger.Debug("CheckConnection: Marathon version = %s", mc.info.Version)
	}
	return nil
}

// SetTimeout used if we need to set login parameters
func (mc *Client) SetTimeout(timeout time.Duration) {

	mc.timeout = timeout
	mc.Session.SetClientTimeout(mc.timeout)
}

// SetBasicAuth used if we need to set login parameters
func (mc *Client) SetBasicAuth(username, password string) {

	mc.Session.SetBasicAuth(username, password)
	mc.auth = mc.Session.GetBasicAuth()
}

//=== Marathon Info interface definitions ===

// MarathonVersion returns version of Marathon
func (mc *Client) Version() string {

	return mc.info.Version
}

// MarathonLeader returns actual Marathon leader server
func (mc *Client) Leader() string {

	return mc.info.Leader
}

// MarathonFramework returns the id of this Marathon on Mesos
func (mc *Client) Framework() string {

	return mc.info.FrameworkID
}

// MarathonZookeeper return Zookeeper server(s) address
func (mc *Client) Zookeeper() string {

	return mc.info.ZookeeperConfig.Zk
}
