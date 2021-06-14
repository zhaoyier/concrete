package utils

import (
	"fmt"
	"io"
	"net/http"
)

func SayHello(w http.ResponseWriter, r *http.Request) {
	fmt.Println("==>>TODO SayHello 001")
	io.WriteString(w, "Welcome to my blog site")
}
