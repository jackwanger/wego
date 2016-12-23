package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"runtime"

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
	runtime.GOMAXPROCS(runtime.NumCPU())
}

// Result for validate and filter requests
type Result struct {
	Result string `json:"result"`
}

func main() {
	flag.StringVar(&dictPath, "dict", "", "Directory path. Multiple directories use comma separated string like a,b,c")
	flag.IntVar(&port, "port", 8000, "listen port")
	flag.Parse()

	dict.Load(dictPath)
	fmt.Println("Listening at", port)

	r := mux.NewRouter()
	r.HandleFunc("/validate", validateFunc).Methods("POST")
	r.HandleFunc("/filter", filterFunc).Methods("POST")

	addr := fmt.Sprintf(":%d", port)
	log.Fatal(http.ListenAndServe(addr, r))
}

func validateFunc(w http.ResponseWriter, r *http.Request) {
	text := r.FormValue("message")
	if dict.ExistInvalidWord(text) {
		writeJSON(w, &Result{"true"})
	} else {
		writeJSON(w, &Result{"false"})
	}
}

func filterFunc(w http.ResponseWriter, r *http.Request) {
	text := r.FormValue("message")
	text = dict.ReplaceInvalidWords(text)
	writeJSON(w, &Result{text})
}

func writeJSON(w http.ResponseWriter, v interface{}) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(v)
}
