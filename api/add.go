package api

import (
	"backend/database"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func Add(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	if r.Method == "POST" {

		body, err := ioutil.ReadAll(r.Body)

		if err != nil {

			http.Error(w, "Error reading request body", http.StatusInternalServerError)
		}

		var element database.People
		json.Unmarshal(body, &element)
		database.Db = append(database.Db, element)
		var notify map[string]string = map[string]string{
			"notify": "People are added !!!",
		}
		json.NewEncoder(w).Encode(notify)

	} else {

		var notify map[string]string = map[string]string{

			"notify": "Not Allowed!!!",
		}
		json.NewEncoder(w).Encode(notify)
	}
}
