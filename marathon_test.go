package marathon

import (
	"github.com/dotWicho/marathon/mockserver"
	"github.com/stretchr/testify/assert"
	"net/url"
	"testing"
	"time"
)

func Test_New(t *testing.T) {

	t.Run("nil Client if send invalid baseURL", func(t *testing.T) {

		// Try to create Application
		_client := New("")

		// Application is nil
		assert.Nil(t, _client)
	})

	t.Run("valid Client if send valid baseURL", func(t *testing.T) {

		// Try to create Application
		_client := New("http://127.0.0.1:8080")

		// Application is nil
		assert.NotNil(t, _client)
	})
}

func Test_NewFromURL(t *testing.T) {

	t.Run("nil Client if send invalid baseURL", func(t *testing.T) {

		// We define some vars
		baseURL := &url.URL{
			Scheme:     "",
			Opaque:     "",
			User:       nil,
			Host:       "",
			Path:       "",
			RawPath:    "",
			ForceQuery: false,
			RawQuery:   "",
			Fragment:   "",
		}

		// Try to create Application
		_client := NewFromURL(baseURL)

		// Application is nil
		assert.Nil(t, _client)
	})

	t.Run("valid Client if send valid baseURL", func(t *testing.T) {

		// We define some vars
		baseURL := &url.URL{
			Scheme:     "https",
			Opaque:     "",
			User:       url.UserPassword("anonymous", "password"),
			Host:       "127.0.0.1:8080",
			Path:       "",
			RawPath:    "",
			ForceQuery: false,
			RawQuery:   "",
			Fragment:   "",
		}

		// Try to create Application
		_client := NewFromURL(baseURL)

		// Application is nil
		assert.NotNil(t, _client)
	})
}

func TestClient_New(t *testing.T) {

	t.Run("nil Client if send invalid baseURL", func(t *testing.T) {

		// We define some vars
		baseURL := &url.URL{
			Scheme:     "",
			Opaque:     "",
			User:       nil,
			Host:       "",
			Path:       "",
			RawPath:    "",
			ForceQuery: false,
			RawQuery:   "",
			Fragment:   "",
		}

		// Define a fake client
		aClient := &Client{}

		// Try to create Application
		_client := aClient.New(baseURL)

		// Application is nil
		assert.Nil(t, _client)
	})

	t.Run("valid Client if send valid baseURL", func(t *testing.T) {

		// We define some vars
		baseURL := &url.URL{
			Scheme:     "https",
			Opaque:     "",
			User:       nil,
			Host:       "127.0.0.1:8080",
			Path:       "",
			RawPath:    "",
			ForceQuery: false,
			RawQuery:   "",
			Fragment:   "",
		}

		// Define a fake client
		aClient := &Client{}

		// Try to create Application
		_client := aClient.New(baseURL)

		// Application is nil
		assert.NotNil(t, _client)
	})
}

func TestClient_Connect(t *testing.T) {

	t.Run("nil Client.Session if send invalid baseURL", func(t *testing.T) {

		// We define some vars
		baseURL := &url.URL{
			Scheme:     "",
			Opaque:     "",
			User:       nil,
			Host:       "",
			Path:       "",
			RawPath:    "",
			ForceQuery: false,
			RawQuery:   "",
			Fragment:   "",
		}

		// Define a fake client
		aClient := &Client{}

		// Try to create Application
		aClient.Connect(baseURL.String())

		// Application is nil
		assert.Nil(t, aClient.Session)
	})

	t.Run("valid Client.Session if send valid baseURL", func(t *testing.T) {

		// We define some vars
		baseURL := &url.URL{
			Scheme:     "https",
			Opaque:     "",
			User:       nil,
			Host:       "127.0.0.1:8080",
			Path:       "",
			RawPath:    "",
			ForceQuery: false,
			RawQuery:   "",
			Fragment:   "",
		}

		// Define a fake client
		aClient := &Client{}

		// Try to create Application
		aClient.Connect(baseURL.String())

		// Application is nil
		assert.NotNil(t, aClient.Session)
	})
}

func TestClient_CheckConnection(t *testing.T) {

	// We create a Mock Server
	server := mockserver.MockServer()
	defer server.Close()

	// Try to create Client
	_client := New(server.URL)

	// Client is not nil
	assert.NotNil(t, _client)

	// Fire up CheckConnection
	err := _client.CheckConnection()

	// We get nil error
	assert.Nil(t, err)

	// Check some values
	assert.Equal(t, "v1.0.0", _client.Version())
	assert.Equal(t, "127.0.0.10:8080", _client.Leader())
	assert.Equal(t, "97c136bf-5a28-4821-9d94-480d9fbb01c8", _client.Framework())
	assert.Equal(t, "127.0.0.10:2181", _client.Zookeeper())

}

func TestClient_SetTimeout(t *testing.T) {

	// We define some vars
	timeout := 10 * time.Second

	// Try to create Client
	_client := New("http://127.0.0.1:8080")

	// Client is not nil
	assert.NotNil(t, _client)

	// We set to 10 seconds our default Timeout
	_client.SetTimeout(timeout)

	// was modified out Client?
	assert.NotNil(t, _client)

	// if not empty our default Timeout
	assert.NotEmpty(t, _client.timeout)

	// We have a Timeout set?
	assert.Equal(t, timeout, _client.timeout)
}

func TestClient_SetBasicAuth(t *testing.T) {

	// Try to create Client
	_client := New("http://127.0.0.1:8080")

	// was modified out Client?
	assert.NotNil(t, _client)

	t.Run("get empty Auth if set empty Username and empty Password", func(t *testing.T) {

		// We set some variables
		username := ""
		password := ""
		expected := ""

		// Set empty Basic Auth
		_client.SetBasicAuth(username, password)

		// was modified out Client?
		assert.NotNil(t, _client)

		// our data is correct?
		assert.EqualValues(t, expected, _client.auth)
	})

	t.Run("get empty Auth if set valid Username and empty Password", func(t *testing.T) {

		// We set some variables
		username := "anonymous"
		password := ""
		expected := ""

		_client.SetBasicAuth(username, password)

		// was modified out Client?
		assert.NotNil(t, _client)

		// our data is correct?
		assert.EqualValues(t, expected, _client.auth)
	})

	t.Run("get empty Auth if set empty Username and valid Password", func(t *testing.T) {

		// We set some variables
		username := ""
		password := "Password123"
		expected := ""

		_client.SetBasicAuth(username, password)

		// was modified out Client?
		assert.NotNil(t, _client)

		// our data is correct?
		assert.EqualValues(t, expected, _client.auth)
	})

	t.Run("get valid Auth if set valid Username and valid Password", func(t *testing.T) {

		// We set some variables
		username := "anonymous"
		password := "Password123"
		expected := "anonymous:Password123"

		_client.SetBasicAuth(username, password)

		// was modified out Client?
		assert.NotNil(t, _client)

		// our data is correct?
		assert.EqualValues(t, expected, _client.auth)
	})
}
