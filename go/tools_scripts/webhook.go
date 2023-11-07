package tools_scripts

import (
	"encoding/json"
	"log"
	"net/http"
)

func StartWebhook() error {
	s := http.NewServeMux()

	s.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		decoder := json.NewDecoder(r.Body)
		data := map[string]any{}
		if err := decoder.Decode(&data); err != nil {
			log.Println(err)
		} else {
			log.Println(data)
		}
		w.WriteHeader(http.StatusOK)
	})

	return http.ListenAndServe(":8888", s)
}
