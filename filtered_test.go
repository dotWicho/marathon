package marathon

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func Test_NewFilteredApps(t *testing.T) {

	t.Run("nil FilteredApps if send nil client", func(t *testing.T) {

		// Try to create FilteredApps
		_apps := NewFilteredApps(nil)

		// FilteredApps is nil
		assert.Nil(t, _apps)
	})

	t.Run("valid Application if send valid client", func(t *testing.T) {

		// Try to create FilteredApps
		_apps := NewFilteredApps(New("http://127.0.0.1:8080"))

		// FilteredApps is not nil
		assert.NotNil(t, _apps)
	})
}

func TestFilteredApps_Get(t *testing.T) {

	// We create a Mock Server
	server := MockMarathonServer()
	defer server.Close()

	t.Run("get empty FilteredApps ref if filter is empty", func(t *testing.T) {

		// Try to create FilteredApp
		_apps := NewFilteredApps(New(server.URL))

		// try to Get with empty filter
		_refapps := _apps.Get("")

		// Application is equal after Get fire up
		assert.Equal(t, _refapps, _apps)

		// FilteredApps Apps ref is empty
		assert.Empty(t, _apps.apps.Apps)
	})

	t.Run("get valid FilteredApps ref if filter is valid", func(t *testing.T) {

		// We define some vars
		filter := "/infra"

		// Try to create FilteredApp
		_apps := NewFilteredApps(New(server.URL))

		// try to Get with empty app id
		_refapps := _apps.Get(filter)

		// FilteredApps is equal after Get fire up
		assert.Equal(t, _refapps, _apps)

		// Application ref is not empty
		assert.NotEmpty(t, _apps.apps.Apps)

		// Check some values on response
		assert.Equal(t, true, strings.HasPrefix(_apps.apps.Apps[0].ID, filter))
	})
}

func TestFilteredApps_Scale(t *testing.T) {

	// We create a Mock Server
	server := MockMarathonServer()
	defer server.Close()

	t.Run("get error when Scale is called with app empty", func(t *testing.T) {

		// Try to create FilteredApp
		_apps := NewFilteredApps(New(server.URL))

		// try to Get with empty app id
		err := _apps.Scale(2, true)

		//
		assert.NotNil(t, err)
		assert.Equal(t, "filteredApps Scale was called with an empty set", err.Error())
	})

	t.Run("get not error when Scale is called with a valid app", func(t *testing.T) {

		// We define some vars
		filter := "/infra"

		// Try to create FilteredApp
		_apps := NewFilteredApps(New(server.URL))

		// try to Get with empty app id
		err := _apps.Get(filter).Scale(2, true)

		// We get not error
		assert.Nil(t, err)

		// Check some values on response
		// assert.Equal(t, 2, _apps.apps.Apps[0].Instances)
	})
}

func TestFilteredApps_Stop(t *testing.T) {

	// We create a Mock Server
	server := MockMarathonServer()
	defer server.Close()

	t.Run("get error when Stop is called with app empty", func(t *testing.T) {

		// Try to create FilteredApp
		_apps := NewFilteredApps(New(server.URL))

		// try to Get with empty app id
		err := _apps.Stop(true)

		//
		assert.NotNil(t, err)
		assert.Equal(t, "filteredApps Stop was called with an empty set", err.Error())
	})

	t.Run("get not error when Stop is called with a valid app", func(t *testing.T) {

		// We define some vars
		filter := "/infra"

		// Try to create FilteredApp
		_apps := NewFilteredApps(New(server.URL))

		// try to Get with empty app id
		err := _apps.Get(filter).Stop(true)

		// We get not error
		assert.Nil(t, err)
	})
}

func TestFilteredApps_Start(t *testing.T) {

	// We create a Mock Server
	server := MockMarathonServer()
	defer server.Close()

	t.Run("get error when Start is called with app empty", func(t *testing.T) {

		// Try to create FilteredApp
		_apps := NewFilteredApps(New(server.URL))

		// try to Get with empty app id
		err := _apps.Start(2, true)

		// We get an error
		assert.NotNil(t, err)
		assert.Equal(t, "filteredApps Start was called with an empty set", err.Error())
	})

	t.Run("get not error when Start is called with a valid app", func(t *testing.T) {

		// We define some vars
		filter := "/infra"

		// Try to create FilteredApp
		_apps := NewFilteredApps(New(server.URL))

		// try to Get with empty app id
		err := _apps.Get(filter).Start(2, true)

		// We get not error
		assert.Nil(t, err)

		// Check some values on response
		assert.Equal(t, 2, _apps.apps.Apps[0].Instances)
		assert.Equal(t, 2, _apps.apps.Apps[1].Instances)
	})
}

func TestFilteredApps_Restart(t *testing.T) {

	// We create a Mock Server
	server := MockMarathonServer()
	defer server.Close()

	t.Run("get error when Restart is called with app empty", func(t *testing.T) {

		// Try to create FilteredApp
		_apps := NewFilteredApps(New(server.URL))

		// try to Get with empty app id
		err := _apps.Restart(true)

		// We get an error
		assert.NotNil(t, err)
		assert.Equal(t, "filteredApps Restart was called with an empty set", err.Error())
	})

	t.Run("get not error when Restart is called with a valid app", func(t *testing.T) {

		// We define some vars
		filter := "/infra"

		// Try to create FilteredApp
		_apps := NewFilteredApps(New(server.URL))

		// try to Get with empty app id
		err := _apps.Get(filter).Restart(true)

		// We get not error
		assert.Nil(t, err)
	})
}

func TestFilteredApps_Suspend(t *testing.T) {

	// We create a Mock Server
	server := MockMarathonServer()
	defer server.Close()

	t.Run("get error when Suspend is called with app empty", func(t *testing.T) {

		// Try to create FilteredApp
		_apps := NewFilteredApps(New(server.URL))

		// try to Get with empty app id
		err := _apps.Suspend(true)

		// We get an error
		assert.NotNil(t, err)
		assert.Equal(t, "filteredApps Stop was called with an empty set", err.Error())
	})

	t.Run("get not error when Suspend is called with a valid app", func(t *testing.T) {

		// We define some vars
		filter := "/infra"

		// Try to create FilteredApp
		_apps := NewFilteredApps(New(server.URL))

		// try to Get with empty app id
		err := _apps.Get(filter).Suspend(true)

		// We get not error
		assert.Nil(t, err)

		// Check some values on response
		assert.Equal(t, 0, _apps.apps.Apps[0].Instances)
		assert.Equal(t, 0, _apps.apps.Apps[1].Instances)
	})
}

func TestFilteredApps_Load(t *testing.T) {

	// We create a Mock Server
	server := MockMarathonServer()
	defer server.Close()

	t.Run("get nil when Load is called with invalid file", func(t *testing.T) {

		// We define some vars
		fileName := "dumpfile.txt"

		// Try to create FilteredApp
		_apps := NewFilteredApps(New(server.URL))

		// try to Get with empty app id
		_apps = _apps.Load(fileName, "")
		defer os.Remove(fileName)

		// We get an empty apps ref
		assert.NotNil(t, _apps)
		assert.Empty(t, _apps.apps.Apps)
	})

	t.Run("get App ref when is called with a valid JSON file", func(t *testing.T) {

		// We define some vars
		filter := "/infra"
		fileName := "dumpfile.json"
		appsFiltered := &apps{}
		_ = json.Unmarshal([]byte(appsArray), appsFiltered)

		// We create out file to read as JSON
		errFile := ioutil.WriteFile(fileName, []byte(appsArray), 0644)
		defer os.Remove(fileName)

		// We get not error
		assert.Nil(t, errFile)

		// Try to create FilteredApp
		_apps := NewFilteredApps(New(server.URL))

		// try to Get with empty app id
		_ = _apps.Load(fileName, filter)

		// Check some values on response
		assert.Equal(t, appsFiltered.Apps, _apps.apps.Apps)
	})
}

func TestFilteredApps_Dump(t *testing.T) {

	// We create a Mock Server
	server := MockMarathonServer()
	defer server.Close()

	t.Run("get nil when Dump is called with empty filter", func(t *testing.T) {

		// We define some vars
		fileName := "dumpfile.json"

		// Try to create FilteredApp
		_apps := NewFilteredApps(New(server.URL))

		// try to Get with empty app id
		err := _apps.Dump(fileName)
		defer os.Remove(fileName)

		// We get an error
		assert.NotNil(t, err)
		assert.Equal(t, "filteredApps Dump was called with an empty set", err.Error())
	})

	t.Run("dump App content when Dump is called with a valid filter", func(t *testing.T) {

		// We define some vars
		filter := "/infra"
		fileName := "dumpfile.json"
		appsFiltered := &apps{}
		_ = json.Unmarshal([]byte(appsArray), appsFiltered)

		// We need create a []byte with appsFiltered.Apps because Dump writes that
		appsRef, _ := json.MarshalIndent(appsFiltered.Apps, "", "  ")

		// Try to create FilteredApp
		_apps := NewFilteredApps(New(server.URL))

		// try to Get with empty app id
		err := _apps.Get(filter).Dump(fileName)
		defer os.Remove(fileName)

		// We get not error
		assert.Nil(t, err)

		// Read content of file
		file, _ := ioutil.ReadFile(fileName)

		// Check some values on response
		assert.Equal(t, appsRef, file)
	})
}

func TestFilteredApps_DumpSingly(t *testing.T) {

	// We create a Mock Server
	server := MockMarathonServer()
	defer server.Close()

	t.Run("get nil when DumpSingly is called with empty filter", func(t *testing.T) {

		// We define some vars
		baseName := "dumpfile.json"

		// Try to create FilteredApp
		_apps := NewFilteredApps(New(server.URL))

		// try to Get with empty app id
		err := _apps.DumpSingly(baseName)
		defer os.Remove(baseName + "*")

		// We get an error
		assert.NotNil(t, err)
		assert.Equal(t, "filteredApps Dump was called with an empty set", err.Error())
	})

	t.Run("get err when DumpSingly is called with invalid baseName", func(t *testing.T) {

		// We define some vars
		filter := "/infra"
		baseName := "dumpfile"

		// Try to create FilteredApp
		_apps := NewFilteredApps(New(server.URL))

		// try to Get with empty app id
		err := _apps.Get(filter).DumpSingly(baseName)

		// Error must be "invalid filename extension"
		assert.NotNil(t, err)
		assert.Equal(t, "invalid filename extension", err.Error())
	})

	t.Run("dump App content when DumpSingly is called with a valid filter", func(t *testing.T) {

		// We define some vars
		filter := "/infra"
		baseName := "dumpfile.json"
		appsFiltered := &apps{}
		_ = json.Unmarshal([]byte(appsArray), appsFiltered)

		// We need create a []byte with appsFiltered.Apps because Dump writes that
		appsRef, _ := json.MarshalIndent(appsFiltered.Apps[0], "", "  ")

		// Try to create FilteredApp
		_apps := NewFilteredApps(New(server.URL))

		// try to Get with empty app id
		err := _apps.Get(filter).DumpSingly(baseName)
		baseName = strings.TrimSuffix(baseName, filepath.Ext(baseName))

		defer os.Remove(baseName + "*")

		// We get not error
		assert.Nil(t, err)

		// Read content of file
		file, _ := ioutil.ReadFile(baseName + "-infra-redis-1.json")

		// Check some values on response
		assert.Equal(t, appsRef, file)
	})
}

func TestFilteredApps_FilterBy(t *testing.T) {

	// We create a Mock Server
	server := MockMarathonServer()
	defer server.Close()

	t.Run("get empty Apps when FilterBy is called with not match func", func(t *testing.T) {

		// Try to create FilteredApp
		_apps := NewFilteredApps(New(server.URL))

		// Define a func with always return false
		filterFunc := func(app AppDefinition) bool {

			return false
		}

		// try to Get with empty app id
		_filtered := _apps.FilterBy(filterFunc)

		// Apps reference nust be empty
		assert.Empty(t, _filtered.apps)
	})

	t.Run("get Apps when FilterBy is called with match func", func(t *testing.T) {

		// We define some vars
		filter := "/infra"

		// Define a func with return true when math ENVIRONMENT["BROKERPORT"] == 9092
		filterFunc := func(app AppDefinition) bool {

			if app.Env["BROKERPORT"] == "9092" {
				return true
			}
			return false
		}

		// Try to create FilteredApp
		_apps := NewFilteredApps(New(server.URL))

		// try to Get with empty app id
		_filtered := _apps.Get(filter).FilterBy(filterFunc)

		// Apps reference must be not empty
		assert.NotEmpty(t, _filtered.apps)
		assert.Equal(t, "/infra/broker-0", _filtered.apps.Apps[0].ID)
	})
}

func TestFilteredApps_AsMap(t *testing.T) {

	// We create a Mock Server
	server := MockMarathonServer()
	defer server.Close()

	t.Run("get empty Map if filter is empty", func(t *testing.T) {

		// Try to create FilteredApp
		_apps := NewFilteredApps(New(server.URL))

		// try to Get with empty filter
		_map := _apps.Get("").AsMap()

		// Application is equal after Get fire up
		assert.Empty(t, _map)
	})

	t.Run("get valid Map ref if filter is valid", func(t *testing.T) {

		// We define some vars
		filter := "/infra"

		// Try to create FilteredApp
		_apps := NewFilteredApps(New(server.URL))

		// try to Get with empty app id
		_map := _apps.Get(filter).AsMap()

		// FilteredApps is equal after Get fire up
		assert.Equal(t, len(_map), len(_apps.apps.Apps))

		// Application ref must be not empty
		assert.Equal(t, "/infra/redis-1", _map["/infra/redis-1"].ID)
	})
}

func TestFilteredApps_AsRaw(t *testing.T) {

	// We create a Mock Server
	server := MockMarathonServer()
	defer server.Close()

	t.Run("get empty Ref if filter is empty", func(t *testing.T) {

		// Try to create FilteredApp
		_apps := NewFilteredApps(New(server.URL))

		// try to Get with empty filter
		_raw := _apps.Get("").AsRaw()

		// Application is equal after Get fire up
		assert.Empty(t, _raw)
	})

	t.Run("get valid Ref ref if filter is valid", func(t *testing.T) {

		// We define some vars
		filter := "/infra"

		// Try to create FilteredApp
		_apps := NewFilteredApps(New(server.URL))

		// try to Get with empty app id
		_raw := _apps.Get(filter).AsRaw()

		// FilteredApps is equal after Get fire up
		assert.Equal(t, len(_raw), len(_apps.apps.Apps))

		// Application ref must be not empty
		assert.Equal(t, _raw[0].ID, _apps.apps.Apps[0].ID)
	})
}
