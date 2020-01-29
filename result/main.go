package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"syscall/js"
)

type Vimrc struct {
	Name string `json:"name"`
	Body string `json:"body"`
}

func main() {
	document := js.Global().Get("document")
	vimrcUUID := document.Call("getElementById", "vimrcUUID")
	url := "http://localhost:5000/result/" + vimrcUUID.Get("value").String()

	req, err := http.NewRequest(
		"GET",
		url,
		nil,
	)
	if err != nil {
		return
	}

	req.Header.Add("Content-Type", "application/json")

	client := new(http.Client)

	resp, err := client.Do(req)
	if err != nil {
		return
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	var vimrc Vimrc

	err = json.Unmarshal(body, &vimrc)
	if err != nil {
		log.Println("error json")
		return
	}

	vimrcEle := document.Call("getElementById", "vimrc")
	vimrcEle.Set("innerText", vimrc.Body)

	user := document.Call("getElementById", "user")
	user.Set("innerText", "このvimrcを作ったのは"+vimrc.Name)

	fmt.Println(vimrc.Body)
}
