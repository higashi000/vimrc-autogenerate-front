package main

import (
  "net/http"
	"encoding/json"
  "bytes"
  "io/ioutil"
  "fmt"
  "log"
)

type Data struct {
  Data int `json:"data"`
}

func main() {
  generate()
}

func generate() {
  url := "http://localhost:5000/generate"

  sendData := `{"data":10}`

  req, err := http.NewRequest(
    "POST",
    url,
    bytes.NewBuffer([]byte(sendData)),
  )
  if err != nil {
    log.Println("error req")
    return
  }

  req.Header.Add("Content-Type", "application/json")

  client := new(http.Client)
  resp, err := client.Do(req)
  if err != nil {
    log.Println("error resp")
    log.Println(err)
    return
  }
  defer resp.Body.Close()

  body, err := ioutil.ReadAll(resp.Body)

  fmt.Println(string(body))

  var data Data
  err = json.Unmarshal(body, &data)
  if err != nil {
    log.Println("error json")
    return
  }

  fmt.Println(data.Data)
}
