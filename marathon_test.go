package marathon

import (
	"github.com/stretchr/testify/assert"
	"net/url"
	"testing"
)

func Test_New(t *testing.T) {

	t.Run("nil Marathon if send invalid baseURL", func(t *testing.T) {

		// Try to create Marathon Marathon
		_client := New("")

		// Marathon Marathon is nil
		assert.Nil(t, _client)
	})

	t.Run("valid Marathon if send valid baseURL", func(t *testing.T) {

		// Try to create an Apis reference
		_client := New("http://127.0.0.1:8001")

		// Marathon Marathon is nil
		assert.NotNil(t, _client)
	})
}

func TestNewClientFromURL(t *testing.T) {

	t.Run("nil Marathon if send empty baseURL", func(t *testing.T) {

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

		// Try to create Marathon Marathon
		_client := NewClientFromURL(baseURL)

		// Marathon Marathon is nil
		assert.Nil(t, _client)
	})

	t.Run("nil Marathon if send invalid baseURL", func(t *testing.T) {

		// We define some vars
		baseURL := &url.URL{
			Scheme:     "file",
			Opaque:     "",
			User:       nil,
			Host:       "127.0.0.1:8000",
			Path:       "",
			RawPath:    "",
			ForceQuery: false,
			RawQuery:   "",
			Fragment:   "",
		}

		// Try to create Marathon Marathon
		_client := NewClientFromURL(baseURL)

		// Marathon Marathon is nil
		assert.Nil(t, _client)
	})

	t.Run("valid Marathon if send valid baseURL", func(t *testing.T) {

		// We define some vars
		baseURL := &url.URL{
			Scheme:     "https",
			Opaque:     "",
			User:       url.UserPassword("anonymous", "password"),
			Host:       "127.0.0.1:8000",
			Path:       "",
			RawPath:    "",
			ForceQuery: false,
			RawQuery:   "",
			Fragment:   "",
		}

		// Try to create Marathon Marathon
		_client := NewClientFromURL(baseURL)

		// Marathon Marathon is nil
		assert.NotNil(t, _client)
	})
}

func TestClient_Connect(t *testing.T) {

	t.Run("nil Marathon if send invalid baseURL", func(t *testing.T) {

		// We define some vars
		baseURL := "http://127.0.0.1:8000"

		// Try to create Marathon Marathon
		_client := New(baseURL)

		// Marathon Marathon is not nil
		assert.NotNil(t, _client)

		_client.Connect("")

		// Requist Marathon is nil
		assert.Nil(t, _client.Client)
	})

	t.Run("valid Marathon if send valid baseURL", func(t *testing.T) {

		// We define some vars
		initialURL := "http://192.168.0.1:8000"
		baseURL := "http://127.0.0.1:8000"

		// Try to create Marathon Marathon
		_client := New(initialURL)

		// Marathon Marathon is not nil
		assert.NotNil(t, _client)

		_client.Connect(baseURL)

		// Requist Marathon is nil
		assert.NotNil(t, _client.Client)
	})
}