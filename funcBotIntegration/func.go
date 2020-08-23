package funcBotIntegration

import (
	"cloud.google.com/go/firestore"
	"context"
	log "github.com/sirupsen/logrus"
	"github.com/sp0x/ihcph/common"
	"github.com/spf13/viper"
)

var initialized = false
var globalContext *Context

type Context struct {
	Firebase *firestore.Client
	ctx      context.Context
}

func Initialize() (*Context, error) {
	if initialized {
		return globalContext, nil
	}
	log.SetLevel(log.InfoLevel)
	common.BindConfig()
	verboseValue := viper.GetString("VERBOSE")
	if verboseValue != "" {
		log.SetLevel(log.DebugLevel)
	}
	initialized = true
	ctxt := &Context{}
	fbase, err := common.NewFirebaseFromEnv()
	if err != nil {
		return nil, err
	}
	ctxt.Firebase = fbase
	ctxt.ctx = context.Background()
	globalContext = ctxt
	return ctxt, nil
}
