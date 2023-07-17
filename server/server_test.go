package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"reflect"
)

func TestGetEnvKey(t *testing.T) {

	testCases := []struct {
		name            string
		keyToSet        string
		keyToSearch     string
		valueToSet      string
		expStatCode     int
		expResponseBody string
		failureMessage  string
	}{
		{
			name:            "test with existing key",
			expStatCode:     http.StatusOK,
			keyToSet:        "TestKey",
			keyToSearch:     "TestKey",
			valueToSet:      "TestValue",
			expResponseBody: "TestValue",
			failureMessage:  "failed to get an existing key",
		}, {
			name:            "test with non existing key",
			expStatCode:     http.StatusNotFound,
			keyToSet:        "TestKey",
			valueToSet:      "TestValue",
			keyToSearch:     "InvalidKey",
			expResponseBody: "",
			failureMessage:  "get non existing key",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/env/%s", tc.keyToSearch), nil)

			res := httptest.NewRecorder()

			os.Setenv(tc.keyToSet, tc.valueToSet)

			defer os.Unsetenv(tc.keyToSet)

			envKeyHandler(res, req)

			assert.Equal(t, tc.expStatCode, res.Code, tc.failureMessage)

			assert.Equal(t, tc.expResponseBody, res.Body.String(), tc.failureMessage)
		})
	}

	t.Run("sending request with method post", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodPost, "/env/key", nil)

		res := httptest.NewRecorder()

		envHandler(res, req)

		assert.Equal(t, 404, res.Code)

	})
	t.Run("sending request with method patch", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodPatch, "/env", nil)

		res := httptest.NewRecorder()

		envHandler(res, req)

		assert.Equal(t, 404, res.Code)

	})

	t.Run("sending request with method delete", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodDelete, "/env", nil)
		res := httptest.NewRecorder()

		envHandler(res, req)

		assert.Equal(t, 404, res.Code)

	})

}

func TestGetEnv(t *testing.T) {

	type keyValue struct {
		key   string
		value string
	}
	testCases := []struct {
		name         string
		keysToSet    []keyValue
		expStatCode  int
		expKeyValMap map[string]string
	}{
		{
			name:      "test should pass with equal maps",
			keysToSet: []keyValue{{key: "key1", value: "value1"}, {key: "key2", value: "value2"}, {key: "key3", value: "value3"}},
			expKeyValMap: map[string]string{
				"key1": "value1",
				"key2": "value2",
				"key3": "value3",
			},
			expStatCode: http.StatusOK,
		},
		{
			name:         "test empty env",
			keysToSet:    []keyValue{},
			expKeyValMap: map[string]string{},
			expStatCode:  http.StatusOK,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			req, _ := http.NewRequest(http.MethodGet, "/env", nil)

			res := httptest.NewRecorder()

			oldEnv := os.Environ()

			os.Clearenv()

			defer func() {
				os.Clearenv()
				for _, env := range oldEnv {
					pair := strings.SplitN(env, "=", 2)
					os.Setenv(pair[0], pair[1])
				}
			}()

			for _, pair := range tc.keysToSet {
				os.Setenv(pair.key, pair.value)
			}

			envHandler(res, req)

			assert.Equal(t, tc.expStatCode, res.Code, "got %d status code than 200", res.Code)

			assert.Equal(t, "application/json", res.Header().Get("Content-Type"), " got %s content type than json", res.Header().Get("Content-Type"))

			var got map[string]string
			json.NewDecoder(res.Body).Decode(&got)

			if !reflect.DeepEqual(got, tc.expKeyValMap) {
				t.Errorf("expected map %v but got %v", tc.expKeyValMap, got)
			}
		})
	}

	t.Run("sending request with method post", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodPost, "/env", nil)

		res := httptest.NewRecorder()

		envHandler(res, req)

		assert.Equal(t, 404, res.Code)

	})

	t.Run("sending request with method patch", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodPatch, "/env", nil)

		res := httptest.NewRecorder()

		envHandler(res, req)

		assert.Equal(t, 404, res.Code)

	})

	t.Run("sending request with method delete", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodDelete, "/env", nil)
		res := httptest.NewRecorder()

		envHandler(res, req)

		assert.Equal(t, 404, res.Code)

	})

}


func TestSetPort(t * testing.T){
	app := App{}

	app.SetPort(50)

	want :=50

	assert.Equal(t,want,app.port)
}

