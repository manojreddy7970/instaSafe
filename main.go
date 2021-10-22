package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/manojreddy7970/instaSafe/tickethandler"
)

// The new router function creates the router and
// returns it to us. We can now use this function
// to instantiate and test the router outside of the main function
func newRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/hello", handler).Methods("GET")

	staticFileDirectory := http.Dir("./assets/")

	staticFileHandler := http.StripPrefix("/assets/", http.FileServer(staticFileDirectory))
	r.PathPrefix("/assets/").Handler(staticFileHandler).Methods("GET")

	r.HandleFunc("/platformTicket", tickethandler.GetBalancePFTHandler).Methods("GET")
	r.HandleFunc("/platformTicket", tickethandler.EvaluatePFTHandler).Methods("POST")
	return r
}

func main() {
	r := newRouter()
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		panic(err.Error())
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}