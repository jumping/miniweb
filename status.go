package miniweb

import(
	"fmt"
)

type Status struct {}

// 404 Not Found
func (s Status)Status404(res Resource) {
	fmt.Fprintf(res.W, "<b>404 Not Found</b>")
}

// 405 Method Not Allowed
func (s Status)Status405(res Resource) {
	fmt.Fprintf(res.W, "<b>405 Method Not Allowed</b>")
}

// 400 Bad Request
func (s Status)Status400(res Resource) {
	fmt.Fprintf(res.W, "<b>400 Bad Request</b>")
}