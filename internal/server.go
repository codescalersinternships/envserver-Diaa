package server

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"
)

// ErrInvalidPort is an error for port validation
var ErrInvalidPort = errors.New("port should be between 1024:65000")

// App is a struct contains the program configs like port
type App struct {
	Port int
}

// NewApp factory function for the App struct. returns a newApp instance with the port
func NewApp(p int) (*App, error) {

	if p < 1 || p > 65535 {
		return nil, ErrInvalidPort
	}
	return &App{Port: p}, nil
}

// Run is a function that starts the server
func (app *App) Run() error {
	mux := http.NewServeMux()

	mux.HandleFunc("/env", envHandler)

	mux.HandleFunc("/env/", envHandler)

	err := http.ListenAndServe(fmt.Sprintf(":%d", app.Port), mux)

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

	key := strings.TrimPrefix(r.URL.Path, "/env")

	switch key {
	case "":
		handleGetEnv(w, r)
	default:
		handleGetKey(w, r)

	}

}

func handleGetEnv(w http.ResponseWriter, r *http.Request) {
	encoder := json.NewEncoder(w)
	w.Header().Set("Content-Type", "application/json")

	env := make(map[string]string)
	for _, envVar := range os.Environ() {

		pair := strings.SplitN(envVar, "=", 2)

		env[pair[0]] = pair[1]
	}

	err := encoder.Encode(env)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func handleGetKey(w http.ResponseWriter, r *http.Request) {
	key := strings.TrimPrefix(r.URL.Path, "/env/")
	encoder := json.NewEncoder(w)

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
