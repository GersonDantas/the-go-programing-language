package main

import (
	"log"
	"math/rand"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"time"

	"github.com/GersonDantas/gopl/ch1/lissajous"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	if len(os.Args) > 1 && os.Args[1] == "web" {
		handler := func(w http.ResponseWriter, r *http.Request) {
			lissajous.Lissajous(w, 0)
		}
		http.HandleFunc("/", handler)
		log.Fatal(http.ListenAndServe("localhost:8000", nil))
		return
	}
	if len(os.Args) > 1 && os.Args[1] == "file" {
		gf, err := os.Create(outputPath("out.gif"))
		if err != nil {
			log.Fatal(err)
		}
		lissajous.Lissajous(gf, 0)
		gf.Close()
		return
	}
	lissajous.Lissajous(os.Stdout, 0)
}

func outputPath(name string) string {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		return name
	}
	return filepath.Join(filepath.Dir(filename), name)
}
