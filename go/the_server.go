// Serving stuff for the_script
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"path"
	"runtime"
	"strings"
)

var (
	the_bag_of_json = map[string]string{
		"a_key": `{"a":1}`,
		"b_key": `{"b":2}`,
	}
)

func main() {
	http.HandleFunc("/api/", ApiHandler)
	http.HandleFunc("/log", LogHandler)
	http.ListenAndServe("127.0.0.1:9999", nil)
}

func ApiHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("[api_request]: %s\n", r.URL.Path)
	key := strings.Split(r.URL.Path, "/")[2]
	response := the_bag_of_json[key]
	fmt.Fprintf(w, response)
}

func LogHandler(w http.ResponseWriter, r *http.Request) {
	_, THE_SERVER_PATH, _, _ := runtime.Caller(0)
	THE_THING_DIR := path.Dir(path.Dir(THE_SERVER_PATH))
	value, err := ioutil.ReadFile(THE_THING_DIR + "/log/the_log")
	if err != nil {
		fmt.Fprintf(w, fmt.Sprintf("%s", err))
		return
	}
	js, _ := json.Marshal(string(value))
	callback := r.URL.Query().Get("callback")
	fmt.Fprintf(w, callback+"("+string(js)+")")
}
