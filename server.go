package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()

	if err != nil {
		fmt.Printf("error loading env vars: %v", err)
		os.Exit(1)
	}

	mux := http.NewServeMux()

	mux.HandleFunc("/env", envHandler)

	mux.HandleFunc("/env/", envKeyHandler)

	err = http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("PORT")), mux)

	if err == http.ErrServerClosed {
		fmt.Println("server closed")
		os.Exit(1)
	} else if err != nil {
		fmt.Printf("error starting the server %v", err)
		os.Exit(1)
	} else {
		fmt.Println("server running on ", os.Getenv("PORT"))
	}

}

func envHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	env := make(map[string]string)
	for _, envVar := range os.Environ() {

		pair := strings.SplitN(envVar, "=", 2)

		env[pair[0]] = pair[1]
	}

	w.WriteHeader(http.StatusOK)

	w.Header().Set("Content-Type", "application/json")

	encoder := json.NewEncoder(w)
	err := encoder.Encode(env)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func envKeyHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	key := strings.TrimPrefix(r.URL.Path, "/env/")

	value := os.Getenv(key)

	if value == "" {
		w.WriteHeader(http.StatusNotFound)
	} else {
		w.WriteHeader(http.StatusOK)
	}

	_, err := w.Write([]byte(value))

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

}
