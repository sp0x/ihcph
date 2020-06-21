package main

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/sp0x/torrentd/bots"
	"github.com/sp0x/torrentd/indexer"
	"github.com/sp0x/torrentd/indexer/search"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

func runWatcher(_ *cobra.Command, _ []string) {
	indexer.Loader = getIndexLoader()
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
	var resultsChan <-chan search.ExternalResultItem
	if isSingleRun {
		resultsChan = indexer.GetAllPagesFromIndex(indexerFacade, nil)
	} else {
		resultsChan = indexer.Watch(indexerFacade, nil, watchIntervalSec)
	}
	waitForResultsAndBroadcastThem(resultsChan)
}

//Reads the channel that's the result of watching an indexer.
func waitForResultsAndBroadcastThem(resultsChan <-chan search.ExternalResultItem) {
	chatMessagesChannel := make(chan bots.ChatMessage)
	token := viper.GetString("telegram_token")
	telegram, err := bots.NewTelegram(token, tgbotapi.NewBotAPI)
	if err != nil {
		fmt.Printf("Couldn't initialize telegram bot: %v", err)
		os.Exit(1)
	}
	go func() {
		_ = telegram.Run()
	}()
	go func() {
		_ = telegram.FeedBroadcast(chatMessagesChannel)
	}()
	for result := range resultsChan {
		//log.Infof("New result: %s\n", result)
		if result.IsNew() || result.IsUpdate() {
			price := result.GetField("price")
			reserved := result.GetField("reserved")
			if reserved == "true" {
				reserved = "It's currently reserved"
			} else {
				reserved = "And not reserved yet!!!"
			}
			msgText := fmt.Sprintf("I found a new property\n"+
				"[%s](%s)\n"+
				"*%s* - %s", result.Title, result.Link, price, reserved)
			message := bots.ChatMessage{Text: msgText, Banner: result.Banner}
			chatMessagesChannel <- message
			area := result.Size
			fmt.Printf("[%s][%d][%s] %s - %s\n", price, area, reserved, result.ResultItem.Title, result.Link)
		}

	}
}
