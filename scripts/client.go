package main

import (
	"encoding/json"
	"github.com/centrifugal/centrifuge-go"
	"log"
)

func main() {
	client := centrifuge.NewJsonClient("ws://localhost:8123/connection/websocket", centrifuge.Config{
		Token: "eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiJjZW50cmlmdWdvIiwiZXhwIjoxNjg4MTkxMTA0LCJpc3MiOiJMT0dJVFkiLCJzdWIiOiJlZmYzOGE5Yy1iZDM0LTQ0MTktYjNhMC1iYjNhZmMxNWM1NGQifQ.2zEQG5o3pI8hqlH0pM0JrFvny8g-KSV5uFwiKRQ2WyeOBcoSXaGlPcEjRQsmBGSVAsWoDVbvdmBKC0_QNh987g",
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
