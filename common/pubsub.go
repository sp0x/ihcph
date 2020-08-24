package common

import (
	"context"
	"github.com/gogo/protobuf/proto"
	"github.com/lileio/pubsub"
	"github.com/lileio/pubsub/middleware/defaults"
	"github.com/lileio/pubsub/providers/google"
)

func setupPubSub(projectId string) error {
	provider, err := google.NewGoogleCloud(projectId)
	if err != nil {
		return err
	}
	//Service credentials exposed through: GOOGLE_APPLICATION_CREDENTIALS
	pubsub.SetClient(&pubsub.Client{
		ServiceName: "my-service-name",
		Provider:    provider,
		Middleware:  defaults.Middleware,
	})
	return nil
}

func publishSchemeStatus(ctx context.Context, data proto.Message) {
	pubsub.Publish(ctx, "scrapescheme", data)
}
