package mqtt

import (
	"encoding/json"
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type CupMessage struct {
	Class   string  `json:"class"`
	Id      int     `json:"id"`
	Command string  `json:"command"`
	Timeout float32 `json:"timeout"`
	Mode    string  `json:"mode"`
}

type message struct {
	Command string  `json:"command"`
	Timeout float32 `json:"timeout"`
	Mode    string  `json:"mode"`
}

func (m *CupMessage) SendMessage(client mqtt.Client) (err error) {
	mes := message{m.Command, m.Timeout, m.Mode}
	mqttMessage, err := json.Marshal(mes)
	if err != nil {
		return err
	}
	topic := fmt.Sprintf("%s/%d", m.Class, m.Id)
	token := client.Publish(topic, 0, false, mqttMessage)
	token.Wait()
	fmt.Println(string(mqttMessage))
	return nil
}
