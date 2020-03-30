package gpubsub

import (
	"context"

	"cloud.google.com/go/pubsub"
	"github.com/sirupsen/logrus"
)

type publishReturn struct {
	status string
	msg    string
	id     string
}

func Publish(projectID, topicID, msg string) *publishReturn {

	ctx := context.Background()

	// create connection using projectID
	client, err := pubsub.NewClient(ctx, projectID)

	if err != nil {
		logrus.Fatal(err)
		return &publishReturn{status: "failure", msg: "google pubsup connection failed"}
	}

	// establish channel to the topic
	topic := client.Topic(topicID)

	//publish the message

	result := topic.Publish(ctx, &pubsub.Message{Data: []byte(msg)})

	// wait for the server generated ID for the published msg

	msgId, err := result.Get(ctx)

	if err != nil {
		logrus.Fatal("google pubsup message publish failed")
		client.Close()
		return &publishReturn{status: "failure", msg: "google pubsup message publish failed"}
	}

	client.Close()
	return &publishReturn{status: "success", msg: "message published successfully", id: msgId}
}
