package main

import (
	"fmt"
	"syscall/js"

	"github.com/higashi000/vimrc-autogenerate/list/getList"
)

func main() {
	data := getList.GetList()
	fmt.Println(data)
	document := js.Global().Get("document")
	vimrc1 := document.Call("getElementById", "vimrc1")
	vimrc2 := document.Call("getElementById", "vimrc2")
	vimrc3 := document.Call("getElementById", "vimrc3")
	vimrc4 := document.Call("getElementById", "vimrc4")
	vimrc5 := document.Call("getElementById", "vimrc5")

	vimrc1.Set("href", "../result/index.html?uuid="+data.Inform[0].Uuid)
	vimrc1.Set("innerText", "1. "+data.Inform[0].Name)
	vimrc2.Set("href", "../result/index.html?uuid="+data.Inform[1].Uuid)
	vimrc2.Set("innerText", "2. "+data.Inform[1].Name)
	vimrc3.Set("href", "../result/index.html?uuid="+data.Inform[2].Uuid)
	vimrc3.Set("innerText", "3. "+data.Inform[2].Name)
	vimrc4.Set("href", "../result/index.html?uuid="+data.Inform[3].Uuid)
	vimrc4.Set("innerText", "4. "+data.Inform[3].Name)
	vimrc5.Set("href", "../result/index.html?uuid="+data.Inform[4].Uuid)
	vimrc5.Set("innerText", "5. "+data.Inform[4].Name)
}
