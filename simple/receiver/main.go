package main

import (
	"context"
	"fmt"
	"log"

	cloudevents "github.com/cloudevents/sdk-go"
)

func Receive(event cloudevents.Event) {
	// do something with event.Context and event.Data (via event.DataAs(foo)
	fmt.Printf("Event received.\n====\n%s====\n", event)
}

func main() {
	c, err := cloudevents.NewDefaultClient()
	if err != nil {
		log.Fatalf("failed to create client, %v", err)
	}
	log.Fatal(c.StartReceiver(context.Background(), Receive))
}
