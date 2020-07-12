package function

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/sp0x/torrentd/bots"
	"github.com/sp0x/torrentd/config"
	"github.com/sp0x/torrentd/indexer"
	"github.com/spf13/viper"
	"os"
)

type Result struct {
	Code    int
	Message string
	Token   string
}

type Body struct {
	Message string
}

const (
	appName = "ihcph"
)

var initialized = false
var globalContext *Context

type Context struct {
	Bot         *BotInterface
	IndexFacade *indexer.Facade
}

//Executed on Cold boot.
func Initialize() *Context {
	var err error
	if initialized {
		return globalContext
	}
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
	context.Bot = loadTelegram(cfg)
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

func loadTelegram(cfg config.Config) *BotInterface {
	token := viper.GetString("telegram_token")
	tmpTelegram, err := bots.NewTelegram(token, cfg, tgbotapi.NewBotAPI)
	if err != nil {
		fmt.Printf("Couldn't initialize telegram Bot: %v", err)
		os.Exit(1)
	}
	return &BotInterface{Telegram: tmpTelegram}
}
