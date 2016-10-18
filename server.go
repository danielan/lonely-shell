package main

import (
    "fmt"
    "io/ioutil"
    "log"
    "net/http"
)

func main() {
    http.HandleFunc("/static/", func (w http.ResponseWriter, r *http.Request) {
        if r.Method == "POST" {
            fmt.Println("Got post!")
            content, _ := ioutil.ReadAll(r.Body)
            fmt.Println(string(content))
        } else {
            fmt.Println("Got connection!")
            http.ServeFile(w, r, r.URL.Path[1:])
        }
    })
    log.Fatal(http.ListenAndServeTLS(":443", "server.crt", "server.key", nil))
}
