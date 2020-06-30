package function

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

type Result struct {
	Code    int
	Message string
	Token   string
}

type Body struct {
	Message string
}

const (
	signingKey = "Anekpattanakij-Golang-CD"
)

func TestRequest(w http.ResponseWriter, r *http.Request) {
	data, _ := ioutil.ReadAll(r.Body)
	body := Body{}
	body.Message = string(data)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"Message": body.Message,
		"exp":     time.Now().Add(time.Hour * 72).Unix(),
	})
	tokenString, err := token.SignedString([]byte(signingKey))

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(Result{
			500,
			err.Error(),
			"",
		})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(Result{
		200,
		body.Message,
		tokenString,
	})
	return

}
