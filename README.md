# vimrc-autogenerate-front

## Description
This is front end for [vimrc-autogenerate](https://github.com/higashi000/vimrc-autogenerate).<br>
golang >= 1.13.6

## Usage
```
$ git clone https://github.com/higashi000/vimrc-autogenerate-front

$ cd vimrc-autogenerate-front

$ cp $(go env GOROOT)/misc/wasm/wasm_exec.js ./makeVimrc/

$ cp $(go env GOROOT)/misc/wasm/wasm_exec.js ./result/

$ cd makeVimrc

$ go build

$ ../

$ cd result

$ go build

$ ../

$ go run main.go
```

After, please access http://localhost:4000
