package Serverless

import (
	"encoding/json"
	"net/http"
	"time"
)

type TimeResponse struct {
	Time string `json:"time"`
}

func CurrentTime(w http.ResponseWriter, r *http.Request) {
	time := time.Now()

	res := TimeResponse{Time: time.String()}

	if err := json.NewEncoder(w).Encode(res); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
