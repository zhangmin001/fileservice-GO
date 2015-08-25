package main
import (
    "fmt"
    "net/http"
)

func main() {
    http.HandleFunc("/", homeHandler)
    http.HandleFunc("/static/", func(w http.ResponseWriter, r *http.Request) {

        http.ServeFile(w, r, r.URL.Path)
    })
    panic(http.ListenAndServe(":17901", nil))
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Welcome Home! Show Image!")
}