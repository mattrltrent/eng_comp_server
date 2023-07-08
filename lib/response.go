package lib

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

func Err(w http.ResponseWriter, msg string) {
	fmt.Fprint(os.Stderr, msg)
	w.Header().Add("content-type", "application/json")
	j, err := json.Marshal(map[string]string{"error": msg})
	if err != nil {
		panic(err)
	}
	w.Write(j)
}

func JSON(w http.ResponseWriter, d interface{}) {
	w.Header().Add("content-type", "application/json")
	if err := json.NewEncoder(w).Encode(&d); err != nil {
		panic(err)
	}
}
