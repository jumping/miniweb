package miniweb

import(
	"fmt"
)

// 404 Not Found
func status404(res Resource) {
	fmt.Fprintf(res.W, "<b>404 Not Found</b>")
}

// 405 Method Not Allowed
func status405(res Resource) {
	fmt.Fprintf(res.W, "<b>405 Method Not Allowed</b>")
}

// 400 Bad Request
func status400(res Resource) {
	fmt.Fprintf(res.W, "<b>400 Bad Request</b>")
}