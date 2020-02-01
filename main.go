package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/fatih/color"
)

var links = map[string]string{}

func main() {
	go heartBeat()

	port := "8090"
	if value, ok := os.LookupEnv("PORT"); ok {
		port = value
	}

	jsonData, _ := ioutil.ReadFile("./data.json")
	json.Unmarshal(jsonData, &links)

	http.HandleFunc("/", handler)
	color.Cyan("--> Go Server Listening @ " + port + "!!")
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

func heartBeat() {
	for range time.Tick(time.Second * 1) {
		res, err := http.Get("https://r.reddy.is/go")
		fmt.Println("HeartBeat ::")
		if err != nil {
			fmt.Print(err.Error())
		}
		responseData, err := ioutil.ReadAll(res.Body)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(responseData))

	}
}
