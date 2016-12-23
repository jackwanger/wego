package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"runtime"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/repong/wego/dict"
)

// Auto versioning
var (
	env string
)

var port int
var dictPath string

func init() {
	flag.StringVar(&dictPath, "dict", "", "Comma separated string like dict1,dict2,...")
	flag.IntVar(&port, "port", 8000, "listen port")
}

// Result for validate and filter requests
type Result struct {
	Result interface{} `json:"result"`
}

func main() {
	flag.Parse()

	runtime.GOMAXPROCS(runtime.NumCPU())
	dict.Load(dictPath)

	r := mux.NewRouter()
	r.HandleFunc("/validate", validate).Methods("POST")
	r.HandleFunc("/filter", filter).Methods("POST")

	addr := fmt.Sprintf(":%d", port)
	fmt.Println("Listening at", addr)

	if env == "release" {
		log.Fatal(http.ListenAndServe(addr, r))
	} else {
		loggedRouter := handlers.LoggingHandler(os.Stdout, r)
		log.Fatal(http.ListenAndServe(addr, loggedRouter))
	}
}

func writeJSON(w http.ResponseWriter, v interface{}) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(v)
}

func validate(w http.ResponseWriter, r *http.Request) {
	text := r.FormValue("message")
	invalid := dict.ExistInvalidWord(text)
	if invalid {
		writeJSON(w, &Result{"false"})
	} else {
		writeJSON(w, &Result{"true"})
	}
}

func filter(w http.ResponseWriter, r *http.Request) {
	text := r.FormValue("message")
	filtered := dict.ReplaceInvalidWords(text)
	writeJSON(w, &Result{filtered})
}
