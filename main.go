package main

import (
    "encoding/json"
    "log"
    "net/http"
)

type Question struct {

 Sentence string `json:"sentence"`
 Word string `json:"word"`
 Option [] string `json:"options"`
 Answer string `json:"answer"`
}

func main() {
    //API endpoint
    http.HandleFunc("/api/quiz", func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Access-Control-Allow-Origin", "*")
        w.Header().Set("Content-Type", "application/json")

        q := Question{
            Sentence: "She quickly finished the report.",
            Word: "quickly",
            Option: []string{"Adjective", "Adverb", "verb", "Noun"},
            Answer: "Adverb",
        }
        json.NewEncoder(w).Encode(q)
    })

    // Serve HTML
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        http.ServeFile(w, r, "index.html")
    })
    log.Println("Server running on http://localhost:8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
