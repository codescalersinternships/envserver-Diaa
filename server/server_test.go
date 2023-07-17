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

	t.Run("testing /env", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodGet, "/env", nil)

		res := httptest.NewRecorder()

		want := make(map[string]string)
		for _, env := range os.Environ() {
			pair := strings.SplitN(env, "=", 2)
			want[pair[0]] = pair[1]
		}

		envHandler(res, req)

		assert.Equal(t, 200, res.Code, "got %d status code than 200", res.Code)

		assert.Equal(t, "application/json", res.Header().Get("Content-Type"), " got %s content type than json", res.Header().Get("Content-Type"))

		var got map[string]string
		json.NewDecoder(res.Body).Decode(&got)

		if !reflect.DeepEqual(got, want) {
			t.Errorf("expected map %v but got %v", want, got)
		}
	})

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

func TestSetPort(t *testing.T) {
	app := App{}

	app.SetPort(50)

	want := 50

	assert.Equal(t, want, app.port)
}
