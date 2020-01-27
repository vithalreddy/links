package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/fatih/color"
)

func main() {

	http.HandleFunc("/", handler)
	color.Cyan("Go Server Listening @ 8090!")
	log.Fatal(http.ListenAndServe(":8090", nil))

}

func handler(w http.ResponseWriter, r *http.Request) {
	urlPath := strings.ToLower(r.URL.Path)
	if urlPath == "/go" {
		fmt.Fprintln(w, ":)")
		return
	} else {
		redirectTO := "https://vithalreddy.com"

		val, ok := links[urlPath]
		if ok {
			redirectTO = val
		}

		fmt.Printf("::Redircted:: %s -> %s \n", urlPath, val)
		http.Redirect(w, r, redirectTO, 302)
	}

}

var links = map[string]string{
	"/notes": "https://github.com/vithalreddy/notes",
	"/cv":    "https://drive.google.com/file/d/109fRG2KxV8WWjPDxNiLpj4TbYK7DxFGf/view",
}
