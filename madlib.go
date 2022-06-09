package main

import (
    "fmt"
    "net/http"
    "io"
)

var sentence = `It was a %s day. I went downstairs to see if I could %s dinner. I asked, "Does the stew need fresh %s?"`

func madlib(w http.ResponseWriter, req *http.Request) {
    response := fmt.Sprintf(sentence,  randomWord("/adjective"), randomWord("/verb"), randomWord("/noun"))
    fmt.Fprintf(w, response + "\n")
}

func main() {
    http.HandleFunc("/madlib", madlib)

    http.ListenAndServe(":80", nil)
}


func randomWord(partOfSpeech string) string {
    resp, err := http.Get("https://reminiscent-steady-albertosaurus.glitch.me" + partOfSpeech)
    if err != nil {
        panic(err)
    }

    defer resp.Body.Close()
    body, err := io.ReadAll(resp.Body)
    if err != nil {
        panic(err)
    }

    return string(body)
}
