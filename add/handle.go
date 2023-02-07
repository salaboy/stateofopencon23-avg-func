package function

import (
	"context"
	"encoding/json"
	"net/http"

	dapr "github.com/dapr/go-sdk/client"
)

var STATE_STORE_NAME = "statestore"
var daprClient dapr.Client

type MyValues struct {
	Values []string
}

// Handle an HTTP Request.
func Handle(ctx context.Context, res http.ResponseWriter, req *http.Request) {

	daprClient, err := dapr.NewClient()
	if err != nil {
		panic(err)
	}

	value := req.URL.Query().Get("value")

	result, _ := daprClient.GetState(ctx, STATE_STORE_NAME, "values", nil)
	myValues := MyValues{}
	if result.Value != nil {
		json.Unmarshal(result.Value, &myValues)
	}

	if myValues.Values == nil || len(myValues.Values) == 0 {
		myValues.Values = []string{value}
	} else {
		myValues.Values = append(myValues.Values, value)
	}

	jsonData, err := json.Marshal(myValues)

	err = daprClient.SaveState(ctx, STATE_STORE_NAME, "values", jsonData, nil)
	if err != nil {
		panic(err)
	}

	respondWithJSON(res, http.StatusOK, myValues)

}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
