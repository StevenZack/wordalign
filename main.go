package main

import (
	"fmt"
	"github.com/StevenZack/openurl"
	"io"
	"net/http"
	"os"
	"os/exec"
)

func main() {
	port := "10290"
	http.HandleFunc("/", home)
	openurl.Open("http://127.0.0.1:" + port)
	e := http.ListenAndServe(":"+port, nil)
	if e != nil {
		fmt.Println(e)
		return
	}
}
func home(w http.ResponseWriter, r *http.Request) {
	f, e := os.OpenFile("index.html", os.O_RDONLY, 0644)
	if e != nil {
		fmt.Println(e)
		return
	}
	io.Copy(w, f)
	f.Close()
}
