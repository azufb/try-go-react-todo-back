package main

import "net/http"

type homeHandler struct{}

// ServeHTTP implements http.Handler.
func (h *homeHandler) ServeHTTP(http.ResponseWriter, *http.Request) {
	panic("unimplemented")
}

func main() {
	mux := http.NewServeMux()

	mux.Handle("/", &homeHandler{})

	http.ListenAndServe(":8080", mux)
}
