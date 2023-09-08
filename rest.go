package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func version(w http.ResponseWriter, req *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, "{\"version\": \"0.1.mock\"	}")
}

func getById(w http.ResponseWriter, req *http.Request) {
	if req.Method != "GET" {
		w.WriteHeader(404)
		return
	}
	vars := mux.Vars(req)
	id, ok := vars["id"]
	if !ok {
		fmt.Println("id is missing in parameters")
	}
	fmt.Println(`getById() id := `, id)

	w.Header().Set("Content-Type", "application/json")
	jsonFile, err := os.Open("./mockdata/" + id + "_data.json")

	if err != nil {
		fmt.Println(err)
		fmt.Fprintf(w, "{\"error\": \"item not found\"}")
	} else {
		fmt.Println("Successfully opened mock data")
		defer jsonFile.Close()
		jByteValue, _ := ioutil.ReadAll(jsonFile)
		w.Write(jByteValue)
	}

}

func postItem(w http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" {
		w.WriteHeader(404)
		return
	}
	fmt.Fprintf(w, "Post Item Called")
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/v1/version", version)

	router.HandleFunc("/v1/item/{id}", getById)

	router.HandleFunc("/v1/item/", getById)

	router.HandleFunc("/v1/item", postItem)

	http.ListenAndServe("127.0.0.1:8090", router)

}
