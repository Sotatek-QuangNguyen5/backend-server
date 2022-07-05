package api

import (

	"io/ioutil"
	"net/http"
	"encoding/json"
	"backend/database"
)

func Edit(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	if r.Method == "PUT" {

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
		
		for i := 0; i <  len(database.Db); i++ {

			if database.Db[i].Id == element.Id {

				database.Db[i] = element
			}
		}

		var notify map[string]string = map[string]string{

			"notify" : "Edit Successed!!",
		}
		json.NewEncoder(w).Encode(notify)

	} else {

		var notify map[string]string = map[string]string{

			"notify" : "Not Allowed!!",
		}
		json.NewEncoder(w).Encode(notify)
	}
}
