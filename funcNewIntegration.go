package ihcph

import (
	"encoding/json"
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
	newBot := NewBotRequest{}
	err := json.NewDecoder(r.Body).Decode(&newBot)
	if err != nil {
		http.Error(w, "Couldn't decode body", 400)
		return
	}
	ctxt := funcBotIntegration.Initialize()
	newIntegration := &telegram.Integration{
		Token: newBot.Token,
	}
	err = ctxt.Bot.StoreNewIntegration(newIntegration)
	if err != nil {
		http.Error(w, "Couldn't save integration", 500)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(NewBotRequestResponse{
		newIntegration.Id,
	})
}
