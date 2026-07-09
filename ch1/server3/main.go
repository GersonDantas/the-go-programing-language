package main

import (
	"log"
	"net/http"
	"strconv"
	"sync"

	"github.com/GersonDantas/gopl/ch1/lissajous"
)

var mu sync.Mutex
var count int

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		cyclesStr := r.URL.Query().Get("cycles")
		log.Printf("Received request with cycles: %s", cyclesStr)
		var cycles int
		if cyclesStr != "" {
			mu.Lock()
			c, err := strconv.Atoi(cyclesStr)
			if err != nil {
				http.Error(w, "Invalid cycles value", http.StatusBadRequest)
				mu.Unlock()
				return
			}
			cycles = c
			mu.Unlock()
		}
		lissajous.Lissajous(w, cycles)
	})
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// func handler(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "%s %s %s", r.Method, r.URL.Path, r.Proto)
// 	for k, v := range r.Header {
// 		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
// 	}
// 	fmt.Fprintf(w, "Host = %q\n", r.Host)
// 	fmt.Fprintf(w, "RemoteAddr = %q\n", r.RemoteAddr)
// 	if err := r.ParseForm(); err != nil {
// 		log.Print(err)
// 	}
// 	for k, v := range r.Form {
// 		fmt.Fprintf(w, "Form[%q] = %q\n", k, v)
// 	}
// }
