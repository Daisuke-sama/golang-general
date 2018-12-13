package src

import (
    "errors"
    "fmt"
    "net/http"
    "strconv"
)

type sentence struct {
    id  int
    val string
}

// The structure with a range of sentences
type sentencesHandler struct {
    sentences []sentence
}

// The constructor.
func newSentenceHandler() *sentencesHandler {
    return &sentencesHandler{
        sentences: []sentence{
            {id: 1, val: "The first one."},
            {id: 2, val: "The first second."},
            {id: 3, val: "The first third."},
            {id: 4, val: "The first forth."},
            {id: 5, val: "The first fifth."},
            {id: 6, val: "The first sixth."},
            {id: 7, val: "The first seventh."},
        },
    }
}

// The expanding method for the sentenceHandler, which makes
// the sentenceHandler implementing a special http interface.
func (sh *sentencesHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    id, err := strconv.Atoi(r.URL.Path[len("/sentences/"):])
    if err != nil {
        w.WriteHeader(http.StatusBadRequest)
        return
    }

    s, err := sh.search(id)
    if err == errUnknownSentence {
        http.Error(w, errUnknownSentence.Error(), http.StatusNotFound)
        return
    }

    fmt.Fprintln(w, s.val)
}

var errUnknownSentence = errors.New("Your sentence has not been found.")

func (sh *sentencesHandler) search(id int) (*sentence, error) {
    for _, s := range sh.sentences {
        if id == s.id {
            return &s, nil
        }
    }

    return nil, errUnknownSentence
}
