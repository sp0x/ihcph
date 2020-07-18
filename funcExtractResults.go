package ihcph

import (
	"encoding/json"
	"github.com/sp0x/ihcph/funcExtractResults"
	"github.com/sp0x/ihcph/telegram"
	"github.com/sp0x/torrentd/indexer"
	"github.com/sp0x/torrentd/indexer/search"
	"net/http"
)

type ExtractionRequest struct {
	BotToken string `json:"bot_id"`
}

type ExtractionResponse struct {
	Items []string `json:"items"`
}

func ExtractResults(w http.ResponseWriter, r *http.Request) {
	extractionRequest := ExtractionRequest{}
	err := json.NewDecoder(r.Body).Decode(&extractionRequest)
	if err != nil {
		http.Error(w, "bad request", 400)
		return
	}
	fnContext := funcExtractResults.Initialize()
	botInterface := telegram.NewBotInterface()
	err = botInterface.Initialize(extractionRequest.BotToken)
	if err != nil {
		http.Error(w, "error fetching bot", 500)
		return
	}
	resultsChan := indexer.GetAllPagesFromIndex(fnContext.IndexFacade, nil)
	proxyChan := make(chan *search.ExternalResultItem)
	go botInterface.BroadcastResults(proxyChan)
	response := ExtractionResponse{Items: []string{}}
	for result := range resultsChan {
		proxyChan <- result
		if result == nil {
			break
		}
		response.Items = append(response.Items, result.Title)
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(response)
}
