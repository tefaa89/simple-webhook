package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

type Payload struct {
	WaitInSeconds int
	StatusCode    int
}

func initPort() int {
	port := 8080
	if len(os.Args) > 1 {
		arg := os.Args[1]
		log.Printf("Start webhook service at port: %s\n", arg)
		i1, err := strconv.Atoi(arg)
		if err == nil {
			port = i1
		}
	} else {
		log.Printf("Start webhook service at port: %d\n", port)
	}
	return port
}

func handleCall(writer http.ResponseWriter, request *http.Request) {
	defer request.Body.Close()
	var payload Payload

	switch request.Method {
	case "POST":
		// read body
		body, err := ioutil.ReadAll(request.Body)
		if err != nil {
			panic(err)
		}

		// parse JSON from body
		err = json.Unmarshal(body, &payload)
		if err != nil {
			fmt.Fprintf(writer, "payload is wrong: %s", string(body))
			panic(err)
		}

		// process waiting
		if payload.WaitInSeconds > 0 {
			log.Printf("Waiting %d seconds", payload.WaitInSeconds)
			time.Sleep(time.Duration(payload.WaitInSeconds) * time.Second)
		}

		// process response status code
		if http.StatusText(payload.StatusCode) != "" {
			writer.WriteHeader(payload.StatusCode)
		} else {
			writer.WriteHeader(200)
		}
	}
}

func main() {
	port := initPort()

	http.HandleFunc("/call", handleCall)
	if err := http.ListenAndServe(":"+strconv.Itoa(port), nil); err != nil {
		panic(err)
	}
}
