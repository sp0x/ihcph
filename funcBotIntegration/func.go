package funcBotIntegration

import (
	"cloud.google.com/go/firestore"
	"context"
	"github.com/sp0x/ihcph/telegram"
	"github.com/sp0x/torrentd/indexer"
	"github.com/spf13/viper"
	"google.golang.org/api/option"
)

var initialized = false
var globalContext *Context

type Context struct {
	Bot         *telegram.BotInterface
	IndexFacade *indexer.Facade
	Firebase    *firestore.Client
	ctx         context.Context
}

func (c *Context) StoreNewIntegration(integration *telegram.Integration) error {
	fbase := c.Firebase
	nsbots := fbase.Collection("bots")
	newDoc := nsbots.Doc(integration.Token)
	existing, err := newDoc.Get(c.ctx)
	if existing != nil {
		return nil
	}
	integration.Id = newDoc.ID
	_, err = newDoc.Create(c.ctx, integration)
	return err
}

func initConfig() {
	_ = viper.BindEnv("firebase_project")
	_ = viper.BindEnv("firebase_credentials")
}

func Initialize() *Context {
	if initialized {
		return globalContext
	}
	initConfig()
	initialized = true
	ctxt := &Context{}
	fbase, err := newFirebase(viper.GetString("firebase_project"), viper.GetString("firebase_credentials"))
	if err != nil {
		panic(err)
	}
	ctxt.Firebase = fbase
	ctxt.ctx = context.Background()
	return ctxt
}

func newFirebase(projectId string, credentialsFile string) (*firestore.Client, error) {
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
