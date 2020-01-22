package main

import (
  "github.com/higashi000/vimrc-autogenerate/post"
)

type Data struct {
  Data int `json:"data"`
}

func main() {
  post.Generate(10, "af;kja")
}
