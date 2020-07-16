package funcBotIntegration

import (
	"cloud.google.com/go/firestore"
	"context"
	"github.com/sp0x/ihcph/common"
	"github.com/sp0x/ihcph/telegram"
	"github.com/spf13/viper"
)

var initialized = false
var globalContext *Context

type Context struct {
	Bot      *telegram.BotInterface
	Firebase *firestore.Client
	ctx      context.Context
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
	fbase, err := common.NewFirebaseFromEnv()
	if err != nil {
		panic(err)
	}
	ctxt.Firebase = fbase
	ctxt.ctx = context.Background()
	ctxt.Bot = telegram.NewBotInterface()
	globalContext = ctxt
	return ctxt
}
