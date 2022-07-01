package main

import (
	"backend/api"
	"fmt"
	"log"
	"net/http"
)

const PORT = ":5050"
//a
func main() {

	http.HandleFunc("/add", api.Add)
	http.HandleFunc("/edit", api.Edit)
	http.HandleFunc("/getList", api.GetList)
	http.HandleFunc("/getPeople", api.GetPeople)
	http.HandleFunc("/delete", api.Delete)

	fmt.Println("Server is running on port ", PORT)
	log.Fatal(http.ListenAndServe(PORT, nil))
}
