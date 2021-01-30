package mqtt

import (
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"testing"
)

func TestSettings_MqttConnect(t *testing.T) {
	s := Settings{Broker: "mosquitto", Port: 1883}
	client, err := s.MqttConnect()
	if err != nil {
		t.Error(err)
	}
	client.Disconnect(250)
}

func TestCupMessage_SendMessage(t *testing.T) {
	s := Settings{Broker: "mosquitto", Port: 1883}
	client, err := s.MqttConnect()
	if err != nil {
		t.Error(err)
	}

	sub(client)

	message := CupMessage{
		Class:   "cup",
		Id:      1,
		Command: "on",
		Timeout: 1,
		Mode:    "heat",
	}
	message.SendMessage(client)
	client.Disconnect(250)
}

func sub(client mqtt.Client) {
	topic := "cup/1"
	token := client.Subscribe(topic, 1, nil)
	token.Wait()
}