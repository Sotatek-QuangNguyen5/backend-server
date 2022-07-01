package api

import (

	"backend/database"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func Delete(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	if r.Method == "DELETE" {

		body, err := ioutil.ReadAll(r.Body)

		if err != nil {

			http.Error(w, "Error reading request body", http.StatusInternalServerError)
		}

		var element database.People
		json.Unmarshal(body, &element)
		for i := 0; i < len(database.Db); i++ {

			if database.Db[i].Id == element.Id {

				database.Db = append(database.Db[:i], database.Db[i+1:]...)
				break
			}
		}
		var notify map[string]string = map[string]string{

			"notify" : "Delete Successed!!!",
		}
		json.NewEncoder(w).Encode(notify)

	} else {

		var notify map[string]string = map[string]string{

			"notify" : "Not Allowed!!!",
		}
		json.NewEncoder(w).Encode(notify)
	}
}
