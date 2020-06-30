package ihcph

import (
	"fmt"
	"github.com/sp0x/torrentd/bots"
	"github.com/sp0x/torrentd/indexer/search"
	"time"
)

type BotInterface struct {
	Telegram *bots.TelegramRunner
}

func (b *BotInterface) BroadcastResults(resultsChan <-chan *search.ExternalResultItem) {
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
				b.Telegram.Broadcast(message)
			}
		case <-time.After(10 * time.Second):
			fmt.Printf("Timed out waiting for result")
			break
		}
	}
}
