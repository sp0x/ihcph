package ihcph

import (
	"github.com/gorilla/mux"
	"github.com/steinfletcher/apitest"
	"net/http"
	"testing"
)

func Test_TestFunction(t *testing.T) {
	r := mux.NewRouter()
	r.HandleFunc("/funcExtractResults/method", TestRequest)
	apitest.New().
		Handler(r).
		Get("/funcExtractResults/method").
		Expect(t).
		Status(http.StatusOK).
		End()

}
