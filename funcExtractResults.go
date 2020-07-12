package ihcph

import (
	"encoding/json"
	"github.com/sp0x/ihcph/function"
	"github.com/sp0x/torrentd/indexer"
	"net/http"
)

type Result struct {
	Code    int
	Message string
	Token   string
}

type Body struct {
	Message string
}

func TestRequest(w http.ResponseWriter, r *http.Request) {
	fnContext := function.Initialize()
	resultsChan := indexer.GetAllPagesFromIndex(fnContext.IndexFacade, nil)
	fnContext.Bot.BroadcastResults(resultsChan)
	body := Body{}
	body.Message = "Scanned for new results."
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(Result{
		200,
		body.Message,
		"",
	})
}
