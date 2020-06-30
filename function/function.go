package function

import (
	"encoding/json"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/sp0x/ihcph"
	"github.com/sp0x/torrentd/bots"
	"github.com/sp0x/torrentd/config"
	"github.com/sp0x/torrentd/indexer"
	"github.com/sp0x/torrentd/indexer/search"
	"github.com/spf13/viper"
	"net/http"
	"os"
	"time"
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

var (
	indexFacade *indexer.Facade
)

var bot *ihcph.BotInterface

func init() {
	var err error
	indexer.Loader = ihcph.GetIndexLoader(appName)
	//Construct our facade based on the needed indexer.
	cfg := getConfig()
	indexFacade, err = indexer.NewFacade("ihcph", cfg)
	if err != nil {
		fmt.Printf("Couldn't initialize the named indexer `%s`: %s", "ihcph", err)
		os.Exit(1)
	}
	if indexFacade == nil {
		fmt.Printf("Indexer facade was nil")
		os.Exit(1)
	}
	bot = loadTelegram(cfg)
}

func getConfig() config.Config {
	return &config.ViperConfig{}
}

func loadTelegram(cfg config.Config) *ihcph.BotInterface {
	token := viper.GetString("telegram_token")
	tmpTelegram, err := bots.NewTelegram(token, cfg, tgbotapi.NewBotAPI)
	if err != nil {
		fmt.Printf("Couldn't initialize telegram bot: %v", err)
		os.Exit(1)
	}
	return &ihcph.BotInterface{Telegram: tmpTelegram}
}

func broadcastResults(resultsChan <-chan *search.ExternalResultItem) {
	for {
		select {
		case result := <-resultsChan:
			//This signals that our channel has been closed.
			if result == nil {
				return
			}
			if result.IsNew() || result.IsUpdate() {
				link := result.Site
				availableTime := result.GetField("time")
				if availableTime == "" {
					continue
				}
				msgText := fmt.Sprintf("I found a new opening at %s:\t%s\n", link, availableTime)
				message := &bots.ChatMessage{Text: msgText, Banner: result.Banner}
				bot.Telegram.Broadcast(message)
			}
		case <-time.After(10 * time.Second):
			fmt.Printf("Timed out waiting for result")
			break
		}
	}
}

func TestRequest(w http.ResponseWriter, r *http.Request) {
	resultsChan := indexer.GetAllPagesFromIndex(indexFacade, nil)
	broadcastResults(resultsChan)
	body := Body{}
	body.Message = "Scanned for new results."
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(Result{
		200,
		body.Message,
		"",
	})
}
