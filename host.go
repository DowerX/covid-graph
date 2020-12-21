package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

func host() error {
	fmt.Println("Started webserver!")
	r := mux.NewRouter()
	r.HandleFunc("/", handleIndex)
	r.HandleFunc("/static/{file}", handleStatic)
	err := http.ListenAndServe(conf.Address, r)
	if err != nil {
		return err
	}
	return nil
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadFile("./static/index.html")
	if err != nil {
		panic(err)
	}
	w.Write(data)
}

func handleStatic(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	data, err := ioutil.ReadFile("./static/" + vars["file"])
	if err != nil {
		panic(err)
	}
	w.Write(data)
}
