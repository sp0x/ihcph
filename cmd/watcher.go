package main

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/sp0x/ihcph/funcExtractResults"
	"github.com/sp0x/ihcph/telegram"
	"github.com/sp0x/torrentd/bots"
	"github.com/sp0x/torrentd/config"
	"github.com/sp0x/torrentd/indexer"
	"github.com/sp0x/torrentd/indexer/search"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
	"time"
)

var bot *telegram.BotInterface

func runWatcher(_ *cobra.Command, _ []string) {
	indexer.Loader = funcExtractResults.GetIndexLoader(appName)
	//Construct our facade based on the needed indexer.
	indexerFacade, err := indexer.NewFacade(indexSite, &appConfig)
	if err != nil {
		fmt.Printf("Couldn't initialize the named indexer `%s`: %s", indexSite, err)
		os.Exit(1)
	}
	if indexerFacade == nil {
		fmt.Printf("Indexer facade was nil")
		os.Exit(1)
	}
	watchIntervalSec := 30
	isSingleRun := viper.GetBool("single_run")
	bot = loadTelegram()
	var resultsChan <-chan *search.ExternalResultItem
	if isSingleRun {
		resultsChan = indexer.GetAllPagesFromIndex(indexerFacade, nil)
		broadcastResults(resultsChan)
	} else {
		resultsChan = indexer.Watch(indexerFacade, nil, watchIntervalSec)
		waitForResultsAndBroadcastThem(resultsChan)
	}
}

func loadTelegram() *telegram.BotInterface {
	//token := viper.GetString("telegram_token")
	//tmpTelegram, err := bots.NewTelegram(token, &appConfig, tgbotapi.NewBotAPI)
	//if err != nil {
	//	fmt.Printf("Couldn't initialize telegram bot: %v", err)
	//	os.Exit(1)
	//}
	return telegram.NewBotInterface()
}

func broadcastResults(resultsChan <-chan *search.ExternalResultItem) {
	for {
		tgram, err := bots.NewTelegram(viper.GetString("telegram_token"), &config.ViperConfig{}, tgbotapi.NewBotAPI)
		if err != nil {
			panic(err)
		}
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
				tgram.Broadcast(message)
			}
		case <-time.After(10 * time.Second):
			fmt.Printf("Timed out waiting for result")
			return
		}
	}
}

//Reads the channel that's the result of watching an indexer.
func waitForResultsAndBroadcastThem(resultsChan <-chan *search.ExternalResultItem) {
	chatMessagesChannel := make(chan bots.ChatMessage)
	tgram, err := bots.NewTelegram(viper.GetString("telegram_token"), &config.ViperConfig{}, tgbotapi.NewBotAPI)
	if err != nil {
		panic(err)
	}
	go func() {
		err := tgram.Run()
		if err != nil {
			fmt.Print(err)
			os.Exit(1)
		}
	}()
	go func() {
		_ = tgram.FeedBroadcast(chatMessagesChannel)
	}()
	for result := range resultsChan {
		if result == nil {
			break
		}
		if result.IsNew() || result.IsUpdate() {
			link := result.Site
			availableTime := result.GetField("time")
			msgText := fmt.Sprintf("I found a new opening at %s:\t%s\n", link, availableTime)
			message := bots.ChatMessage{Text: msgText, Banner: result.Banner}
			chatMessagesChannel <- message
			fmt.Printf("Found a new opening: %s", availableTime)
		}
	}
}
