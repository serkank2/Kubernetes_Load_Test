package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type handler struct{}

func (h *handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	temp := 1
	for i := 2; i < 2000000; i++ {

		fmt.Println(temp * i)
	}
	w.Write([]byte(fmt.Sprint(temp)))
}

func main() {
	mux := mux.NewRouter()
	mux.Handle("/", &handler{})
	fmt.Println("Server Started")
	http.ListenAndServe(":80", mux)

}
