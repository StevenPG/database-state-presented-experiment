package pkg

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

type States struct {
	States	[]CState	`json:"state"`
}

type CState struct {
	Value		string	`json:"key"`
}

type OutboundRequest struct {
	Key		string	`json:"key"`
	Value	string	`json:"value"`
	States	[]CState	`json:"state"`
}

// LoopRequest perform the loop
func LoopRequest(d time.Duration, f func(key string), uniqueKey string) {
	for {
		time.Sleep(d)
		log.Printf("Sending request for %s", uniqueKey)
		go f(uniqueKey)
	}
}

func LoopRequestFullSpeed( f func(key string), uniqueKey string) {
	for {
		log.Printf("Sending request for %s", uniqueKey)
		go f(uniqueKey)
	}
}

// MakeRequestWithRandomUpdate make out bound call
func MakeRequestWithRandomUpdate(uniqueKey string) {
	// call get endpoint to get current states
	res, _ := http.Get("http://localhost:8080/get")

	var states *States
	data, _ := ioutil.ReadAll(res.Body)
	json.Unmarshal(data, &states)

	randomValue := strconv.Itoa(rand.Int())

	request := OutboundRequest{
		Key:    uniqueKey,
		Value:  randomValue,
		States: states.States,
	}

	jsonValue, _ := json.Marshal(request)
	postResult, _ := http.Post("http://localhost:8080/put", "application/json", bytes.NewBuffer(jsonValue))

	if postResult.StatusCode != 202 {
		data, _ := ioutil.ReadAll(postResult.Body)
		fmt.Println(string(data))
	}
}
