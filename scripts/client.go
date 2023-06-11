package main

import (
	"encoding/json"
	"github.com/centrifugal/centrifuge-go"
	"log"
)

func main() {
	client := centrifuge.NewJsonClient("ws://localhost:8123/connection/websocket", centrifuge.Config{
		Token: "eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiJjZW50cmlmdWdvIiwiZXhwIjoxNjg2NTAzMzM5LCJpc3MiOiJMT0dJVFkiLCJzdWIiOiIwNmU5NTFmMi04OTQzLTQ2YTgtYTA4Yi04ZTgxNTZhMjQzNDIifQ.kH7NrkrfNpePGC7QuyBEGQX4hLu-XNXPnKsl-pamrFWDQHE6d7eZLIamLoO1t2PgUGjmAFC3X29_CImGhehdsg",
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
