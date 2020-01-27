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
  indentSettingNum := document.Call("getElementById", "indentSettingNum")

  orderVimrc := post.VimrcOption{}
  orderVimrc.UserName = userName.Get("value").String()
  indentValue, err := strconv.Atoi(indent.Get("value").String())
  if err != nil {
    log.Println(err)
  }
  orderVimrc.Indent = indentValue
  orderVimrc.ColorScheme = colorscheme.Get("value").String()

  num, err := strconv.Atoi(indentSettingNum.Get("innerText").String())
  if err != nil {
    log.Println(err)
  }

  for i := 0; i < num; i++ {
    indentSetting := document.Call("getElementById", "indent" + strconv.Itoa(i))
    lang := document.Call("getElementById", "lang" + strconv.Itoa(i))

    indentValue, err = strconv.Atoi(indentSetting.Get("value").String())
    if err != nil {
      log.Println(err)
    }

    language := lang.Get("value").String()

    langSetting := post.LanguageSettings{}
    langSetting.Language = language
    langSetting.Indent = indentValue

    orderVimrc.LangSetting = append(orderVimrc.LangSetting, langSetting)
  }

  return orderVimrc
}
