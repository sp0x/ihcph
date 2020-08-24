package ihcph

import (
	"encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/sp0x/ihcph/funcBotIntegration"
	"github.com/sp0x/ihcph/telegram"
	"net/http"
)

type NewBotRequest struct {
	Token string `json:"token"`
}

type NewBotRequestResponse struct {
	Id string `json:"id"`
}

func NewBotIntegration(w http.ResponseWriter, r *http.Request) {
	_, err := funcBotIntegration.Initialize()
	if err != nil {
		http.Error(w, fmt.Sprintf("couldn't initialize: %v", err), 500)
		return
	}
	newBot := NewBotRequest{}
	err = json.NewDecoder(r.Body).Decode(&newBot)
	if err != nil {
		http.Error(w, "Couldn't decode body", 400)
		return
	}
	newIntegration := &telegram.Integration{
		Token: newBot.Token,
	}
	bot := telegram.NewBotInterface()
	err = bot.StoreNewIntegration(newIntegration)
	if err != nil {
		log.Errorf("error while creating integration: %v", err)
		http.Error(w, "Couldn't save integration", 500)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(NewBotRequestResponse{
		newIntegration.Id,
	})
}
