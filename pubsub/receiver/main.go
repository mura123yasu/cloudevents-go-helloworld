package main

import (
	"context"
	"fmt"
	"log"
	"os"

	cloudevents "github.com/cloudevents/sdk-go"
	"github.com/cloudevents/sdk-go/pkg/cloudevents/client"
	cepubsub "github.com/cloudevents/sdk-go/pkg/cloudevents/transport/pubsub"
	pscontext "github.com/cloudevents/sdk-go/pkg/cloudevents/transport/pubsub/context"
	"github.com/kelseyhightower/envconfig"
)

type envConfig struct {
	ProjectID string `envconfig:"GOOGLE_CLOUD_PROJECT"`

	TopicID string `envconfig:"PUBSUB_TOPIC" default:"demo_cloudevents" required:"true"`

	SubscriptionID string `envconfig:"PUBSUB_SUBSCRIPTION" default:"demo_cloudevents_subscriber" required:"true"`
}

type Model struct {
	Sequence int    `json:"id"`
	Message  string `json:"message"`
}

func receive(ctx context.Context, event cloudevents.Event, resp *cloudevents.EventResponse) error {
	fmt.Printf("Event Context: %+v\n", event.Context)

	fmt.Printf("Transport Context: %+v\n", pscontext.TransportContextFrom(ctx))

	data := &Model{}
	if err := event.DataAs(data); err != nil {
		fmt.Printf("Got Data Error: %s\n", err.Error())
	}
	fmt.Printf("Data: %+v\n", data)

	fmt.Printf("----------------------------\n")
	return nil
}

func main() {
	ctx := context.Background()

	var env envConfig
	if err := envconfig.Process("", &env); err != nil {
		log.Printf("[ERROR] Failed to process env var: %s", err)
		os.Exit(1)
	}

	log.Printf("[INFO] ProjectID: %s", env.ProjectID)
	log.Printf("[INFO] TopicID: %s", env.TopicID)
	log.Printf("[INFO] SubscriptionID: %s", env.SubscriptionID)
	t, err := cepubsub.New(context.Background(),
		cepubsub.WithProjectID(env.ProjectID),
		// cepubsub.WithTopicID(env.TopicID),
		// cepubsub.WithSubscriptionAndTopicID(env.SubscriptionID, env.TopicID))
		cepubsub.WithSubscriptionID(env.SubscriptionID))
	if err != nil {
		log.Fatalf("failed to create pubsub transport, %s", err.Error())
	}
	c, err := client.New(t)
	if err != nil {
		log.Fatalf("failed to create client, %s", err.Error())
	}

	log.Println("Created client, listening...")

	if err := c.StartReceiver(ctx, receive); err != nil {
		log.Fatalf("failed to start pubsub receiver, %s", err.Error())
	}
}
