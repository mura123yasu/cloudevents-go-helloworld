package main

import (
	"context"

	cloudevents "github.com/cloudevents/sdk-go"
)

func main() {
	event := cloudevents.NewEvent()
	event.SetID("ABC-123")
	event.SetType("com.cloudevents.readme.sent")
	// event.SetSource("http://localhost:8080/")
	event.SetSource("somesource")
	event.SetData(map[string]string{"hello": "world"})

	t, err := cloudevents.NewHTTPTransport(
		cloudevents.WithTarget("http://localhost:8080/"),
		cloudevents.WithEncoding(cloudevents.HTTPBinaryV1),
	)
	if err != nil {
		panic("failed to create transport, " + err.Error())
	}

	c, err := cloudevents.NewClient(t)
	if err != nil {
		panic("unable to create cloudevent client: " + err.Error())
	}

	_, _, err = c.Send(context.Background(), event)
	if err != nil {
		panic("failed to send cloudevent: " + err.Error())
	}
}
