package scrapeschemes

import (
	"context"
	"fmt"
	"github.com/lileio/pubsub"
	"github.com/lileio/pubsub/middleware/defaults"
	"github.com/lileio/pubsub/providers/google"
	"github.com/spf13/viper"
	"os"
)

const schemeTopic = "scrapecheme"

type ScrapeSchemeMessage struct{}

func setupPubsubConfig() {
	viper.AutomaticEnv()
	_ = viper.BindEnv("firebase_project")
	_ = viper.BindEnv("firebase_credentials_file")
}

func init() {
	setupPubsubConfig()
	projectId := viper.GetString("firebase_project")
	provider, err := google.NewGoogleCloud(projectId)
	if err != nil {
		fmt.Printf("couldn't initialize google pubsub provider")
		os.Exit(1)
	}
	//Service credentials exposed through: GOOGLE_APPLICATION_CREDENTIALS
	pubsub.SetClient(&pubsub.Client{
		ServiceName: "ihcph",
		Provider:    provider,
		Middleware:  defaults.Middleware,
	})
}

func PublishSchemeStatus(ctx context.Context, data *ScrapeSchemeMessage) {
	pubsub.PublishJSON(ctx, schemeTopic, data)
}
