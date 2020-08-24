package telegram

import (
	"cloud.google.com/go/firestore"
	"context"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/sp0x/ihcph/common"
	"github.com/sp0x/torrentd/bots"
	"github.com/sp0x/torrentd/indexer/search"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"time"
)

type BotInterface struct {
	Firestore *firestore.Client
	bot       *bots.TelegramRunner
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

func (b *BotInterface) BroadcastResults(resultsChan <-chan *search.ExternalResultItem) {
	for {
		telegram := b.bot
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
			return
		}
	}
}

func (b *BotInterface) StoreNewIntegration(integration *Integration) error {
	fbase := b.Firestore
	nsbots := fbase.Collection("bots")
	newDoc := nsbots.Doc(integration.Token)
	ctx := context.Background()
	existing, err := newDoc.Get(ctx)
	if err != nil {
		errCode := grpc.Code(err)
		if errCode != codes.NotFound {
			return err
		} else if errCode == codes.NotFound {
			integration.Id = newDoc.ID
			_, err = newDoc.Create(ctx, integration)
			return err
		}
	}
	return existing.DataTo(integration)
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

func (b *BotInterface) Initialize(token string) error {
	integration, err := b.GetBotIntegration(token)
	if err != nil {
		return err
	}
	bot, err := bots.NewTelegram(integration.Token, common.GetConfig(), tgbotapi.NewBotAPI)
	if err != nil {
		return err
	}
	b.bot = bot
	return nil
}

func (b *BotInterface) GetBot() *bots.TelegramRunner {
	return b.bot
}
