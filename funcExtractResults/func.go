package funcExtractResults

import (
	"fmt"
	"github.com/sp0x/ihcph/telegram"
	"github.com/sp0x/torrentd/config"
	"github.com/sp0x/torrentd/indexer"
	"github.com/spf13/viper"
	"os"
)

const (
	appName = "ihcph"
)

var initialized = false
var globalContext *Context

type Context struct {
	Bot         *telegram.BotInterface
	IndexFacade *indexer.Facade
}

//Executed on Cold boot.
func Initialize() *Context {
	if initialized {
		return globalContext
	}
	var err error
	initialized = true
	indexer.Loader = GetIndexLoader(appName)
	//Construct our facade based on the needed indexer.
	cfg := getConfig()
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
	context.Bot = telegram.GetTelegram(viper.GetString("telegram_token"), cfg)
	context.IndexFacade = indexFacade
	globalContext = context
	return context
}

func getConfig() config.Config {
	c := &config.ViperConfig{}
	_ = c.Set("storage", "firebase")
	_ = c.Set("firebase_project", "firebase")
	_ = c.Set("firebase_credentials_file", "firebase")
	return c
}
