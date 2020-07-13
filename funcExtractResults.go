package ihcph

import (
	"encoding/json"
	"github.com/sp0x/ihcph/funcExtractResults"
	"github.com/sp0x/torrentd/indexer"
	"net/http"
)

type ExtractionResponse struct {
}

func TestRequest(w http.ResponseWriter, r *http.Request) {
	fnContext := funcExtractResults.Initialize()
	resultsChan := indexer.GetAllPagesFromIndex(fnContext.IndexFacade, nil)
	fnContext.Bot.BroadcastResults(resultsChan)

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(ExtractionResponse{})
}
