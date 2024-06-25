package main

import (
    "fmt"
    "net/http"
)

func main() {
    http.HandleFunc("/unsplash", servesplash)
    http.Handle("/", http.FileServer(http.Dir(".")))

    fmt.Println("Server is running on http://localhost:8080")
    err := http.ListenAndServe(":8080", nil)
    if err != nil {
        fmt.Printf("Server failed: %v\n", err)
    }
}

func servesplash(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, "Index.html")
}
