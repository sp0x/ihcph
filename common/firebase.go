package common

import (
	"cloud.google.com/go/firestore"
	"context"
	"google.golang.org/api/option"
)

func NewFirebaseFromEnv() (*firestore.Client, error) {
	c := GetFirebaseConfig()
	return NewFirebase(c.Project, c.Credentials)
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
