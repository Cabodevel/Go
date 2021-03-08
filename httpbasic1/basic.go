package httpBasic

import (
	"encoding/json"
	"log"
	"net/http"
)

//StartServer starts the server
func StartServer() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		names := r.URL.Query()["name"]
		var name string
		if len(names) == 1 {
			name = names[0]
		}

		m := map[string]string{"name": name}
		enc := json.NewEncoder(w)
		enc.Encode(m)

		w.Write([]byte("Hola " + name))
	})

	err := http.ListenAndServe(":5001", nil)

	if err != nil {
		log.Fatal(err)
	}
}
