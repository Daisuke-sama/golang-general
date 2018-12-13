package src

import (
    "fmt"
    "log"
    "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Testing http.HandleFunc().")
}

func ServeIt() {
    http.HandleFunc("/test", handler)

    log.Println("Server is switching on *** ")
    log.Fatal(http.ListenAndServe("127.0.0.1:8889", nil))
}

type inviteHandler struct{}

func (ph *inviteHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Here we are.")
}

func ServeItMature() {
    gh := &inviteHandler{}
    http.Handle("/signin/", gh)

    ServeIt()
}

func ServeItGreat() {
    sh := newSentenceHandler()
    http.Handle("/sentences/", sh)

    ServeItMature()

    StartMuxSever()
}

func ServeWithMultiplexer() {
    m := &mux{}
    http.Handle("/", m)

    ServeItGreat()
}
