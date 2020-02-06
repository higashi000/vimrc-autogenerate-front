package getList

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type Data struct {
	Name string `json:"name"`
	Uuid string `json:"uuid"`
}

type VimrcList struct {
	Inform []struct {
		Name string `json:"name"`
		Uuid string `json:"uuid"`
	} `json:"inform"`
}

func GetList() VimrcList {
	url := "http://localhost:5000/list"

	req, err := http.NewRequest(
		"GET",
		url,
		nil,
	)
	if err != nil {
		log.Println(err)
	}

	req.Header.Add("Content-Type", "application/json")

	client := new(http.Client)

	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}

	var data VimrcList

	err = json.Unmarshal(body, &data)
	if err != nil {
		log.Println(err)
	}

	resp.Body.Close()

	return data
}
