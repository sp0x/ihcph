package telegram

import (
	"cloud.google.com/go/firestore"
	"context"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/sp0x/ihcph/common"
	"github.com/sp0x/torrentd/bots"
	"github.com/sp0x/torrentd/config"
	"github.com/sp0x/torrentd/indexer/search"
	"time"
)

type BotInterface struct {
	Firestore *firestore.Client
}

func NewBotInterface() *BotInterface {
	b := &BotInterface{}
	fstore, err := common.NewFirebaseFromEnv()
	if err != nil {
		panic(err)
	}
	b.Firestore = fstore
	return b
}

func (b *BotInterface) BroadcastResults(integration *Integration, resultsChan <-chan *search.ExternalResultItem) {
	for {
		telegram, err := bots.NewTelegram(integration.Token, &config.ViperConfig{}, tgbotapi.NewBotAPI)
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
				telegram.Broadcast(message)
			}
		case <-time.After(10 * time.Second):
			fmt.Printf("Timed out waiting for result")
			break
		}
	}
}

func (b *BotInterface) StoreNewIntegration(integration *Integration) error {
	fbase := b.Firestore
	nsbots := fbase.Collection("bots")
	newDoc := nsbots.Doc(integration.Token)
	ctx := context.Background()
	existing, err := newDoc.Get(ctx)
	if existing != nil {
		return nil
	}
	integration.Id = newDoc.ID
	_, err = newDoc.Create(ctx, integration)
	return err
}

func (b *BotInterface) GetBotIntegration(token string) (*Integration, error) {
	fbase := b.Firestore
	nsbots := fbase.Collection("bots")
	newDoc := nsbots.Doc(token)
	ctx := context.Background()
	existing, err := newDoc.Get(ctx)
	if existing == nil {
		return nil, err
	}
	ign := Integration{}
	err = existing.DataTo(&ign)
	if err != nil {
		return nil, err
	}
	return &ign, nil
}
