package main

import (
  "github.com/higashi000/vimrc-autogenerate/post"
  "github.com/higashi000/vimrc-autogenerate/getData"
)

func main() {
  vimrcOption := getData.GetData()
  post.Generate(vimrcOption)
}
