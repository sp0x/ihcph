package common

import (
	"cloud.google.com/go/firestore"
	"context"
	"github.com/spf13/viper"
	"google.golang.org/api/option"
)

func NewFirebaseFromEnv() (*firestore.Client, error) {
	return NewFirebase(viper.GetString("firebase_project"), viper.GetString("firebase_credentials"))
}

func NewFirebase(projectId string, credentialsFile string) (*firestore.Client, error) {
	ctx := context.Background()
	var options []option.ClientOption
	if credentialsFile != "" {
		options = append(options, option.WithCredentialsFile(credentialsFile))
	}
	// credentials file option is optional, by default it will use GOOGLE_APPLICATION_CREDENTIALS
	// environment variable, this is a default method to connect to Google services
	client, err := firestore.NewClient(ctx, projectId, options...)
	return client, err
}
