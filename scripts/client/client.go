package main

import (
	"encoding/json"
	"github.com/centrifugal/centrifuge-go"
	"log"
)

const Host = `ws://logity-centrifugo/connection/websocket`
const Token = "eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiJjZW50cmlmdWdvIiwiZXhwIjoxNjkyMDE4NTUxLCJpc3MiOiI4MDgwIiwic3ViIjoiN2ExMjFiMjUtOGY2My00ZjZkLWEzYjctMTJlNDkzNjcxMTQzIn0.vu0btsc6WSX4J0sU5riapaLkMqQzsJFhhnTWP0jaW5DdfozAOINfTSbrrNuxnIwC4c3zMvnW5hKC-vHQdLkasg"

func main() {
	client := centrifuge.NewJsonClient(Host, centrifuge.Config{
		Token: Token,
	})

	defer client.Close()

	client.OnConnecting(func(e centrifuge.ConnectingEvent) {
		log.Printf("Connecting - %d (%s)\n", e.Code, e.Reason)
	})
	client.OnConnected(func(e centrifuge.ConnectedEvent) {
		log.Printf("Connected with ID %s\n", e.ClientID)
	})
	client.OnPublication(func(event centrifuge.ServerPublicationEvent) {
		var msg map[string]interface{}
		_ = json.Unmarshal(event.Data, &msg)
		bt, _ := json.MarshalIndent(msg, "  ", " ")
		log.Printf("Publication: (%s)\n", string(bt))
	})
	client.OnDisconnected(func(e centrifuge.DisconnectedEvent) {
		log.Printf("Disconnected: %d (%s)\n", e.Code, e.Reason)
	})

	_ = client.Connect()
	for {

	}
}
