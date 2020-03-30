# cloudevents-go-helloworld
My first trial of cloud events sdk-go.

## Try simple
```sh
git checkout https://github.com/mura123yasu/cloudevents-go-helloworld.git
go mod download
# start receiver
go run ./simple/receiver/main.go
# start sender at another process
go run ./simple/sender/main.go
```

## Try with Cloud Pub/Sub

### Setup Cloud Pub/Sub
```sh
gcloud pubsub topics create <YOUR PUBSUB TOPIC>
gcloud pubsub subscriptions create <YOUR PUBSUB SUBSCRIPTION> --topic=<YOUR PUBSUB TOPIC>
```

### Set env
```sh
# required
export GOOGLE_APPLICATION_CREDENTIALS=<YOUR CREDENTIAL>
export GOOGLE_CLOUD_PROJECT=<YOUR GCP PROJECT>
# optional
export PUBSUB_TOPIC=<YOUR PUBSUB TOPIC> # default is "demo_cloudevents"
export PUBSUB_SUBSCRIPTION=<YOUR PUBSUB SUBSCRIPTION> # default is "demo_cloudevents_subscriber"
```

### Run
```sh
git checkout https://github.com/mura123yasu/cloudevents-go-helloworld.git
go mod download
# start receiver
go run ./pubsub/receiver/main.go
# start sender at another process
go run ./pubsub/sender/main.go
```