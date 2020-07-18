package funcBotIntegration

import (
	"cloud.google.com/go/firestore"
	"context"
	"github.com/sp0x/ihcph/common"
	"github.com/sp0x/ihcph/telegram"
)

var initialized = false
var globalContext *Context

type Context struct {
	Bots     *telegram.BotInterface
	Firebase *firestore.Client
	ctx      context.Context
}

func Initialize() *Context {
	if initialized {
		return globalContext
	}
	common.BindConfig()
	initialized = true
	ctxt := &Context{}
	fbase, err := common.NewFirebaseFromEnv()
	if err != nil {
		panic(err)
	}
	ctxt.Firebase = fbase
	ctxt.ctx = context.Background()
	ctxt.Bots = telegram.NewBotInterface()
	globalContext = ctxt
	return ctxt
}
