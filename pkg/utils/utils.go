package utils

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func ParseBody(r *http.Request, x interface{}) {
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatal(err.Error())
		return
	}

	_ = json.Unmarshal([]byte(body), x)
}

func GetParamID(r *http.Request, param string) int {
	params := mux.Vars(r)
	param_id := params[param]

	id, err := strconv.Atoi(param_id)

	if err != nil {
		log.Fatal("error while parsing param")
	}

	return id
}
