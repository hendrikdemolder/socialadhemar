package hello

import (
    "fmt"
    "http"
)

func init() {
    http.HandleFunc("/bla", handler)
}

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "Hello world2!")
}
