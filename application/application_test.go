package application

import (
	"encoding/json"
	"github.com/dotWicho/marathon"
	"github.com/dotWicho/marathon/mockserver"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"os"
	"testing"
)

func Test_NewApplication(t *testing.T) {

	t.Run("nil Application if send nil client", func(t *testing.T) {

		// Try to create Application
		_app := New(nil)

		// Application must be nil
		assert.Nil(t, _app)
	})

	t.Run("valid Application if send valid client", func(t *testing.T) {

		// Try to create Application
		_app := New(marathon.New("http://127.0.0.1:8080"))

		// Application must be not nil
		assert.NotNil(t, _app)
	})
}

func TestApplication_Get(t *testing.T) {

	// We create a Mock Server
	server := mockserver.MockServer()
	defer server.Close()

	t.Run("get empty Application ref if id is empty", func(t *testing.T) {

		// Try to create Application
		_app := New(marathon.New(server.URL))

		// try to Get with empty app id
		_refapp := _app.Get("")

		// Application is equal after Get fire up
		assert.Equal(t, _refapp, _app)

		// Application ref must be empty
		assert.Empty(t, _app.app.App)
	})

	t.Run("get valid Application ref if id is valid", func(t *testing.T) {

		// We define some vars
		appID := "/infra/redis-1"

		// Try to create Application
		_app := New(marathon.New(server.URL))

		// try to Get with empty app id
		_refapp := _app.Get(appID)

		// Application is equal after Get fire up
		assert.Equal(t, _refapp, _app)

		// Application ref must be not empty
		assert.NotEmpty(t, _app.app.App)

		// Check some values on response, must be equals
		assert.Equal(t, appID, _app.app.App.ID)
	})
}

func TestApplication_Set(t *testing.T) {

	// We create a Mock Server
	server := mockserver.MockServer()
	defer server.Close()

	t.Run("get empty Application ref if id is empty", func(t *testing.T) {

		// we define some vars
		emptyApp := AppDefinition{}

		// Try to create Application
		_app := New(marathon.New(server.URL))

		// try to Get with empty app id
		_refapp := _app.Set(emptyApp)

		// Application is equal after Get fire up
		assert.Equal(t, _refapp, _app)

		// Application ref must be empty
		assert.Empty(t, _app.app.App)
	})

	t.Run("get valid Application ref if app is valid", func(t *testing.T) {

		// we define some vars
		redisApp := &App{}
		_ = json.Unmarshal([]byte(mockserver.AppRedis), redisApp)

		// Try to create Application
		_app := New(marathon.New(server.URL))

		// try to Set with valid app
		_refapp := _app.Set(redisApp.App)

		// Application is equal after Get fire up
		assert.Equal(t, _refapp, _app)

		// Application ref must be not empty
		assert.NotEmpty(t, _app.app.App)

		// Check some values on response
		assert.Equal(t, redisApp.App.ID, _app.app.App.ID)
	})
}

func TestApplication_Create(t *testing.T) {

	// We create a Mock Server
	server := mockserver.MockServer()
	defer server.Close()

	t.Run("get empty Application ref if app is empty", func(t *testing.T) {

		// we define some vars
		emptyApp := AppDefinition{}
		// Try to create Application
		_app := New(marathon.New(server.URL))

		// try to Get with empty app id
		_refapp := _app.Create(emptyApp)

		// Application is equal after Get fire up
		assert.Equal(t, _refapp, _app)

		// Application ref must be empty
		assert.Empty(t, _app.app.App)
	})

	t.Run("get valid Application ref if app is valid", func(t *testing.T) {

		// we define some vars
		redisApp := &App{}
		_ = json.Unmarshal([]byte(mockserver.AppRedis), redisApp)

		// Try to create Application
		_app := New(marathon.New(server.URL))

		// try to Get with empty app id
		_refapp := _app.Create(redisApp.App)

		// Application is equal after Get fire up
		assert.Equal(t, _refapp, _app)

		// Application ref is not empty
		assert.NotEmpty(t, _app.app.App)

		// Our app ref must be not empty
		assert.NotEmpty(t, _app.app.App.ID)

		// Check some values on response
		assert.Equal(t, redisApp.App.ID, _app.app.App.ID)
	})
}

func TestApplication_Destroy(t *testing.T) {

	// We create a Mock Server
	server := mockserver.MockServer()
	defer server.Close()

	t.Run("get error if Application ref is empty", func(t *testing.T) {

		// Try to create Application
		_app := New(marathon.New(server.URL))

		// try to Get with empty app id
		err := _app.Destroy()

		// We get an error (app cannot be null nor empty)
		assert.NotNil(t, err)
		assert.Equal(t, "app cannot be null nor empty", err.Error())

		// Application ref must be empty
		assert.Empty(t, _app.app.App)
	})

	t.Run("erase Application if app ref is valid", func(t *testing.T) {

		// we define some vars
		redisApp := &App{}
		_ = json.Unmarshal([]byte(mockserver.AppRedis), redisApp)

		// Try to create Application
		_app := New(marathon.New(server.URL))

		// try to Set with valid app
		_ = _app.Set(redisApp.App)

		// try to Get with empty app id
		err := _app.Destroy()

		// We not get an error
		assert.Nil(t, err)

		// Our app ref must be not empty
		assert.Empty(t, _app.app.App)
	})
}

func TestApplication_Update(t *testing.T) {

	// We create a Mock Server
	server := mockserver.MockServer()
	defer server.Close()

	t.Run("get empty Application ref if app is empty", func(t *testing.T) {

		// we define some vars
		emptyApp := AppDefinition{}

		// Try to create Application
		_app := New(marathon.New(server.URL))

		// try to Get with empty app id
		err := _app.Update(emptyApp)

		// We get an error (app cannot be null nor empty)
		assert.NotNil(t, err)
		assert.Equal(t, "app cannot be null nor empty", err.Error())

		// Application ref is empty
		assert.Empty(t, _app.app.App)
	})

	t.Run("get valid Application ref if app is valid", func(t *testing.T) {

		// we define some vars
		redisApp := &App{}
		_ = json.Unmarshal([]byte(mockserver.AppRedis), redisApp)

		// Try to create Application
		_app := New(marathon.New(server.URL))

		// try to Get with empty app id
		err := _app.Update(redisApp.App)

		// We not get an error
		assert.Nil(t, err)

		// Application ref must be not empty
		assert.NotEmpty(t, _app.app.App)

		// Check some values on response
		assert.Equal(t, redisApp.App.ID, _app.app.App.ID)
	})
}

func TestApplication_Instances(t *testing.T) {

	// We create a Mock Server
	server := mockserver.MockServer()
	defer server.Close()

	t.Run("get -1 when Instances is called with app empty", func(t *testing.T) {

		// Try to create Application
		_app := New(marathon.New(server.URL))

		// try to Get with empty app id
		instances := _app.Instances()

		// Check results
		assert.Equal(t, -1, instances)

		// Application ref must be empty
		assert.Empty(t, _app.app.App)
	})

	t.Run("get valid Instances (>= 0) when is called with valid app", func(t *testing.T) {

		// we define some vars
		redisApp := &App{}
		_ = json.Unmarshal([]byte(mockserver.AppRedis), redisApp)

		// Try to create Application
		_app := New(marathon.New(server.URL))

		// try to Get with empty app id
		instances := _app.Get(redisApp.App.ID).Instances()

		// Application ref must be not empty
		assert.NotEmpty(t, _app.app.App)

		// Check some values on response
		assert.Equal(t, redisApp.App.Instances, instances)
	})
}

func TestApplication_Scale(t *testing.T) {

	// We create a Mock Server
	server := mockserver.MockServer()
	defer server.Close()

	t.Run("get error when Scale is called with app empty", func(t *testing.T) {

		// Try to create Application
		_app := New(marathon.New(server.URL))

		// try to Get with empty app id
		err := _app.Scale(2, true)

		// Error must be "app cannot be null nor empty"
		assert.NotNil(t, err)
		assert.Equal(t, "app cannot be null nor empty", err.Error())

		// Application ref must be empty
		assert.Empty(t, _app.app.App)
	})

	t.Run("get not error when Scale is called with a valid app", func(t *testing.T) {

		// we define some vars
		redisApp := &App{}
		_ = json.Unmarshal([]byte(mockserver.AppRedis), redisApp)

		// Try to create Application
		_app := New(marathon.New(server.URL))

		// try to Get with empty app id
		err := _app.Get(redisApp.App.ID).Scale(2, true)

		// We get not error
		assert.Nil(t, err)

		// Check some values on response
		assert.Equal(t, 2, _app.Instances())
	})
}

func TestApplication_Start(t *testing.T) {

	TestApplication_Scale(t)
}

func TestApplication_Stop(t *testing.T) {

	// We create a Mock Server
	server := mockserver.MockServer()
	defer server.Close()

	t.Run("get error when Stop is called with app empty", func(t *testing.T) {

		// Try to create Application
		_app := New(marathon.New(server.URL))

		// try to Get with empty app id
		err := _app.Stop(true)

		// Error must be "app cannot be null nor empty"
		assert.NotNil(t, err)
		assert.Equal(t, "app cannot be null nor empty", err.Error())

		// Application ref must be empty
		assert.Empty(t, _app.app.App)
	})

	t.Run("get not error when Stop is called with a valid app", func(t *testing.T) {

		// we define some vars
		redisApp := &App{}
		_ = json.Unmarshal([]byte(mockserver.AppRedis), redisApp)

		// Try to create Application
		_app := New(marathon.New(server.URL))

		// try to Get with empty app id
		err := _app.Get(redisApp.App.ID).Stop(true)

		// We get not error
		assert.Nil(t, err)

		// Check some values on response
		assert.Equal(t, 0, _app.Instances())
	})
}

func TestApplication_Restart(t *testing.T) {

	// We create a Mock Server
	server := mockserver.MockServer()
	defer server.Close()

	t.Run("get error when Restart is called with app empty", func(t *testing.T) {

		// Try to create Application
		_app := New(marathon.New(server.URL))

		// try to Get with empty app id
		err := _app.Restart(true)

		// Error must be "app cannot be null nor empty"
		assert.NotNil(t, err)
		assert.Equal(t, "app cannot be null nor empty", err.Error())

		// Application ref must be empty
		assert.Empty(t, _app.app.App)
	})

	t.Run("get not error when Restart is called with a valid app", func(t *testing.T) {

		// we define some vars
		redisApp := &App{}
		_ = json.Unmarshal([]byte(mockserver.AppRedis), redisApp)

		// Try to create Application
		_app := New(marathon.New(server.URL))

		// try to Get with empty app id
		err := _app.Get(redisApp.App.ID).Restart(true)

		// We get not error
		assert.Nil(t, err)

		// Check some values on response
		assert.Equal(t, 1, _app.Instances())
	})
}

func TestApplication_Suspend(t *testing.T) {

	// We create a Mock Server
	server := mockserver.MockServer()
	defer server.Close()

	t.Run("get error when Suspend is called with app empty", func(t *testing.T) {

		// Try to create Application
		_app := New(marathon.New(server.URL))

		// try to Get with empty app id
		err := _app.Suspend(true)

		// Error must be "app cannot be null nor empty"
		assert.NotNil(t, err)
		assert.Equal(t, "app cannot be null nor empty", err.Error())

		// Application ref must be empty
		assert.Empty(t, _app.app.App)
	})

	t.Run("get not error when Suspend is called with a valid app", func(t *testing.T) {

		// we define some vars
		redisApp := &App{}
		_ = json.Unmarshal([]byte(mockserver.AppRedis), redisApp)

		// Try to create Application
		_app := New(marathon.New(server.URL))

		// try to Get with empty app id
		err := _app.Get(redisApp.App.ID).Suspend(true)

		// We get not error
		assert.Nil(t, err)

		// Check some values on response
		assert.Equal(t, 0, _app.Instances())
	})
}

func TestApplication_GetTag(t *testing.T) {

	// We create a Mock Server
	server := mockserver.MockServer()
	defer server.Close()

	t.Run("get error when GetTag is called with app empty", func(t *testing.T) {

		// Try to create Application
		_app := New(marathon.New(server.URL))

		// try to Get with empty app id
		tag, err := _app.GetTag()

		// Error must be "app cannot be null nor empty"
		assert.NotNil(t, err)
		assert.Equal(t, "app cannot be null nor empty", err.Error())

		// Application ref must be empty
		assert.Empty(t, tag)
	})

	t.Run("get not error when GetTag is called with a valid app", func(t *testing.T) {

		// we define some vars
		redisApp := &App{}
		_ = json.Unmarshal([]byte(mockserver.AppRedis), redisApp)

		// Try to create Application
		_app := New(marathon.New(server.URL))

		// try to Get with empty app id
		tag, err := _app.Get(redisApp.App.ID).GetTag()

		// We get not error
		assert.Nil(t, err)

		// Check some values on response
		assert.Equal(t, "5.0.5", tag)
	})
}

func TestApplication_SetTag(t *testing.T) {

	// We create a Mock Server
	server := mockserver.MockServer()
	defer server.Close()

	t.Run("get error when SetTag is called with app empty", func(t *testing.T) {

		// Try to create Application
		_app := New(marathon.New(server.URL))

		// try to Get with empty app id
		err := _app.SetTag("4.0.10", true)

		// Error must be "app cannot be null nor empty"
		assert.NotNil(t, err)
		assert.Equal(t, "app cannot be null nor empty", err.Error())
	})

	t.Run("get not error when SetTag is called with a valid app", func(t *testing.T) {

		// we define some vars
		redisApp := &App{}
		_ = json.Unmarshal([]byte(mockserver.AppRedis), redisApp)

		// Try to create Application
		_app := New(marathon.New(server.URL))

		// try to Get with empty app id
		err := _app.Get(redisApp.App.ID).SetTag("4.0.10", true)

		// We get not error
		assert.Nil(t, err)

		// Get app tag value to compare
		tag, _ := _app.GetTag()

		// Check some values on response
		assert.Equal(t, "4.0.10", tag)
	})
}

func TestApplication_Env(t *testing.T) {

	// We create a Mock Server
	server := mockserver.MockServer()
	defer server.Close()

	t.Run("get nil when Env is called with app empty", func(t *testing.T) {

		// Try to create Application
		_app := New(marathon.New(server.URL))

		// try to Get with empty app id
		envs := _app.Env()

		// envs must be nil
		assert.Nil(t, envs)
	})

	t.Run("get Env when Env is called with a valid app", func(t *testing.T) {

		// we define some vars
		redisApp := &App{}
		_ = json.Unmarshal([]byte(mockserver.AppRedis), redisApp)

		// Try to create Application
		_app := New(marathon.New(server.URL))

		// try to Get with empty app id
		envs := _app.Get(redisApp.App.ID).Env()

		// Get values to compare
		_env := _app.Env()

		// Check some values on response
		assert.Equal(t, redisApp.App.Env, _env)
		assert.Equal(t, _env, envs)
	})
}

func TestApplication_SetEnv(t *testing.T) {

	// We create a Mock Server
	server := mockserver.MockServer()
	defer server.Close()

	t.Run("get nil when SetEnv is called with app empty", func(t *testing.T) {

		// Try to create Application
		_app := New(marathon.New(server.URL))

		// try to Get with empty app id
		err := _app.SetEnv("TESTED", "YES", true)

		// Error must be "app cannot be null nor empty"
		assert.NotNil(t, err)
		assert.Equal(t, "app cannot be null nor empty", err.Error())
	})

	t.Run("get Env when SetEnv is called with a valid app", func(t *testing.T) {

		// we define some vars
		redisApp := &App{}
		_ = json.Unmarshal([]byte(mockserver.AppRedis), redisApp)

		// Try to create Application
		_app := New(marathon.New(server.URL))

		// try to Get with empty app id
		err := _app.Get(redisApp.App.ID).SetEnv("TESTED", "YES", true)

		// We get not error
		assert.Nil(t, err)

		_env := _app.Env()

		// Check some values on response
		assert.Equal(t, "YES", _env["TESTED"])
	})
}

func TestApplication_DelEnv(t *testing.T) {

	// We create a Mock Server
	server := mockserver.MockServer()
	defer server.Close()

	t.Run("get nil when DelEnv is called with app empty", func(t *testing.T) {

		// Try to create Application
		_app := New(marathon.New(server.URL))

		// try to Get with empty app id
		err := _app.DelEnv("REDISPORT", true)

		// Error must be "app cannot be null nor empty"
		assert.NotNil(t, err)
		assert.Equal(t, "app cannot be null nor empty", err.Error())
	})

	t.Run("del Env is called with a valid app", func(t *testing.T) {

		// we define some vars
		redisApp := &App{}
		_ = json.Unmarshal([]byte(mockserver.AppRedis), redisApp)

		// Try to create Application
		_app := New(marathon.New(server.URL))

		// try to Get with empty app id
		_ = _app.Get(redisApp.App.ID).SetEnv("TESTED", "YES", true)
		err := _app.DelEnv("TESTED", true)

		// We get not error
		assert.Nil(t, err)

		_env := _app.Env()

		// Check some values on response
		assert.Equal(t, "", _env["TESTED"])
	})
}

func TestApplication_Cpus(t *testing.T) {

	// We create a Mock Server
	server := mockserver.MockServer()
	defer server.Close()

	t.Run("get -1 when Cpus is called with app empty", func(t *testing.T) {

		// Try to create Application
		_app := New(marathon.New(server.URL))

		// try to Get with empty app id
		cpus := _app.Cpus()

		// Number of cpus must be negative
		assert.Equal(t, float64(-1), cpus)
	})

	t.Run("get valid Cpus number is called with a valid app", func(t *testing.T) {

		// we define some vars
		redisApp := &App{}
		_ = json.Unmarshal([]byte(mockserver.AppRedis), redisApp)

		// Try to create Application
		_app := New(marathon.New(server.URL))

		// try to Get with empty app id
		cpus := _app.Get(redisApp.App.ID).Cpus()

		// Check some values on response
		assert.Equal(t, redisApp.App.Cpus, cpus)
	})
}

func TestApplication_SetCpus(t *testing.T) {

	// We create a Mock Server
	server := mockserver.MockServer()
	defer server.Close()

	t.Run("get err when SetCpus is called with app empty", func(t *testing.T) {

		// Try to create Application
		_app := New(marathon.New(server.URL))

		// try to Get with empty app id
		err := _app.SetCpus(float64(4), true)

		// Error must be "app cannot be null nor empty"
		assert.NotNil(t, err)
		assert.Equal(t, "app cannot be null nor empty", err.Error())
	})

	t.Run("get not err when SetCpus is called with a valid app", func(t *testing.T) {

		// we define some vars
		redisApp := &App{}
		_ = json.Unmarshal([]byte(mockserver.AppRedis), redisApp)

		// Try to create Application
		_app := New(marathon.New(server.URL))

		// try to Get with empty app id
		err := _app.Get(redisApp.App.ID).SetCpus(float64(4), true)

		// err must be nil
		assert.Nil(t, err)

		// Check some values on response
		assert.Equal(t, float64(4), _app.Cpus())
	})
}

func TestApplication_Memory(t *testing.T) {

	// We create a Mock Server
	server := mockserver.MockServer()
	defer server.Close()

	t.Run("get -1 when Memory is called with app empty", func(t *testing.T) {

		// Try to create Application
		_app := New(marathon.New(server.URL))

		// try to Get with empty app id
		memory := _app.Memory()

		// memory must be negative
		assert.Equal(t, float64(-1), memory)
	})

	t.Run("get valid Memory number is called with a valid app", func(t *testing.T) {

		// we define some vars
		redisApp := &App{}
		_ = json.Unmarshal([]byte(mockserver.AppRedis), redisApp)

		// Try to create Application
		_app := New(marathon.New(server.URL))

		// try to Get with empty app id
		memory := _app.Get(redisApp.App.ID).Memory()

		// Check some values on response
		assert.Equal(t, redisApp.App.Mem, memory)
	})
}

func TestApplication_SetMemory(t *testing.T) {

	// We create a Mock Server
	server := mockserver.MockServer()
	defer server.Close()

	t.Run("get err when SetMemory is called with app empty", func(t *testing.T) {

		// Try to create Application
		_app := New(marathon.New(server.URL))

		// try to Get with empty app id
		err := _app.SetMemory(float64(4096), true)

		// Error must be "app cannot be null nor empty"
		assert.NotNil(t, err)
		assert.Equal(t, "app cannot be null nor empty", err.Error())
	})

	t.Run("get not err when SetMemory is called with a valid app", func(t *testing.T) {

		// we define some vars
		redisApp := &App{}
		_ = json.Unmarshal([]byte(mockserver.AppRedis), redisApp)

		// Try to create Application
		_app := New(marathon.New(server.URL))

		// try to Get with empty app id
		err := _app.Get(redisApp.App.ID).SetMemory(float64(4096), true)

		// err must be nil
		assert.Nil(t, err)

		// Check some values on response
		assert.Equal(t, float64(4096), _app.Memory())
	})
}

func TestApplication_Role(t *testing.T) {

	// We create a Mock Server
	server := mockserver.MockServer()
	defer server.Close()

	t.Run("get empty when Role is called with app empty", func(t *testing.T) {

		// Try to create Application
		_app := New(marathon.New(server.URL))

		// try to Get with empty app id
		role := _app.Role()

		// role must be empty
		assert.Equal(t, "", role)
	})

	t.Run("get valid Role is called with a valid app", func(t *testing.T) {

		// we define some vars
		redisApp := &App{}
		_ = json.Unmarshal([]byte(mockserver.AppRedis), redisApp)

		// Try to create Application
		_app := New(marathon.New(server.URL))

		// try to Get with empty app id
		role := _app.Get(redisApp.App.ID).Role()

		// Check some values on response
		assert.Equal(t, redisApp.App.Role, role)
	})
}

func TestApplication_SetRole(t *testing.T) {

	// We create a Mock Server
	server := mockserver.MockServer()
	defer server.Close()

	t.Run("get err when SetRole is called with app empty", func(t *testing.T) {

		// Try to create Application
		_app := New(marathon.New(server.URL))

		// try to Get with empty app id
		err := _app.SetRole("slave_private", true)

		// Error must be "app cannot be null nor empty"
		assert.NotNil(t, err)
		assert.Equal(t, "app cannot be null nor empty", err.Error())
	})

	t.Run("get not err when SetRole is called with a valid app", func(t *testing.T) {

		// we define some vars
		redisApp := &App{}
		_ = json.Unmarshal([]byte(mockserver.AppRedis), redisApp)

		// Try to create Application
		_app := New(marathon.New(server.URL))

		// try to Get with empty app id
		err := _app.Get(redisApp.App.ID).SetRole("slave_private", true)

		// err must be nil
		assert.Nil(t, err)

		// Check some values on response
		assert.Equal(t, "slave_private", _app.Role())
	})
}

func TestApplication_Container(t *testing.T) {

	// We create a Mock Server
	server := mockserver.MockServer()
	defer server.Close()

	t.Run("get nil when Parameters is called with app empty", func(t *testing.T) {

		// Try to create Application
		_app := New(marathon.New(server.URL))

		// try to Get with empty app id
		containerRef := _app.Container()

		// container must be nil
		assert.Nil(t, containerRef)
	})

	t.Run("get Params when Parameters is called with a valid app", func(t *testing.T) {

		// we define some vars
		redisApp := &App{}
		_ = json.Unmarshal([]byte(mockserver.AppRedis), redisApp)

		// Try to create Application
		_app := New(marathon.New(server.URL))

		// try to Get with empty app id
		containerRef := _app.Get(redisApp.App.ID).Container()

		// We get not error
		assert.NotNil(t, containerRef)

		// Check some values on response
		assert.Equal(t, "DOCKER", containerRef.Type)
		assert.Equal(t, "docker.io/redis-ha:5.0.5", containerRef.Docker.Image)
	})
}

func TestApplication_SetContainer(t *testing.T) {

	// We create a Mock Server
	server := mockserver.MockServer()
	defer server.Close()

	t.Run("get nil when Parameters is called with app empty", func(t *testing.T) {

		// We define some vars
		container := &marathon.Container{
			Type: "DOCKER",
			Docker: marathon.Docker{
				Image:          "fake.registry.org/redis:5.0.6",
				Network:        "",
				Privileged:     false,
				Parameters:     nil,
				ForcePullImage: true,
			},
			Volumes:      nil,
			PortMappings: nil,
		}

		// Try to create Application
		_app := New(marathon.New(server.URL))

		// try to Get with empty app id
		err := _app.SetContainer(container, true)

		// Error must be "app cannot be null nor empty"
		assert.NotNil(t, err)
		assert.Equal(t, "app cannot be null nor empty", err.Error())
	})

	t.Run("get Params when Parameters is called with a valid app", func(t *testing.T) {

		// we define some vars
		container := &marathon.Container{
			Type: "DOCKER",
			Docker: marathon.Docker{
				Image:          "fake.registry.org/redis:5.0.6",
				Network:        "",
				Privileged:     false,
				Parameters:     nil,
				ForcePullImage: true,
			},
			Volumes:      nil,
			PortMappings: nil,
		}
		redisApp := &App{}
		_ = json.Unmarshal([]byte(mockserver.AppRedis), redisApp)

		// Try to create Application
		_app := New(marathon.New(server.URL))

		// try to Get with empty app id
		err := _app.Get(redisApp.App.ID).SetContainer(container, true)

		// We get not error
		assert.Nil(t, err)

		// Check some values on response
		assert.Equal(t, container, _app.Container())
	})
}

func TestApplication_Parameter(t *testing.T) {

	// We create a Mock Server
	server := mockserver.MockServer()
	defer server.Close()

	t.Run("get nil when Parameters is called with app empty", func(t *testing.T) {

		// Try to create Application
		_app := New(marathon.New(server.URL))

		// try to Get with empty app id
		params, err := _app.Parameters()

		// err must be nil
		assert.NotNil(t, err)

		// params must be empty
		assert.Nil(t, params)
	})

	t.Run("get nil Params when Parameters is called with a valid app but that app dont have Params", func(t *testing.T) {

		// we define some vars
		kongAppRef := &App{}
		_ = json.Unmarshal([]byte(mockserver.KongApp), kongAppRef)

		// Try to create Application
		_app := New(marathon.New(server.URL))

		// try to Get with empty app id
		params, err := _app.Get(kongAppRef.App.ID).Parameters()

		// We get not error
		assert.NotNil(t, err)
		assert.Equal(t, "the Marathon app /infra/kong-v2 has no Docker parameters", err.Error())
		assert.Nil(t, params)
	})

	t.Run("get Params when Parameters is called with a valid app", func(t *testing.T) {

		// we define some vars
		redisApp := &App{}
		_ = json.Unmarshal([]byte(mockserver.AppRedis), redisApp)

		// Try to create Application
		_app := New(marathon.New(server.URL))

		// try to Get with empty app id
		params, err := _app.Get(redisApp.App.ID).Parameters()

		// We get not error
		assert.Nil(t, err)

		// Check some values on response
		assert.Equal(t, map[string]string{"add-host": "10.128.64.32"}, params)
	})
}

func TestApplication_AddParameter(t *testing.T) {

	// We create a Mock Server
	server := mockserver.MockServer()
	defer server.Close()

	t.Run("get nil when AppParameters is called with app empty", func(t *testing.T) {

		// Try to create Application
		_app := New(marathon.New(server.URL))

		// try to Get with empty app id
		err := _app.AddParameter("TESTED", "YES", true)

		// Error must be "app cannot be null nor empty"
		assert.NotNil(t, err)
		assert.Equal(t, "app cannot be null nor empty", err.Error())
	})

	t.Run("get Parameters when AppParameters is called with a valid app", func(t *testing.T) {

		// we define some vars
		redisApp := &App{}
		_ = json.Unmarshal([]byte(mockserver.AppRedis), redisApp)

		// Try to create Application
		_app := New(marathon.New(server.URL))

		// try to Get with empty app id
		err := _app.Get(redisApp.App.ID).AddParameter("new-host", "10.64.128.1", true)

		// We get not error
		assert.Nil(t, err)

		_map, _ := _app.Parameters()

		// Check some values on response
		assert.Equal(t, "10.128.64.32", _map["add-host"])
		assert.Equal(t, "10.64.128.1", _map["new-host"])
	})
}

func TestApplication_DelParameter(t *testing.T) {

	// We create a Mock Server
	server := mockserver.MockServer()
	defer server.Close()

	t.Run("get nil when DelParameter is called with app empty", func(t *testing.T) {

		// Try to create Application
		_app := New(marathon.New(server.URL))

		// try to Get with empty app id
		err := _app.DelParameter("add-host", true)

		// Error must be "app cannot be null nor empty"
		assert.NotNil(t, err)
		assert.Equal(t, "app cannot be null nor empty", err.Error())
	})

	t.Run("get error when DelParameter is called with an invalid Parameter", func(t *testing.T) {

		// we define some vars
		redisApp := &App{}
		_ = json.Unmarshal([]byte(mockserver.AppRedis), redisApp)

		// Try to create Application
		_app := New(marathon.New(server.URL))

		// try to Get with empty app id
		err := _app.Get(redisApp.App.ID).DelParameter("add-folder", true)

		// Error must be "parameters add-folder dont exist in Marathon app /infra/redis-1"
		assert.NotNil(t, err)
		assert.Equal(t, "parameters add-folder dont exist in Marathon app /infra/redis-1", err.Error())
	})

	t.Run("get not error when DelParameter is called with a valid Parameter", func(t *testing.T) {

		// we define some vars
		redisApp := &App{}
		_ = json.Unmarshal([]byte(mockserver.AppRedis), redisApp)

		// Try to create Application
		_app := New(marathon.New(server.URL))

		// try to Get with empty app id
		err := _app.Get(redisApp.App.ID).DelParameter("add-host", true)

		// Error must be nil
		assert.Nil(t, err)
	})
}

func TestApplication_Load(t *testing.T) {

	// We create a Mock Server
	server := mockserver.MockServer()
	defer server.Close()

	t.Run("get nil when Load is called with invalid file", func(t *testing.T) {

		// We define some vars
		fileName := "dumpfile.txt"

		// Try to create Application
		_app := New(marathon.New(server.URL))

		// try to Get with empty app id
		_app = _app.Load(fileName)
		defer os.Remove(fileName)

		// Apps data reference must be empty
		assert.Empty(t, _app.app.App)
	})

	t.Run("get App ref when is called with a valid file", func(t *testing.T) {

		// We define some vars
		fileName := "dumpfile.json"
		redisApp := &App{}
		_ = json.Unmarshal([]byte(mockserver.AppRedis), redisApp)

		// We need create a []byte with redisApp.App because Dump writes that
		redisRef, _ := json.MarshalIndent(redisApp.App, "", "  ")

		// We create out file to read as JSON
		errFile := ioutil.WriteFile(fileName, redisRef, 0644)

		// We get not error
		assert.Nil(t, errFile)

		// Try to create Application
		_app := New(marathon.New(server.URL))

		// try to Get with empty app id
		_ = _app.Load(fileName)
		defer os.Remove(fileName)

		// Check some values on response
		assert.Equal(t, redisApp.App, _app.app.App)
	})
}

func TestApplication_Dump(t *testing.T) {

	// We create a Mock Server
	server := mockserver.MockServer()
	defer server.Close()

	t.Run("get nil when Dump is called with app empty", func(t *testing.T) {

		// We define some vars
		fileName := "dumpfile.json"

		// Try to create Application
		_app := New(marathon.New(server.URL))

		// try to Get with empty app id
		err := _app.Dump(fileName)
		defer os.Remove(fileName)

		// Error must be "app cannot be null nor empty"
		assert.NotNil(t, err)
		assert.Equal(t, "app cannot be null nor empty", err.Error())
	})

	t.Run("dump App content is called with a valid app", func(t *testing.T) {

		// We define some vars
		fileName := "dumpfile.json"
		redisApp := &App{}
		_ = json.Unmarshal([]byte(mockserver.AppRedis), redisApp)

		// We need create a []byte with redisApp.App because Dump writes that
		redisRef, _ := json.MarshalIndent(redisApp.App, "", "  ")

		// Try to create Application
		_app := New(marathon.New(server.URL))

		// try to Get with empty app id
		err := _app.Get(redisApp.App.ID).Dump(fileName)
		defer os.Remove(fileName)

		// We get not error
		assert.Nil(t, err)

		// Read content of file
		file, _ := ioutil.ReadFile(fileName)

		// Check some values on response
		assert.Equal(t, redisRef, file)
	})
}
