package main

import (
    "fmt"
    "github.com/jboursiquot/things"
    "gotraining/GSLTraining/src"
)

func main() {
    fmt.Println(things.Thing{Name: "Thing one."})
    fmt.Println(things.Thing{Name: "Thing two."})

    src.ServeWithMultiplexer()
}
