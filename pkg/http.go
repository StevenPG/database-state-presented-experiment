package pkg

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

type Item struct {
	Key 	string `json:"key"`
	Value	string `json:"value"`
	State	[]State `json:"state"`
}

type State struct {
	Value		string `json:"key"`
}

type Values struct {
	State []State `json:"state"`
}

// RunHttpServer - executes the API that will be hit for the experiment
func RunHttpServer() {
	router := httprouter.New()
	router.GET("/get", allHandler)
	log.Print("Loading /get endpoint...")

	router.POST("/put", putHandler)
	log.Print("Loading /put endpoint...")

	log.Print("Starting HTTP Server on localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func allHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	values := &Values{State: GetBoltContentAsState()}
	valuesAsString, _ := json.Marshal(values)
	fmt.Fprintf(w, "%s\n", string(valuesAsString))
}

func putHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	item := &Item{}

	body, _ := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	json.Unmarshal(body, item)

	err := PutDatabaseItem(item)

	if err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusBadRequest)
	} else {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusAccepted)
	}
}
