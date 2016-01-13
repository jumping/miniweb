package miniweb

import(
	"net/http"
	"fmt"
)

func Run(host string, mux *Router) {
	fmt.Println("Start  services at ", host, ".....")
	http.ListenAndServe(host, mux)
}