package main

import (
        "fmt"
        "net/http"

        "github.com/gorilla/mux"
)

func main() {
        mainRouter := mux.NewRouter().StrictSlash(true)
        mainRouter.HandleFunc("/Hello", GreetingHandler)
        err := http.ListenAndServe(":8080", mainRouter)
        if err != nil {
                fmt.Println("Something is wrong : " + err.Error())
        }
}

func GreetingHandler(w http.ResponseWriter, r *http.Request){
        w.WriteHeader(http.StatusOK)
        w.Header().Set("Content-Type", "text/json")
        w.Write([]byte("Hello from the other side!!!"))
}