package main

import (
	"math/rand"
	"strconv"

	"github.com/RomeroGabriel/go-basic-event-handler/pkg/rabbitmq"
)

func randomString() string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	s := make([]rune, 10)
	for i := range s {
		s[i] = letters[rand.Intn(len(letters))]
	}
	return string(s)
}

func main() {
	ch := rabbitmq.OpenChannel()
	defer ch.Close()
	randomMsg := randomString()
	msg := "From Producer. MSG: " + randomMsg + "ID: " + strconv.Itoa(rand.Intn(100))
	rabbitmq.Publish(ch, msg, "amq.direct")
}
