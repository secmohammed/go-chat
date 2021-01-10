package main

import (
    "fmt"
    "log"
    "net/http"

    "github.com/secmohammed/private-chat-go/internal"
)

func main() {
    router := internal.NewRouter()
    router.Handle("channel add", internal.AddChannel)

    http.Handle("/", router)
    if err := http.ListenAndServe(":4000", nil); err != nil {

        log.Fatal(err)
    }
    fmt.Println("hello")

}
