package routes

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func JsonResponse(w http.ResponseWriter, data interface{}) {

	jsonData, err := json.Marshal(data)

	w.Header().Add("Content-Type", "application/json")

	if err != nil {

		fmt.Fprintf(w, `{"Message":"Json Decode Error"}`)

		return

	}

	w.Write(jsonData)
}
