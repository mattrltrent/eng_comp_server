package lib

import (
	"encoding/json"
	"net/http"
)

func Err(w http.ResponseWriter, msg string) {
	j, err := json.Marshal(map[string]string{"error": msg})
	if err != nil {
		panic(err)
	}
	w.Write(j)
}
