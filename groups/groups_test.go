package groups

import (
	"encoding/json"
	"github.com/dotWicho/marathon"
	"github.com/dotWicho/marathon/mockserver"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"os"
	"testing"
)

func Test_NewGroups(t *testing.T) {

	t.Run("nil Groups if send nil client", func(t *testing.T) {

		// Try to create Groups
		_groups := NewGroups(nil)

		// Groups must be nil
		assert.Nil(t, _groups)
	})

	t.Run("valid Groups if send valid client", func(t *testing.T) {

		// Try to create Groups
		_groups := NewGroups(marathon.New("http://127.0.0.1:8080"))

		// Groups must be not nil
		assert.NotNil(t, _groups)
	})
}

func TestGroups_Get(t *testing.T) {

	// We create a Mock Server
	server := mockserver.MockServer()
	defer server.Close()

	t.Run("get empty Group ref if id is empty", func(t *testing.T) {

		// Try to create Groups
		_group := NewGroups(marathon.New(server.URL))

		// try to Get with empty app id
		_groupRef := _group.Get("")

		// Initial group and reference by Get call must be equals
		assert.Equal(t, _groupRef, _group)

		// Group ref is empty
		assert.Empty(t, _group.group)
	})

	t.Run("get valid Group ref if id is valid", func(t *testing.T) {

		// We define some vars
		groupID := "/infra"

		// Try to create Groups
		_group := NewGroups(marathon.New(server.URL))

		// try to Get with empty app id
		_groupRef := _group.Get(groupID)

		// Check some values on response
		assert.Equal(t, _groupRef, _group)
		assert.Equal(t, groupID, _group.group.ID)
	})
}

func TestGroups_Create(t *testing.T) {

	// We create a Mock Server
	server := mockserver.MockServer()
	defer server.Close()

	t.Run("get empty Group ref if app is empty", func(t *testing.T) {

		// we define some vars
		emptyGroup := &Group{}

		// Try to create Groups
		_group := NewGroups(marathon.New(server.URL))

		// try to Get with empty app id
		err := _group.Create(emptyGroup)

		// Group is equal after Get fire up
		assert.NotNil(t, err)
		assert.Equal(t, "group cannot be null nor empty", err.Error())

		// Group ref is empty
		assert.Empty(t, _group.group)
	})

	t.Run("get valid Group ref if app is valid", func(t *testing.T) {

		// we define some vars
		validGroup := &Group{}
		_ = json.Unmarshal([]byte(mockserver.GroupsArray), validGroup)

		// Try to create Groups
		_group := NewGroups(marathon.New(server.URL))

		// try to Get with empty app id
		err := _group.Create(validGroup)

		// Group is equal after Get fire up
		assert.Nil(t, err)

		// Group ref must be not empty
		assert.NotEmpty(t, _group.group)

		// Check some values on response
		assert.Equal(t, validGroup.ID, _group.group.ID)
	})
}

func TestGroups_Destroy(t *testing.T) {

	// We create a Mock Server
	server := mockserver.MockServer()
	defer server.Close()

	t.Run("get empty Group ref if app is empty", func(t *testing.T) {

		// Try to create Groups
		_group := NewGroups(marathon.New(server.URL))

		// try to Get with empty app id
		err := _group.Destroy()

		// Group is equal after Get fire up
		assert.NotNil(t, err)
		assert.Equal(t, "group cannot be null nor empty", err.Error())

		// Group ref is empty
		assert.Empty(t, _group.group)
	})

	t.Run("get valid Group ref if app is valid", func(t *testing.T) {

		// we define some vars
		groupID := "/infra"

		// Try to create Groups
		_group := NewGroups(marathon.New(server.URL))

		// try to Get with empty app id
		err := _group.Get(groupID).Destroy()

		// Group is equal after Get fire up
		assert.Nil(t, err)

		// Check some values on response
		assert.Empty(t, _group.group)
	})
}

func TestGroups_Update(t *testing.T) {

	// We create a Mock Server
	server := mockserver.MockServer()
	defer server.Close()

	t.Run("get empty Group ref if app is empty", func(t *testing.T) {

		// we define some vars
		emptyGroup := &Group{}

		// Try to create Groups
		_group := NewGroups(marathon.New(server.URL))

		// try to Get with empty app id
		err := _group.Update(emptyGroup)

		// Group is equal after Get fire up
		assert.NotNil(t, err)
		assert.Equal(t, "group cannot be null nor empty", err.Error())

		// Group ref is empty
		assert.Empty(t, _group.group)
	})

	t.Run("get valid Group ref if app is valid", func(t *testing.T) {

		// we define some vars
		validGroup := &Group{}
		_ = json.Unmarshal([]byte(mockserver.GroupsArray), validGroup)

		// Try to create Groups
		_group := NewGroups(marathon.New(server.URL))

		// try to Get with empty app id
		err := _group.Update(validGroup)

		// Group is equal after Get fire up
		assert.Nil(t, err)

		// Group ref must be not empty
		assert.NotEmpty(t, _group.group)

		// Check some values on response
		assert.Equal(t, validGroup.ID, _group.group.ID)
	})
}

func TestGroups_Scale(t *testing.T) {

	// We create a Mock Server
	server := mockserver.MockServer()
	defer server.Close()

	t.Run("get empty Group ref if app is empty", func(t *testing.T) {

		// Try to create Groups
		_group := NewGroups(marathon.New(server.URL))

		// try to Get with empty app id
		err := _group.Scale(2, true)

		// Group is equal after Get fire up
		assert.NotNil(t, err)
		assert.Equal(t, "group cannot be null nor empty", err.Error())

		// Group ref is empty
		assert.Empty(t, _group.group)
	})

	t.Run("get valid Group ref if app is valid", func(t *testing.T) {

		// we define some vars
		validGroup := &Group{}
		_ = json.Unmarshal([]byte(mockserver.GroupsArray), validGroup)

		// Try to create Groups
		_group := NewGroups(marathon.New(server.URL))

		// try to Get with empty app id
		err := _group.Get(validGroup.ID).Scale(2, true)

		// Group is equal after Get fire up
		assert.Nil(t, err)

		// Group ref must be not empty
		assert.NotEmpty(t, _group.group)

		// Check some values on response
		assert.Equal(t, http.StatusOK, _group.client.StatusCode())
	})
}

func TestGroups_Stop(t *testing.T) {

	// We create a Mock Server
	server := mockserver.MockServer()
	defer server.Close()

	t.Run("get empty Group ref if app is empty", func(t *testing.T) {

		// Try to create Groups
		_group := NewGroups(marathon.New(server.URL))

		// try to Get with empty app id
		err := _group.Stop(true)

		// Group is equal after Get fire up
		assert.NotNil(t, err)
		assert.Equal(t, "group cannot be null nor empty", err.Error())

		// Group ref is empty
		assert.Empty(t, _group.group)
	})

	t.Run("get valid Group ref if app is valid", func(t *testing.T) {

		// we define some vars
		validGroup := &Group{}
		_ = json.Unmarshal([]byte(mockserver.GroupsArray), validGroup)

		// Try to create Groups
		_group := NewGroups(marathon.New(server.URL))

		// try to Get with empty app id
		err := _group.Get(validGroup.ID).Stop(true)

		// Group is equal after Get fire up
		assert.Nil(t, err)

		// Group ref must be not empty
		assert.NotEmpty(t, _group.group)

		// Check some values on response
		assert.Equal(t, http.StatusOK, _group.client.StatusCode())
	})
}

func TestGroups_Start(t *testing.T) {

	// We create a Mock Server
	server := mockserver.MockServer()
	defer server.Close()

	t.Run("get empty Group ref if app is empty", func(t *testing.T) {

		// Try to create Groups
		_group := NewGroups(marathon.New(server.URL))

		// try to Get with empty app id
		err := _group.Start(2, true)

		// Group is equal after Get fire up
		assert.NotNil(t, err)
		assert.Equal(t, "group cannot be null nor empty", err.Error())

		// Group ref is empty
		assert.Empty(t, _group.group)
	})

	t.Run("get valid Group ref if app is valid", func(t *testing.T) {

		// we define some vars
		validGroup := &Group{}
		_ = json.Unmarshal([]byte(mockserver.GroupsArray), validGroup)

		// Try to create Groups
		_group := NewGroups(marathon.New(server.URL))

		// try to Get with empty app id
		err := _group.Get(validGroup.ID).Start(2, true)

		// Group is equal after Get fire up
		assert.Nil(t, err)

		// Group ref must be not empty
		assert.NotEmpty(t, _group.group)

		// Check some values on response
		assert.Equal(t, http.StatusOK, _group.client.StatusCode())
	})
}

func TestGroups_Restart(t *testing.T) {

	// We create a Mock Server
	server := mockserver.MockServer()
	defer server.Close()

	t.Run("get empty Group ref if app is empty", func(t *testing.T) {

		// Try to create Groups
		_group := NewGroups(marathon.New(server.URL))

		// try to Get with empty app id
		err := _group.Restart(true)

		// Group is equal after Get fire up
		assert.NotNil(t, err)
		assert.Equal(t, "group cannot be null nor empty", err.Error())

		// Group ref is empty
		assert.Empty(t, _group.group)
	})

	t.Run("get valid Group ref if app is valid", func(t *testing.T) {

		// we define some vars
		validGroup := &Group{}
		_ = json.Unmarshal([]byte(mockserver.GroupsArray), validGroup)

		// Try to create Groups
		_group := NewGroups(marathon.New(server.URL))

		// try to Get with empty app id
		err := _group.Get(validGroup.ID).Restart(true)

		// Group is equal after Get fire up
		assert.Nil(t, err)

		// Group ref must be not empty
		assert.NotEmpty(t, _group.group)

		// Check some values on response
		assert.Equal(t, http.StatusOK, _group.client.StatusCode())
	})
}

func TestGroups_Suspend(t *testing.T) {

	// We create a Mock Server
	server := mockserver.MockServer()
	defer server.Close()

	t.Run("get empty Group ref if app is empty", func(t *testing.T) {

		// Try to create Groups
		_group := NewGroups(marathon.New(server.URL))

		// try to Get with empty app id
		err := _group.Suspend(true)

		// Group is equal after Get fire up
		assert.NotNil(t, err)
		assert.Equal(t, "group cannot be null nor empty", err.Error())

		// Group ref is empty
		assert.Empty(t, _group.group)
	})

	t.Run("get valid Group ref if app is valid", func(t *testing.T) {

		// we define some vars
		validGroup := &Group{}
		_ = json.Unmarshal([]byte(mockserver.GroupsArray), validGroup)

		// Try to create Groups
		_group := NewGroups(marathon.New(server.URL))

		// try to Get with empty app id
		err := _group.Get(validGroup.ID).Suspend(true)

		// Group is equal after Get fire up
		assert.Nil(t, err)

		// Group ref must be not empty
		assert.NotEmpty(t, _group.group)

		// Check some values on response
		assert.Equal(t, http.StatusOK, _group.client.StatusCode())
	})
}

func TestGroups_Apply(t *testing.T) {

	// We create a Mock Server
	server := mockserver.MockServer()
	defer server.Close()

	t.Run("get nil when Load is called with invalid file", func(t *testing.T) {

		// We define some vars
		fileName := "dumpfile.conf"

		// Try to create Groups
		_group := NewGroups(marathon.New(server.URL))

		// try to Get with empty app id
		err := _group.Load(fileName).Apply(true)
		defer os.Remove(fileName)

		//
		assert.NotNil(t, err)
		assert.Equal(t, "group cannot be null nor empty", err.Error())
		assert.Empty(t, _group.group)
	})

	t.Run("get App ref when is called with a valid file", func(t *testing.T) {

		// we define some vars
		fileName := "dumpfile.json"

		// We create out file to read as JSON
		errFile := ioutil.WriteFile(fileName, []byte(mockserver.GroupsArray), 0644)
		defer os.Remove(fileName)

		// We get not error
		assert.Nil(t, errFile)

		// Try to create Groups
		_group := NewGroups(marathon.New(server.URL))

		// try to Get with empty app id
		err := _group.Load(fileName).Apply(true)

		//
		assert.Nil(t, err)

		// Check some values on response
		assert.Equal(t, http.StatusOK, _group.client.StatusCode())
		assert.Equal(t, "/infra", _group.group.ID)

	})
}

func TestGroups_Load(t *testing.T) {

	// We create a Mock Server
	server := mockserver.MockServer()
	defer server.Close()

	t.Run("get nil when Load is called with invalid file", func(t *testing.T) {

		// We define some vars
		fileName := "dumpfile.conf"

		// Try to create Groups
		_group := NewGroups(marathon.New(server.URL))

		// try to Get with empty app id
		_ = _group.Load(fileName)
		defer os.Remove(fileName)

		//
		assert.NotNil(t, _group)
		assert.Empty(t, _group.group)
	})

	t.Run("get App ref when is called with a valid file", func(t *testing.T) {

		// we define some vars
		fileName := "dumpfile.json"

		// We create out file to read as JSON
		errFile := ioutil.WriteFile(fileName, []byte(mockserver.GroupsArray), 0644)
		defer os.Remove(fileName)

		// We get not error
		assert.Nil(t, errFile)

		// Try to create Groups
		_group := NewGroups(marathon.New(server.URL))

		// try to Get with empty app id
		_ = _group.Load(fileName)

		// Check some values on response
		assert.Equal(t, "/infra", _group.group.ID)
	})
}

func TestGroups_Dump(t *testing.T) {

	// We create a Mock Server
	server := mockserver.MockServer()
	defer server.Close()

	t.Run("get nil when Load is called with invalid file", func(t *testing.T) {

		// We define some vars
		fileName := "dumpfile.conf"

		// Try to create Groups
		_group := NewGroups(marathon.New(server.URL))

		// try to Get with empty app id
		err := _group.Dump(fileName)
		defer os.Remove(fileName)

		//
		assert.NotNil(t, err)
		assert.Equal(t, "group cannot be null nor empty", err.Error())
	})

	t.Run("get App ref when is called with a valid file", func(t *testing.T) {

		// we define some vars
		fileName := "dumpfile.json"

		// Try to create Groups
		_group := NewGroups(marathon.New(server.URL))

		// try to Get with empty app id
		err := _group.Get("/infra").Dump(fileName)
		defer os.Remove(fileName)

		// We need create a []byte with groupArray content because Dump writes that
		groupRef, _ := json.MarshalIndent(_group.group, "", "  ")

		//
		assert.Nil(t, err)

		// We read the file to compare
		file, errFile := ioutil.ReadFile(fileName)

		// We get not error
		assert.Nil(t, errFile)

		// Check some values on response
		assert.Equal(t, groupRef, file)
	})
}
