package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

var CupUrl = os.Getenv("CupUrl")

type server struct{}

type Commands struct {
	Command string `json:"command"`
	Timeout int    `json:"timeout"`
}

func GetCapStatus() string {
	return "OK"
}

func SendCommand(commands *Commands) string {
	data, err := json.Marshal(commands)
	resp, err := http.Post(CupUrl, "application/json",
		bytes.NewBuffer(data))

	if err != nil {
		log.Fatal(err)
	}

	var res string
	json.NewDecoder(resp.Body).Decode(&res)
	return res
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	status := GetCapStatus()
	switch r.Method {
	case "GET":
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(fmt.Sprintf(`{"message": "%s"}`, status)))
	case "POST":
		w.WriteHeader(http.StatusCreated)
		var c Commands
		err := json.NewDecoder(r.Body).Decode(&c)
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}
		result := SendCommand(&c)
		w.Write([]byte(fmt.Sprintf(`{"command": "%s", "timeout":"%d", "result":"%s"}`, c.Command, c.Timeout, result)))
	}
}

func main() {
	s := &server{}
	http.Handle("/", s)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
