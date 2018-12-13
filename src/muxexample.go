package src

import (
    "fmt"
    "log"
    "net/http"
)

type mux struct {
}

func (m *mux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    switch r.URL.Path {
    case "/door":
        func(w http.ResponseWriter, r *http.Request) {
            fmt.Fprintln(w, "You have opened.")
        }(w, r)
    default:
        http.NotFound(w, r)
    }
}

func StartMuxSever() {
    m := &mux{}

    log.Println("We are ready for guests!")
    log.Fatal(http.ListenAndServe("127.0.0.1:8890", m))
}
