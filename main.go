package main

import (
	"net/http"
	"fmt"
)

func main() {
	fmt.Println("Hello World !")
  var chttp = http.NewServeMux()

  chttp.Handle("/", http.FileServer(http.Dir("./")))

  http.HandleFunc("/", HomeHandler) // homepage
  http.ListenAndServe(":8080", nil)

  func HomeHandler(w http.ResponseWriter, r *http.Request) {
    if (strings.Contains(r.URL.Path, ".")) {
        chttp.ServeHTTP(w, r)
    } else {
        http.Error(w, "Not found", 404)
    }
  }
}

