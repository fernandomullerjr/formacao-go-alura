package main

import "net/http"
import "fmt"

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
        fmt.Fprintf(w, "<h1>Hello, World!</h1>")
    })

    http.ListenAndServe(":80", nil)
}
