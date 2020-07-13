package ihcph

import (
	"encoding/json"
	"github.com/sp0x/ihcph/funcExtractResults"
	"github.com/sp0x/torrentd/indexer"
	"net/http"
)

type ExtractionRequest struct {
	BotToken string `json:"bot_id"`
}

type ExtractionResponse struct {
}

func ExtractResults(w http.ResponseWriter, r *http.Request) {
	extractionRequest := ExtractionRequest{}
	err := json.NewDecoder(r.Body).Decode(&extractionRequest)
	if err != nil {
		http.Error(w, "bad request", 400)
		return
	}
	fnContext := funcExtractResults.Initialize()
	botIntegration, err := fnContext.Bot.GetBotIntegration(extractionRequest.BotToken)
	if err != nil {
		http.Error(w, "error fetching bot", 500)
		return
	}
	resultsChan := indexer.GetAllPagesFromIndex(fnContext.IndexFacade, nil)
	fnContext.Bot.BroadcastResults(botIntegration, resultsChan)

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(ExtractionResponse{})
}
