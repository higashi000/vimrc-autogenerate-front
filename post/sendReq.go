package post

import (
  "net/http"
  "encoding/json"
  "bytes"
  "io/ioutil"
  "fmt"
  "log"
  "github.com/google/uuid"
)

type Vimrc struct {
  Body string `json:"body"`
}

type VimrcOption struct {
  Uuid string `json:"uuid"`
  UserName string `json:"user"`
  Indent int `json:"indent"`
  ColorScheme string `json:"colorscheme"`
}

func Generate(orderVimrc VimrcOption) {
  url := "http://localhost:5000/generate"

  u, err := uuid.NewRandom()
  if err != nil {
    fmt.Println(err)
    return
  }

  orderVimrc.Uuid = u.String()

  sendData, err := json.Marshal(orderVimrc)
  if err != nil {
    log.Println("error struct to json")
  }

  req, err := http.NewRequest(
    "POST",
    url,
    bytes.NewBuffer(sendData),
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

  var vimrc Vimrc
  err = json.Unmarshal(body, &vimrc)
  if err != nil {
    log.Println("error json")
    return
  }

  fmt.Println(vimrc.Body)
}
