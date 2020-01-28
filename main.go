package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"os"

	"github.com/fatih/color"
)

func main() {
	port := "8090"
	if value, ok := os.LookupEnv("PORT"); ok {
        port =  value
    }

	http.HandleFunc("/", handler)
	color.Cyan("--> Go Server Listening @ "+port+"!!")
	log.Fatal(http.ListenAndServe(":"+port, nil))

}

func handler(w http.ResponseWriter, r *http.Request) {
	urlPath := strings.ToLower(r.URL.Path)
	if urlPath == "/go" {
		fmt.Fprintln(w, ":)")
		return
	} 

	redirectTO := "https://vithalreddy.com"

	val, ok := links[urlPath]
	if ok {
		redirectTO = val
	}

	fmt.Printf("::Redircted:: %s -> %s \n", urlPath, val)
	http.Redirect(w, r, redirectTO, 302)
	

}


var links = map[string]string{
	"/github": "https://github.com/vithalreddy",
	"/gitlab": "https://gitlab.com/vithalreddy",
	"/v": "https://vithalreddy.com",
	"/fame": "https://stackfame.com",
	"/fame-admin": "https://stackfame.com/wp-admin/edit.php",
	"/notes": "https://github.com/vithalreddy/notes",
	"/cv":    "https://drive.google.com/file/d/109fRG2KxV8WWjPDxNiLpj4TbYK7DxFGf/view",
}
