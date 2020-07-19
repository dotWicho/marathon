package marathon

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func Test_NewDeployments(t *testing.T) {

	t.Run("nil Deployment if send nil client", func(t *testing.T) {

		// Try to create Deployments
		_deploy := NewDeployments(nil)

		// Deployments is nil
		assert.Nil(t, _deploy)
	})

	t.Run("valid Deployment if send valid client", func(t *testing.T) {

		// Try to create Deployments
		_deploy := NewDeployments(New("http://127.0.0.1:8080"))

		// Deployments is not nil
		assert.NotNil(t, _deploy)
	})
}

func TestDeployments_Get(t *testing.T) {

	// We create a Mock Server
	server := MockMarathonServer()
	defer server.Close()

	t.Run("get empty array of Deployments", func(t *testing.T) {

		// Try to create Deployment
		_deploy := NewDeployments(New(server.URL))

		// try to Get with empty app id
		_refdeploy, err := _deploy.Get()

		// err must be nil
		assert.Nil(t, err)

		// Application is equal after Get fire up
		assert.Equal(t, _refdeploy, _deploy)

		// Application ref is empty
		assert.Empty(t, _deploy.deployments)
	})

	t.Run("get valid Deployments if exists", func(t *testing.T) {

		deployArray = someDeployments

		// Try to create Deployment
		_deploy := NewDeployments(New(server.URL))

		// try to Get with empty app id
		_refDeploy, err := _deploy.Get()

		// err must be nil
		assert.Nil(t, err)

		// Application is equal after Get fire up
		assert.Equal(t, _refDeploy, _deploy)

		// Application ref is empty
		assert.NotEmpty(t, _deploy.deployments)

		// Check some values on response
		assert.Equal(t, "97c136bf-5a28-4821-9d94-480d9fbb01c8", _deploy.deployments[0].ID)
	})
}

func TestDeployments_Rollback(t *testing.T) {

	// We create a Mock Server
	server := MockMarathonServer()
	defer server.Close()

	t.Run("get error with invalid Deployment id", func(t *testing.T) {

		// Try to create Deployment
		_deploy := NewDeployments(New(server.URL))

		// try to Get with empty app id
		err := _deploy.Rollback("")

		// We get error if deployment id don't exist
		assert.NotNil(t, err)
		assert.Equal(t, "deployment id cannot be null nor empty", err.Error())

		// Application ref is empty
		assert.Empty(t, _deploy.deployments)
	})

	t.Run("get valid Deployments if exists and then Rollback", func(t *testing.T) {

		deployArray = someDeployments

		// Try to create Deployment
		_deploy := NewDeployments(New(server.URL))

		// try to Get with empty app id
		err := _deploy.Rollback("97c136bf-5a28-4821-9d94-480d9fbb01c8")

		// err must be nil
		assert.Nil(t, err)

		// Application ref is empty
		assert.Empty(t, _deploy.deployments)

		// Check response of new deploy started by the rollback
		assert.Equal(t, "d4b75430-8ee6-47e9-95f2-6cf297aaac00", _deploy.deploy.ID)
	})
}

func TestDeployments_Await(t *testing.T) {

	// We create a Mock Server
	server := MockMarathonServer()
	defer server.Close()

	//
	var timeout time.Duration = 5 * time.Second
	deployArray = someDeployments

	t.Run("returns nil when Deploy don't found", func(t *testing.T) {

		// Try to create Deployment
		_deploy := NewDeployments(New(server.URL))

		// deploy Id to check
		id := "97c136bf-5a28-4821-9d94-480d9fbb01cX"

		// Fire up Await of deploy
		err := _deploy.Await(id, timeout)

		// err must be nil
		assert.Nil(t, err)
	})

	t.Run("returns err when Deploy still existing and timeout was reached", func(t *testing.T) {

		// Try to create Deployment
		_deploy := NewDeployments(New(server.URL))

		// deploy Id to check
		id := "97c136bf-5a28-4821-9d94-480d9fbb01c8"

		// Fire up Await of deploy
		err := _deploy.Await(id, timeout)

		// We get error if deploy still exist after timeout
		assert.NotNil(t, err)
		assert.Equal(t, "exit by timeout... deployment still existing", err.Error())
	})

	t.Run("return nil when Deploy exist and finish before timeout", func(t *testing.T) {

		// Try to create Deployment
		_deploy := NewDeployments(New(server.URL))

		// deploy Id to check
		id := "97c136bf-5a28-4821-9d94-480d9fbb01c8"

		// Fire up Await of deploy
		err := _deploy.Await(id, timeout)

		// err must be nil
		assert.Nil(t, err)

	})
}
