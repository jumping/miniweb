package miniweb

import(
	"net/http"
)

func Run(host string, mux *Router) {
	http.ListenAndServe(host, mux)
}