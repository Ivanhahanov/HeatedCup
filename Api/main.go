package main

import (
	"HeatedCup/Api/mqtt"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type server struct{}

type Commands struct {
	Class   string  `json:"class"`
	Id      int     `json:"id"`
	Command string  `json:"command"`
	Timeout float32 `json:"timeout"`
	Mode    string  `json:"mode"`
}

func GetCapStatus() string {
	return "OK"
}

func SendMqttCommand(commands Commands) string {
	s := mqtt.Settings{Broker: "mosquitto", Port: 1883}
	client, err := s.MqttConnect()
	if err != nil {
		panic(err)
	}
	if commands.Class == "" {
		commands.Class = "cup"
	}
	if commands.Id == 0 {
		commands.Id = 1
	}
	if commands.Mode == "" {
		commands.Mode = "heat"
	}
	if commands.Timeout == 0 {
		commands.Timeout = 0.3
	}
	message := mqtt.CupMessage(commands)
	message.SendMessage(client)
	client.Disconnect(250)
	return ""
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
		result := SendMqttCommand(c)
		w.Write([]byte(fmt.Sprintf(`{"command": "%s", "timeout":"%f", "result":"%s"}`, c.Command, c.Timeout, result)))
	}
}

func main() {
	s := &server{}
	http.Handle("/", s)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
