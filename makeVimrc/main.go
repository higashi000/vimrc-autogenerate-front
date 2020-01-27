package main

import (
  "github.com/higashi000/vimrc-autogenerate/makeVimrc/post"
  "github.com/higashi000/vimrc-autogenerate/makeVimrc/getData"
)

func main() {
  vimrcOption := getData.GetData()
  post.Generate(vimrcOption)
}
