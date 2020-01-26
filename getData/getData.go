package getData

import (
  "syscall/js"
  "strconv"
  "log"
  "github.com/higashi000/vimrc-autogenerate/post"
)

func GetData() post.VimrcOption {
  document := js.Global().Get("document")

  userName := document.Call("getElementById", "username")
  colorscheme := document.Call("getElementById", "colorscheme")
  indent := document.Call("getElementById", "indent")

  orderVimrc := post.VimrcOption{}
  orderVimrc.UserName = userName.Get("value").String()
  indentValue, err := strconv.Atoi(indent.Get("value").String())
  if err != nil {
    log.Println(err)
  }
  orderVimrc.Indent = indentValue
  orderVimrc.ColorScheme = colorscheme.Get("value").String()

  return orderVimrc
}
