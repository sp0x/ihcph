package funcExtractResults

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/sp0x/ihcph/common"
	"github.com/sp0x/torrentd/indexer"
	"os"
)

const (
	appName = "ihcph"
)

var initialized = false
var globalContext *Context

type Context struct {
	IndexFacade *indexer.Facade
}

//Executed on Cold boot.
func Initialize() *Context {
	if initialized {
		return globalContext
	}
	var err error
	initialized = true
	log.SetLevel(log.InfoLevel)
	common.BindConfig()
	cfg := common.GetConfig()
	indexer.Loader = GetIndexLoader(appName)
	//Construct our facade based on the needed indexer.

	indexFacade, err := indexer.NewFacade("ihcph", cfg)
	if err != nil {
		fmt.Printf("Couldn't initialize the named indexer `%s`: %s", "ihcph", err)
		os.Exit(1)
	}
	if indexFacade == nil {
		fmt.Printf("Indexer facade was nil")
		os.Exit(1)
	}
	context := &Context{}
	//context.Bots = telegram.NewBotInterface()
	context.IndexFacade = indexFacade
	globalContext = context
	return context
}
