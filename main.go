package main

import (
	"log"
	"net/http"

	"github.com/Ink-33/httphere/internal/base"
)

type loghandlers struct {
	funcs []http.Handler
}

func main() {
	log.SetPrefix("httphere")
	log.Println("version", base.Version)

	hs := &loghandlers{}
	hs.funcs = append(hs.funcs, http.FileServer(http.Dir(".")))

	log.Println("Listening on 0.0.0.0:8080")
	_ = http.ListenAndServe(":8080", hs)
}

func (h *loghandlers) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	log.Printf("%s\trequest\t%s\n", req.RemoteAddr, req.URL)
	for i := range h.funcs {
		h.funcs[i].ServeHTTP(resp, req)
	}
}
