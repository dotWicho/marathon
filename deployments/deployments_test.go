package deployments

import (
	"github.com/dotWicho/requist"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestApplication_New(t *testing.T) {

	t.Run("nil Client if send invalid baseURL", func(t *testing.T) {

		// Try to create a Marathon Deployments reference
		_client := New()

		// Marathon Deployments is not nil
		assert.NotNil(t, _client)
	})
}

func TestApplication_SetClient(t *testing.T) {

	t.Run("nil Client if send invalid baseURL", func(t *testing.T) {

		// Try to create a Marathon Deployments reference
		_client := New()

		// must be not nil
		assert.NotNil(t, _client)

		_err := _client.SetClient(nil)

		assert.NotNil(t, _err)
		assert.Equal(t, "client reference cannot be null", _err.Error())
	})

	t.Run("valid Client if send valid baseURL", func(t *testing.T) {

		// Try to create a Marathon Deployments reference
		_client := New()

		// must be not nil
		assert.NotNil(t, _client)

		_err := _client.SetClient(requist.New("http://192.168.0.1:8001"))

		assert.Nil(t, _err)
	})
}
