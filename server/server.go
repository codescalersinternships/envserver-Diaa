package server

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"
)

var ErrInvalidPort = errors.New("port should be between 1024:65000")

type App struct {
	port int
}

func NewApp() *App {
	return &App{}
}

func (app *App) SetPort(p int) error {
	if p < 1024 || p > 65000 {
		return ErrInvalidPort
	}
	app.port = p
	return nil
}

func (app *App) Run() error {
	mux := http.NewServeMux()

	mux.HandleFunc("/env", envHandler)

	mux.HandleFunc("/env/", envHandler)

	err := http.ListenAndServe(fmt.Sprintf(":%d", app.port), mux)

	if err == http.ErrServerClosed {
		return fmt.Errorf("server closed")

	}
	if err != nil {
		return fmt.Errorf("error starting the server %v", err)

	}
	return nil

}

func envHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	encoder := json.NewEncoder(w)

	key := strings.TrimPrefix(r.URL.Path, "/env")
	if key == "" {

		env := make(map[string]string)
		for _, envVar := range os.Environ() {

			pair := strings.SplitN(envVar, "=", 2)

			env[pair[0]] = pair[1]
		}

		err := encoder.Encode(env)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}
	key = key[1:]

	value := os.Getenv(key)

	if value == "" {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	err := encoder.Encode(value)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

}
