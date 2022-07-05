package api

import (
	
	"backend/database"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func GetList(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	if r.Method == "GET" {

		json.NewEncoder(w).Encode(database.Db)
	} else {

		var notify map[string]string = map[string]string{

			"notify" : "Not Allowed!!!",
		}

		json.NewEncoder(w).Encode(notify)
	}
}

func GetPeople(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	if r.Method == "GET" {

		body, err := ioutil.ReadAll(r.Body)

		if err != nil {

			http.Error(w, "Error reading request body", http.StatusInternalServerError)
			return
		}

		var element database.People
		err = json.Unmarshal(body, &element)

		if err != nil {

			http.Error(w, "Can’t unmarshal JSON object into struct", http.StatusInternalServerError)
			return
		}

		for i := 0; i < len(database.Db); i++ {

			if database.Db[i].Id == element.Id {

				json.NewEncoder(w).Encode(database.Db[i])
			}
		}
	} else {

		var notify map[string]string = map[string]string{

			"notify" : "Not Allowed!!!",
		}

		json.NewEncoder(w).Encode(notify)
	}
}
