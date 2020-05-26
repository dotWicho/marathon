package marathon

import (
	"errors"
	"fmt"
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
}

// Marathon application implementation
type Client struct {
	Session *requist.Requist
	timeout time.Duration

	//
	info Info

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
	marathon.Session = requist.New(base.String())
	marathon.baseUrl = base.String()

	if base.User.String() != "" {
		if pass, check := base.User.Password(); check {
			marathon.Session.SetBasicAuth(base.User.Username(), pass)
		}
		marathon.auth = marathon.Session.GetBasicAuth()
	}
	marathon.SetTimeout(defaultDeploymentTimeout)
	marathon.Session.Accept("application/json")
	marathon.Session.SetHeader("Cache-Control", "no-cache")
	marathon.Session.SetHeader("Accept-Encoding", "identity")

	return marathon
}

// Connect sets baseUrl and prepares the Client with this
func (mc *Client) Connect(baseUrl string) {
	mc.Session = nil
	nc := mc.Session.New(baseUrl)
	mc.Session = nc
}

// StatusCode returns last responseCode
func (mc *Client) StatusCode() int {
	return mc.Session.StatusCode()
}

// CheckConnection send a request to check Marathon server connectivity
func (mc *Client) CheckConnection() error {

	if _, err := mc.Session.Get(marathonApiPing, nil, nil); err != nil {
		return errors.New(fmt.Sprintf("unable to connect to Marathon server %s.\n", mc.baseUrl))
	}
	if mc.StatusCode() == 200 {
		if _, err := mc.Session.Get(marathonApiInfo, mc.info, mc.fail); err != nil {
			return errors.New(fmt.Sprintf("unable to get info from Marathon server %s.\n", mc.baseUrl))
		}
	}
	return nil
}

// SetBasicAuth used if we need to set login parameters
func (mc *Client) SetTimeout(timeout time.Duration) {

	mc.timeout = timeout
	mc.Session.SetClientTimeout(mc.timeout)
}

// SetBasicAuth used if we need to set login parameters
func (mc *Client) SetBasicAuth(username, password string) {

	mc.Session.SetBasicAuth(username, password)
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
