package funcBotIntegration

import (
	"cloud.google.com/go/firestore"
	"context"
	"github.com/sp0x/ihcph/common"
)

var initialized = false
var globalContext *Context

type Context struct {
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
	globalContext = ctxt
	return ctxt
}
